package api

import (
	"LightGate/midware"
	"LightGate/services"
	"LightGate/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(midware.Cors())

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
		heartbeat.Ip = util.ParseIp(c.Request)
		services.Store(heartbeat)
		c.Status(http.StatusOK)
	})

	r.Any("/:module/*path", func(context *gin.Context) {
		services.DoProxy(context.Writer, context.Request)
	})

	return r
}
