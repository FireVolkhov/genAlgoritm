package main

import (
	"../app/modules"
//	"../app/core"
	"os"
	"log"
	"io/ioutil"
//	"path/filepath"
//	"strings"
	"../app/nnetwork"
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
//
//
//	lastStepFile := 0
//
//	// walk all files in directory
//	filepath.Walk("./result/", func(path string, info os.FileInfo, err error) error {
//		if (!info.IsDir()) {
//			names := strings.Split(info.Name(), ".")
//
//			if (names[1] == "pop") {
//				name := core.StringToInt(names[0])
//
//				if (lastStepFile < name) {
//					lastStepFile = name
//				}
//			}
//		}
//		return nil
//	})
//
//	if (0 < lastStepFile) {
//		file, _ := ioutil.ReadFile("./result/" + core.IntToString(lastStepFile) + ".pop")
//		dataAsString := string(file)
//		indStrings := strings.Split(dataAsString, ";")
//
//		individual := modules.StringToIndividual(indStrings[0])
//		report := modules.GetReport(individual)
//		ioutil.WriteFile("./result/report.csv", []byte(report), 0644)
//		log.Println("### Report completed: " + core.IntToString(lastStepFile) + ".pop ###")
//	}

	file, _ := ioutil.ReadFile("./result/target.net")
	network := nnetwork.NNetworkFromBytes(file)

	report := modules.GetReportForNNetwork(network)
	ioutil.WriteFile("./result/report.csv", []byte(report), 0644)
	log.Println("### Report completed ###")
}
