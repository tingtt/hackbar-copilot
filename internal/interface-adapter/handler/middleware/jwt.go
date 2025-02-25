package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

const contextKeyJWT contextKey = "jwt"

var (
	ErrJWTInvalid  = errors.New("jwt: invalid token")
	ErrJWTNotFound = errors.New("jwt: token not found")
)

func GetJWT(ctx context.Context) (jwtclaims.Claims, error) {
	switch claimsOrErr := ctx.Value(contextKeyJWT).(type) {
	case jwtclaims.Claims:
		return claimsOrErr, nil
	case error:
		return jwtclaims.Claims{}, claimsOrErr
	case nil:
		panic("JWT middleware is not used.")
	default:
		panic(fmt.Sprintf("Unknown JWT claims type: %T", claimsOrErr))
	}
}

func setJWT(req *http.Request, claims jwtclaims.Claims) {
	if req == nil {
		panic("req is nil")
	}
	*req = *req.WithContext(context.WithValue(req.Context(), contextKeyJWT, claims))
}

func setJWTError(req *http.Request, err error) {
	if req == nil {
		panic("req is nil")
	}
	*req = *req.WithContext(context.WithValue(req.Context(), contextKeyJWT, err))
}

func JWT(secret []byte) (middleware func(http.Handler) http.Handler, usedContextKeys []contextKey) {
	if len(secret) == 0 {
		panic("JWT secret (--jwt.secret) is required.")
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				next.ServeHTTP(w, r)
			}()

			authorization := r.Header.Get("Authorization")
			if strings.HasPrefix(authorization, "Bearer ") {
				tokenStr := strings.TrimPrefix(authorization, "Bearer ")
				claims, err := parseJWT(tokenStr, secret)
				if err != nil {
					setJWTError(r, err)
					return
				}
				setJWT(r, claims)
				return
			}

			cookie, err := r.Cookie("jwt")
			if err != nil {
				setJWTError(r, fmt.Errorf("jwt: %w", err))
				return
			}
			if cookie == nil {
				setJWTError(r, ErrJWTNotFound)
				return
			}
			claims, err := parseJWT(cookie.Value, secret)
			if err != nil {
				setJWTError(r, err)
				return
			}
			setJWT(r, claims)
		})
	}, []contextKey{contextKeyJWT}
}

func parseJWT(tokenStr string, secret []byte) (jwtclaims.Claims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return jwtclaims.Claims{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jsonClaims, _ := json.Marshal(claims)
		return jwtclaims.Unmarshal(jsonClaims)
	}
	return jwtclaims.Claims{}, ErrJWTInvalid
}
