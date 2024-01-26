package main

import (
	"LightGate/services"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
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

	serviceCron := cron.New()
	err := serviceCron.AddFunc("*/10 * * * * ?", func() {
		services.ReloadServices()
	})
	if err != nil {
		return
	}
	serviceCron.Start()

	err = r.Run()
	if err != nil {
		return
	}
}
