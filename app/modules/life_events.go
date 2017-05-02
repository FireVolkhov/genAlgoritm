package modules

import (
	//"../../app/core"
)

func FireEvents (config *Config) {
	history := GetHistory()
	//lastRating := history[0].Results[0].Rating

	DisplayTickResult(history)

	// Set count event
	//isChangedRating := false

	//if (config.MutationShockSteps <= len(history)) {
	//	for stepIndex := 0; stepIndex < config.MutationShockSteps; stepIndex++ {
	//		if (lastRating != history[stepIndex].Results[0].Rating) {
	//			isChangedRating = true
	//		}
	//	}
	//
	//	if (!isChangedRating) {
	//		config.Count = int(core.Round(float64(config.Count) * config.MutationShockRatio))
	//		log.Printf("SET Count: %d", config.Count)
	//	}
	//}

	// Finish event
	//isChangedRating = false

	//if (config.FinishSteps <= len(history)) {
	//	for stepIndex := 0; stepIndex < config.FinishSteps; stepIndex++ {
	//		if (lastRating != history[stepIndex].Results[0].Rating) {
	//			isChangedRating = true
	//		}
	//	}
	//
	//	if (!isChangedRating) {
	//		DisplayResult(history)
	//		log.Println("--- FINISH ---")
	//		os.Exit(0)
	//	}
	//}

	//if (config.FinishSteps <= GetTick()) {
	//	DisplayResult(history)
	//	log.Println("--- FINISH ---")
	//	os.Exit(0)
	//}
}
