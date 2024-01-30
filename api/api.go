package api

import (
	"LightGate/midware"
	"LightGate/services"
	"LightGate/template"
	"LightGate/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(midware.Cors())
	r.LoadHTMLGlob("template/*")

	r.GET("/home", func(c *gin.Context) {
		service := services.GetServices()
		service["test"] = []string{"192.68.71.133:8080", "192.68.71.134:8080"}
		service["default"] = []string{"127.0.0.1:80"}
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, template.GetHome(service))
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
