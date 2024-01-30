package main

import (
	"LightGate/api"
	"LightGate/services"
	"flag"
	"github.com/robfig/cron"
	"log"
)

var port = flag.String("port", "8080", "运行端口")

func main() {
	flag.Parse()

	serviceCron := cron.New()
	err := serviceCron.AddFunc("*/10 * * * * ?", func() {
		services.ReloadServices()
	})
	if err != nil {
		return
	}
	serviceCron.Start()

	r := api.Router()
	err = r.RunTLS(":"+*port, "conf/server.crt", "conf/server.key")
	if err != nil {
		log.Println("SSL配置失败，启动HTTP")
		err = r.Run(":" + *port)
		if err != nil {
			return
		}
	}
}
