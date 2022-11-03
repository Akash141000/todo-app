package routes

import (
	"todoBackend/src/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.RouterGroup) {
	c.POST("/signup", user.Signup)

	c.POST("/login", user.Login)
}
