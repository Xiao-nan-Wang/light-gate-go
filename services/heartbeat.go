package services

import (
	"sync"
)

var mutex sync.Mutex
var services map[string][]string
var temporaryHeartbeatStorage []Heartbeat

type Heartbeat struct {
	Name string
	Ip   string
	Port string
}

func GetServices() map[string][]string {
	return services
}

func Store(heartbeat Heartbeat) {
	mutex.Lock()
	temporaryHeartbeatStorage = append(temporaryHeartbeatStorage, heartbeat)
	mutex.Unlock()
}

func ReloadServices() {
	for _, heartbeat := range temporaryHeartbeatStorage {
		addresses, ok := services[heartbeat.Name]
		if ok {
			addresses = append(addresses, heartbeat.Ip+":"+heartbeat.Port)
		} else {
			addresses = []string{heartbeat.Ip + ":" + heartbeat.Port}
			services[heartbeat.Name] = addresses
		}
	}
}
