package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Protect(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			
			return nil, err
		}
		return []byte("==signature=="), nil
	})
	return err
}
