package common

import (
	"time"

	"github.com/donnyirianto/go-be-fiber/configuration"
	"github.com/donnyirianto/go-be-fiber/exception"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string, roles []map[string]interface{}, config configuration.Config) string {
	jwtSecret := config.GetString("JWT_SECRET_KEY")
	jwtExpired := config.GetInt("JWT_EXPIRE_MINUTES_COUNT")

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
