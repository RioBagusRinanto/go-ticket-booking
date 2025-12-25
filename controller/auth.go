package controller

import (
	"net/http"

	"ticket-booking/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var json model.Login
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "lucario" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
