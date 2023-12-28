package utils

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cristalhq/jwt/v5"
	"github.com/google/uuid"
)

var key = []byte(`secret`)
var signer, _ = jwt.NewSignerHS(jwt.HS256, key)
var verifier, _ = jwt.NewVerifierHS(jwt.HS256, key)
var issuer = "open.mxzero.top"

func CreateToken(subject map[string]interface{}) string {
	subjectStr, _ := json.Marshal(subject)

	var current time.Time = time.Now()
	var express time.Time = current.Add(2 * time.Hour)

	claims := &jwt.RegisteredClaims{
		ID:        uuid.New().String(),
		Subject:   string(subjectStr),
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(current),
		ExpiresAt: jwt.NewNumericDate(express),
	}
	token, _ := jwt.NewBuilder(signer).Build(claims)
	return token.String()
}

func VerifyToken(token string) bool {
	var claims jwt.RegisteredClaims
	return jwt.ParseClaims([]byte(token), verifier, &claims) == nil && claims.IsValidAt(time.Now())
}

func ParseToken(token string) (map[string]interface{}, error) {
	if !VerifyToken(token) {
		return nil, errors.New("token无效")
	}

	var claims jwt.RegisteredClaims
	err := jwt.ParseClaims([]byte(token), verifier, &claims)
	if err != nil {
		return nil, err
	}

	var subjectData map[string]interface{}
	err = json.Unmarshal([]byte(claims.Subject), &subjectData)

	return subjectData, err
}
