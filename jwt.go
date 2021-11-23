package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtKey = "abc123ABC"

// JWTEncrypt 加密
func JWTEncrypt(effectTimes int, param ...map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims, 0)
	claims["exp"] = fmt.Sprintf("%d", time.Now().Add(time.Second*time.Duration(effectTimes)).Unix())
	claims["iat"] = fmt.Sprintf("%d", time.Now().Unix())

	if len(param) > 0 {
		for _, val := range param {
			for k, v := range val {
				claims[k] = v
			}
		}
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(jwtKey))
	return tokenString
}

// JWTDecrypt 解密
func JWTDecrypt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil
	}
	return claims
}
