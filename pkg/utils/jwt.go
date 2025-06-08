package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("ct_sys_api_root")

func GenerateToken(userID uint, userRole string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"role": userRole,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(jwtSecret)
}
