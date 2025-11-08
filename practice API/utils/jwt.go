package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey= "asew234rtpoy90_322!3x"

func TokenGeneration(email string, id int ) (string, error) {
	claims := &jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(59 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedInString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedInString, nil
}