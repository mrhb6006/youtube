package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Payload struct {
	UserID int64 `json:"user_id"`
	Exp    int64 `json:"exp"`
}

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("invalid token")

func (payload *Payload) Valid() error {
	if time.Now().Unix() > payload.Exp {
		return ErrExpiredToken
	}
	return nil
}

func CreateToken(userID int64) (string, error) {
	var err error
	payload := Payload{
		UserID: userID,
		Exp:    time.Now().Add(time.Hour * 24).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &payload)
	token, err := t.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
