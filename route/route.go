package route

import (
	"net/http"

	"ticket-booking/controller"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// Health check endpoint
	router.GET("/ping", Ping)

	// Auth endpoints
	router.POST("/login", controller.Login)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
