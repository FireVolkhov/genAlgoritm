package main

import (
	"../app/modules"
	"os"
	"log"
	"io/ioutil"
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

	file, _ := ioutil.ReadFile("./result/target.gen")
	individual := modules.FromString(string(file))
	report := modules.GetReport(individual)
	ioutil.WriteFile("./result/report.csv", []byte(report), 0644)
}