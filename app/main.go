package main

import (
	"../app/modules"
	"os"
	"log"
)

func main() {
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
	}
}
