package jwt_helper

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type DecodedToken struct {
	Iat     int       `json:"iat"`
	IsAdmin bool      `json:"roles"`
	UserId  uuid.UUID `json:"userId"`
	Email   string    `json:"email"`
	Iss     string    `json:"iss"`
}

func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return token
}

func VerifyToken(token string, secret string) (*DecodedToken, error) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !decoded.Valid {
		return nil, err
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)
	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken, nil
}
