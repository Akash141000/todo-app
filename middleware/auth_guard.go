package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthGuard(c *gin.Context) {
	fmt.Println("AUTH MIDDLEWARE >>")
	authToken := c.Request.Header["Authentication"]
	if authToken == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Unauthorized request"})
		return
	}
	// c.Set("user", "Authencated User")
	c.Next()
}
