package common

import (
	"strconv"
	"time"

	"github.com/donnyirianto/go-be-fiber/configuration"
	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string, roles []map[string]interface{}, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
