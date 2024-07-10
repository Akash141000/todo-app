package middleware

import (
	"fmt"
	"net/http"
	"todoBackend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthGuard(c *gin.Context) {
	fmt.Println("AUTH MIDDLEWARE >>")

	// if path is "/" then don't check token
	if c.Request.URL.Path == "/" {
		c.Next()
		return
	}
	authToken := c.Request.Header["Authorization"]
	// fmt.Println("AUTH HEADER>>", authToken)
	if authToken == nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Unauthorized request"})
		return
	}
	// fmt.Println("TOKEN STRING", authToken[0])

	claims := &utils.Claims{}

	token, tokenParseErr := jwt.ParseWithClaims(authToken[0], claims, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("TOKEN CLAIMS>>", claims)
		fmt.Println("TOKEN >>", token)
		return []byte("jwtSigningKey"), nil
	})

	// fmt.Println("CLAIMS", tokenClaims)
	fmt.Println("ERROR>>", tokenParseErr)
	if tokenParseErr != nil {
		if tokenParseErr == jwt.ErrSignatureInvalid {
			fmt.Println("Error Invalid signature!")
		}
	}

	fmt.Println("TOKEN VALID>>", token.Valid)
	if !token.Valid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Next()

}
