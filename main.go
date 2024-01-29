package main

import (
	"LightGate/api"
	"LightGate/services"
	"github.com/robfig/cron"
)

func main() {

	serviceCron := cron.New()
	err := serviceCron.AddFunc("*/10 * * * * ?", func() {
		services.ReloadServices()
	})
	if err != nil {
		return
	}
	serviceCron.Start()

	r := api.Router()
	err = r.Run()
	if err != nil {
		return
	}
}
