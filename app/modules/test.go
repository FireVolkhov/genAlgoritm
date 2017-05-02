package modules

import (
	"../../app/test/data"
	"../../app/core"
	"math"
	"log"
	"sync"
	"fmt"
	"time"
)

var tick int = 0

func GetTick () int {
	return tick
}

func SetTick (_tick int) {
	tick = _tick
}

func DisplayTickResult (history []*HistoryStep) {
	tick++
	historyLen := len(history)
	topRating := history[0].Results[0].Rating
	percent := topRating * 100
	percentYear := (math.Pow(topRating, 1 / (float64(data.GetTableLen()) / 365)) - 1) * 100
	goodMonkeyPercent := history[0].Results[0].GoodMonkey * 100
	var elapsed time.Duration = 0

	if (1 < historyLen) {
		elapsed = history[0].Time.Sub(history[1].Time)
	}

	log.Println(
		fmt.Sprintf("%3.0f ", float64(tick)) +
		fmt.Sprintf("FR: %1.0f%s ", percent, "%") +
		fmt.Sprintf("YR: %3.2f%s ", percentYear, "%") +
		fmt.Sprintf("GS: %3.2f%s ", goodMonkeyPercent, "%") +
		fmt.Sprintf("%s", elapsed))
}

func DisplayResult (history []*HistoryStep) {
	elapsedAllTime := history[0].Time.Sub(history[len(history) - 1].Time)
	topRating := history[0].Results[0].Rating
	percentYear := math.Pow(topRating, 1 / (float64(data.GetTableLen()) / 365)) - 1

	log.Printf("Time: %s", elapsedAllTime)
	log.Println(history[0].Results[0].Individual.ToString())
	log.Println(history[0].Results[0].Individual.ToClearGenome().ToString())
	log.Printf("Result All time: %10.0f%s", topRating * 100, "%")
	log.Printf("Result Year: %.2f%s", percentYear * 100, "%")
}

func Test (individual *Individual) (float64, float64) {
	cacheKey := individual.ToCacheKey()
	cacheResult, isGoodMonkeyResult, ok := getFromCache(cacheKey, individual)

	if (ok) {
		return cacheResult, isGoodMonkeyResult

	} else {
		iterator := data.NewTableIterator(8)
		result := float64(1)
		isGoodMonkeyResult := float64(0)

		for (!iterator.IsFinished) {
			result, _, _, isGoodMonkeyResult = iterationStep(iterator, individual, result, isGoodMonkeyResult)
		}

		isGoodMonkeyResult = isGoodMonkeyResult / float64(data.GetTableLen())

		addToCache(cacheKey, individual, result, isGoodMonkeyResult)

		return result, isGoodMonkeyResult
	}
}

func GetReport (individual *Individual) string {
	iterator := data.NewTableIterator(8)
	result := float64(1)
	history := make([]*historyStep, 0)

	for (!iterator.IsFinished) {
		var futureDay *data.TableRow
		var monkeyResult float64
		isGoodMonkeyResult := float64(0)

		result, futureDay, monkeyResult, isGoodMonkeyResult = iterationStep(iterator, individual, result, isGoodMonkeyResult)

		step := &historyStep{
			tableRow: &*futureDay,
			monkeyResult: monkeyResult,
			balance: result,
		}

		history = append(history, step)
	}

	return formatHistoryToCSV(history)
}



// --- PRIVATE ---------------------------------------------------------------------------------------------------------
type cacheItem struct {
	isGoodMonkeyResult float64
	result float64
}

var resultCache = make(map[string]*cacheItem, 0)
var resultMutex = &sync.RWMutex{}

func addToCache(key string, individual *Individual, result, isGoodMonkeyResult float64) {
	item := &cacheItem{
		isGoodMonkeyResult: isGoodMonkeyResult,
		result: result,
	}
	resultMutex.Lock()
	resultCache[key] = item
	resultMutex.Unlock()
}

func getFromCache(key string, individual *Individual) (float64, float64, bool) {
	resultMutex.RLock()
	item, ok := resultCache[key]
	resultMutex.RUnlock()

	if (ok) {
		return item.result, item.isGoodMonkeyResult, true
	} else {
		return 0, 0, false
	}
}

func iterationStep (iterator *data.TableIterator, individual *Individual, result, isGoodMonkeyResult float64) (float64, *data.TableRow, float64, float64) {
	slice := iterator.Next()
	futureDay := slice[len(slice) - 1]
	slice = slice[:len(slice) - 1]
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
		isGoodMonkeyResult = isGoodMonkeyResult + 1
	} else if (
	(monkeyBay && (marketSell || marketWait)) ||
	(monkeySell && (marketBay || marketWait))) {
		result = result * (1 - futureDayPercent)
	}

	return result, futureDay, monkeyResult, isGoodMonkeyResult
}

type historyStep struct {
	tableRow *data.TableRow
	monkeyResult float64
	balance float64
}

func formatHistoryToCSV (history []*historyStep) string {
	result := "Open,High,Low,Close,Volume,Dif,MonkeyResult,Balance;\n"

	for _, step := range history {
		result = result + core.Float64ToString(step.tableRow.Open) + ","
		result = result + core.Float64ToString(step.tableRow.High) + ","
		result = result + core.Float64ToString(step.tableRow.Low) + ","
		result = result + core.Float64ToString(step.tableRow.Close) + ","
		result = result + core.IntToString(step.tableRow.Volume) + ","
		result = result + core.Float64ToString(step.tableRow.Dif) + ","
		result = result + core.Float64ToString(step.monkeyResult) + ","
		result = result + core.Float64ToString(step.balance) + ";\n"
	}

	return result;
}
