package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"main/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		token := c.GetHeader("Authorization")

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		_, errVerifyToken := utils.VerifyToken(token)

		if errVerifyToken != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error verifying JWT token: " + errVerifyToken.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
