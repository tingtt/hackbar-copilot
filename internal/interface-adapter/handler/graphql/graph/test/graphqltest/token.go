package graphqltest

import (
	"encoding/json"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

type mapClaims jwtclaims.Claims

func (c mapClaims) collect() map[string]any {
	bytes, err := json.Marshal(c)
	if err != nil {
		panic("failed to marshal claim to json: " + err.Error())
	}
	mapped := map[string]any{}
	err = json.Unmarshal(bytes, &mapped)
	if err != nil {
		panic("failed to unmarhsal json to map[string]any: " + err.Error())
	}
	return mapped
}

func NewToken(claims jwtclaims.Claims) string {
	jwt := jwtauth.New("HS256", []byte(jwtSecret), nil)

	mapClaim := mapClaims(claims).collect()
	jwtauth.SetIssuedNow(mapClaim)
	jwtauth.SetExpiryIn(mapClaim, time.Hour)

	_, tokenString, err := jwt.Encode(mapClaim)
	if err != nil {
		panic(err)
	}
	return tokenString
}
