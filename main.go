package main

import (
	"LightGate/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/alive-services", func(c *gin.Context) {
		res := services.GetServices()
		c.JSON(http.StatusOK, res)
	})

	r.POST("/heartbeat", func(c *gin.Context) {
		var heartbeat services.Heartbeat
		err := c.BindJSON(&heartbeat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		heartbeat.Ip = c.ClientIP()
		services.Store(heartbeat)
		c.Status(http.StatusOK)
	})

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
