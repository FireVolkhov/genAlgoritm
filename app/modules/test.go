package modules

import (
	"../../app/test/data"
	"math"
	"log"
)

func DisplayTickResult (history []*HistoryStep) {
	historyLen := len(history)
	percent := history[0].Results[0].Rating * 100

	if (1 < historyLen) {
		elapsed := history[0].Time.Sub(history[1].Time)
		log.Printf("%3.0f STEP %10.0f%s %s", float64(historyLen), percent, "%", elapsed)
	} else {
		log.Printf("%3.0f STEP %10.0f%s", float64(historyLen), percent, "%")
	}
}

func DisplayResult (history []*HistoryStep) {
	elapsedAllTime := history[0].Time.Sub(history[len(history) - 1].Time)
	topRating := history[0].Results[0].Rating
	percentYear := math.Pow(topRating, 1 / (float64(data.GetTableLen()) / 365)) - 1

	log.Printf("Time: %s", elapsedAllTime)
	log.Println(history[0].Results[0].Individual.ToString())
	log.Printf("Result All time: %10.0f%s", topRating * 100, "%")
	log.Printf("Result Year: %.2f%s", percentYear * 100, "%")
}

func Test (individual *Individual) float64 {
	iterator := data.NewTableIterator(8)
	result := float64(1)

	for (!iterator.IsFinished) {
		slice := iterator.Next()
		futureDay := slice[len(slice) - 1]
		futureDayDif := futureDay.Close - futureDay.Open
		futureDayPercent := math.Abs(futureDayDif) / futureDay.Open

		args := make([]float64, len(slice))

		for rowIndex, row := range slice {
			args[rowIndex] = row.Dif
		}

		monkeyResult := individual.Run(args)

		marketBay := 0 < futureDayDif
		marketWait := futureDayDif == 0
		marketSell := futureDayDif < 0

		monkeyBay := 0 < monkeyResult
		//monkeyWait := monkeyResult == 0
		monkeySell := monkeyResult < 0

		if ((monkeyBay && marketBay) ||
			(monkeySell && marketSell)) {
			result = result * (1 + futureDayPercent)
		} else if (
			(monkeyBay && (marketSell || marketWait)) ||
			(monkeySell && (marketBay || marketWait))) {
			result = result * (1 - futureDayPercent)
		}
	}

	return result
}
