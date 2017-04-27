package modules

import "time"

type HistoryStep struct {
	Time time.Time
	Results HistoryStepResults
}

type HistoryStepResults []HistoryStepResult

type HistoryStepResult struct {
	Individual *Individual
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
	lessGens := first.Individual.GetGensCount() < second.Individual.GetGensCount()

	if (moreRating) {
		return true
	} else if (lessRating) {
		return false
	} else {
		return lessGens
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
}

func GetHistory () []*HistoryStep {
	return history
}

var history []*HistoryStep

func init() {
	history = make([]*HistoryStep, 0)
}