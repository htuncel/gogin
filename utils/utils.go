package utils

import (
	"net"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"

	"main/configs"
)

func VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return configs.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       0,
		"identity": "00000000000",
		"title":    "",
		"iat":      time.Now().Unix(),
		"exp":      (time.Now().Local().Add(time.Second * time.Duration(86400)).Unix()),
	})

	return token.SignedString(configs.Secret)
}

func GetToken(c *gin.Context) string {
	reqToken := c.GetHeader("authorization")
	splitToken := strings.Replace(reqToken, "Bearer ", "", 1)
	return splitToken
}

func GetClientIP(c *gin.Context) string {
	if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
		if parts := strings.Split(xff, ","); len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}

	host, _, err := net.SplitHostPort(c.Request.RemoteAddr)

	if err != nil {
		return c.Request.RemoteAddr
	}

	return host
}
