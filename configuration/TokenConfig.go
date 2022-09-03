package configuration

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var api string

var privateKey *rsa.PrivateKey

type TokenHeader struct {
	Username string `json:"un,omitempty"`
	Role     int    `json:"rl,omitempty"`
	Api      string `json:"api,omitempty"`
	jwt.RegisteredClaims
}

func init() {
	api = base64.StdEncoding.EncodeToString([]byte(uuid.Must(uuid.NewRandom()).String()))
	private, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatal("Error Generate Token !!!")
	}
	privateKey = private
}

func CreateToken(username string, email string, role int) (string, bool) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, &TokenHeader{
		Username: username,
		Role:     role,
		Api:      api,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(86400000)),
			ID:        base64.StdEncoding.EncodeToString([]byte(uuid.Must(uuid.NewRandom()).String())),
		},
	})
	res, err := token.SignedString(privateKey)
	if err != nil {
		return "", false
	}
	return res, true
}

func Bind(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "JLMS_TOKEN",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(86400000)),
		HttpOnly: true,
		Secure:   true,
	})
}

func ParseToken(token string) (*TokenHeader, error) {
	tk, err := jwt.ParseWithClaims(token, &TokenHeader{}, func(t *jwt.Token) (interface{}, error) {
		return privateKey.PublicKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, jwt.ErrTokenMalformed
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.ErrTokenExpired
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, jwt.ErrTokenNotValidYet
			} else {
				return nil, jwt.ErrTokenMalformed
			}
		}
	}
	if tk != nil {
		if cl, ok := tk.Claims.(*TokenHeader); ok && tk.Valid {
			return cl, nil
		}
		return nil, jwt.ErrTokenMalformed
	} else {
		return nil, jwt.ErrTokenMalformed
	}
}
