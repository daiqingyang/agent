package main

import (
	"github.com/gin-gonic/gin"
)

func StartHttp() {
	port := ":44444"
	eng := gin.Default()
	eng.GET("/info", func(c *gin.Context) {
		agent := NewAgent()
		c.JSON(200, agent)
	})
	eng.Run(port)
}
