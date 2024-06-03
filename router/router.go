package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}