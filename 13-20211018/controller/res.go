package controller

import (
	"github.com/gin-gonic/gin"
)

func ResponseErr(c *gin.Context, statusCode int, message string) error {
	c.JSON(statusCode, gin.H{
		"message": message,
	})
	return nil
}
