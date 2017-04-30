package main

import (
	"../app/modules"
	"time"
)

func checkState() {
	config := modules.GetAppConfig()

	for {
		modules.CreateTasks(config)
		modules.CheckCompletedTasks()
		time.Sleep(time.Second)
	}
}

func main() {
	go checkState()

	modules.StartServer()
}
