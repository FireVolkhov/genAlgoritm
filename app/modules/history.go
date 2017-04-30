package modules

import (
	"time"
	"../../app/core"
	//"io/ioutil"
)

const HistoryLen int = 200;

type HistoryStep struct {
	Time time.Time
	Results HistoryStepResults
}

type HistoryStepResults []*HistoryStepResult

type HistoryStepResult struct {
	Individual *Individual
	Index int
	Rating float64
}

func (slice HistoryStepResults) Len() int {
	return len(slice)
}

func (slice HistoryStepResults) Less(firstIndex, secondIndex int) bool {
	first := slice[firstIndex]
	second := slice[secondIndex]

	moreRating := first.Rating > second.Rating
	lessRating := first.Rating < second.Rating
	//lessGens := first.Individual.GetGensCount() < second.Individual.GetGensCount()

	if (moreRating) {
		return true
	} else if (lessRating) {
		return false
	} else {
		return core.RandomBool()
		//return !lessGens
	}
}

func (slice HistoryStepResults) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SaveHistoryStep (results HistoryStepResults) {
	newStep := &HistoryStep{
		Time: time.Now(),
		Results: results,
	}

	// Unshift
	history = append([]*HistoryStep{newStep}, history...)

	if (HistoryLen < len(history)) {
		history = history[:HistoryLen]
	}

	//now := time.Now()
	//d1 := []byte(results[0].Individual.ToString())
	//ioutil.WriteFile("./result/" + now.Format("2006-01-02_15-04.gen"), d1, 0644)
}

func GetHistory () []*HistoryStep {
	return history
}

var history = make([]*HistoryStep, 0)
