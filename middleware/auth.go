package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	fmt.Println("middlware running...")
	authHeader := c.GetHeader("Authorization")
	fmt.Println("Authorization header:", authHeader)
	if authHeader != "123" && c.Request.Method != "GET" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}
	c.Next()
}
