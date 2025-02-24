package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Email string `json:"emails"`
}

const contextKeyJWT contextKey = "jwt"

var (
	ErrJWTInvalidClaimsType = errors.New("jwt: invalid claims type")
	ErrJWTInvalid           = errors.New("jwt: invalid token")
	ErrJWTNotFound          = errors.New("jwt: token not found")
)

func GetJWT(ctx context.Context) (JWTClaims, error) {
	switch claimsOrErr := ctx.Value(contextKeyJWT).(type) {
	case JWTClaims:
		return claimsOrErr, nil
	case error:
		return JWTClaims{}, claimsOrErr
	case nil:
		panic("JWT middleware is not used.")
	default:
		panic(fmt.Sprintf("Unknown JWT claims type: %T", claimsOrErr))
	}
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
					*r = *r.WithContext(context.WithValue(r.Context(), contextKeyJWT, err))
					return
				}
				*r = *r.WithContext(context.WithValue(r.Context(), contextKeyJWT, claims))
				return
			}

			cookie, err := r.Cookie("jwt")
			if err != nil {
				*r = *r.WithContext(context.WithValue(r.Context(), contextKeyJWT, fmt.Errorf("jwt: %w", err)))
				return
			}
			claims, err := parseJWT(cookie.Value, secret)
			if err != nil {
				*r = *r.WithContext(context.WithValue(r.Context(), contextKeyJWT, err))
				return
			}
			*r = *r.WithContext(context.WithValue(r.Context(), contextKeyJWT, claims))
		})
	}, []contextKey{contextKeyJWT}
}

func parseJWT(tokenStr string, secret []byte) (JWTClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return JWTClaims{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		emails, ok := claims["emails"].([]interface{})
		if ok {
			if len(emails) == 0 {
				return JWTClaims{}, fmt.Errorf("%w (email not found)", ErrJWTInvalidClaimsType)
			}
			contextClaims := JWTClaims{}
			if email, ok := emails[0].(string); ok {
				contextClaims.Email = email
			} else {
				return JWTClaims{}, fmt.Errorf("%w (email type)", ErrJWTInvalidClaimsType)
			}
			return contextClaims, nil
		} else {
			return JWTClaims{}, fmt.Errorf("%w (email[] type)", ErrJWTInvalidClaimsType)
		}
	} else {
		return JWTClaims{}, ErrJWTInvalid
	}
}
