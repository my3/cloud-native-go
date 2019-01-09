package ping

import "github.com/gin-gonic/gin"

func RespondToPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
