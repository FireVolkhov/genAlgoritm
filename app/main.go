package main

import (
	"../app/modules"
	"os"
	"log"

	//_ "net/http/pprof"
	//"net/http"
)

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	log.SetOutput(os.Stdout)

	log.Println("	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	log.Println("	░░░░░░░▄▄▀▀▀▀▀▀▀▀▀▀▄▄█▄░░░░▄░░░░█░░░░░░░")
	log.Println("	░░░░░░█▀░░░░░░░░░░░░░▀▀█▄░░░▀░░░░░░░░░▄░")
	log.Println("	░░░░▄▀░░░░░░░░░░░░░░░░░▀██░░░▄▀▀▀▄▄░░▀░░")
	log.Println("	░░▄█▀▄█▀▀▀▀▄░░░░░░▄▀▀█▄░▀█▄░░█▄░░░▀█░░░░")
	log.Println("	░▄█░▄▀░░▄▄▄░█░░░▄▀▄█▄░▀█░░█▄░░▀█░░░░█░░░")
	log.Println("	▄█░░█░░░▀▀▀░█░░▄█░▀▀▀░░█░░░█▄░░█░░░░█░░░")
	log.Println("	██░░░▀▄░░░▄█▀░░░▀▄▄▄▄▄█▀░░░▀█░░█▄░░░█░░░")
	log.Println("	██░░░░░▀▀▀░░░░░░░░░░░░░░░░░░█░▄█░░░░█░░░")
	log.Println("	██░░░░░░░░░░░░░░░░░░░░░█░░░░██▀░░░░█▄░░░")
	log.Println("	██░░░░░░░░░░░░░░░░░░░░░█░░░░█░░░░░░░▀▀█▄")
	log.Println("	██░░░░░░░░░░░░░░░░░░░░█░░░░░█░░░░░░░▄▄██")
	log.Println("	░██░░░░░░░░░░░░░░░░░░▄▀░░░░░█░░░░░░░▀▀█▄")
	log.Println("	░▀█░░░░░░█░░░░░░░░░▄█▀░░░░░░█░░░░░░░▄▄██")
	log.Println("	░▄██▄░░░░░▀▀▀▄▄▄▄▀▀░░░░░░░░░█░░░░░░░▀▀█▄")
	log.Println("	░░▀▀▀▀░░░░░░░░░░░░░░░░░░░░░░█▄▄▄▄▄▄▄▄▄██")
	log.Println("	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	log.Println("### Start App ###")

	config := modules.GetAppConfig()
	population := modules.NewPopulation(config)

	for {
		population.Selection(config)
		modules.FireEvents(config)
		population.Mutation(config)
		population.Dump(config)
	}
}
