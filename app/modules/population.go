package modules

import (
	"../../app/core"
	"sort"
	"time"
	"io/ioutil"
	"strings"
	"path/filepath"
	"os"
)

type Population struct {
	individuals []*Individual
}

func NewPopulation (config *Config) *Population {
	individuals, ok := getPopulationFromFile(config)

	if (ok) {
		return &Population{
			individuals: individuals,
		}
	} else {
		population := Population{
			individuals: make([]*Individual, config.Count),
		}

		for indIndex := range population.individuals {
			gensCount := core.RandomInt(config.StartGenCountMin, config.StartGenCountMax)
			individual := NewIndividual(config.EnterGens, gensCount)

			population.individuals[indIndex] = individual
		}

		return &population
	}
}

var inputChanel = make(chan *HistoryStepResult, 100)
var outputChanel = make(chan *HistoryStepResult, 100)

func (this *Population) Selection (config *Config) {
	indLen := len(this.individuals)
	results := make(HistoryStepResults, indLen)

	for indIndex, individual := range this.individuals {
		results[indIndex] = &HistoryStepResult{
			Individual: individual.Clone(),
			Index: indIndex,
			Rating: 0,
		}
	}

	resultIndex := 0
	completedResults := 0

	for {
		select {
		case historyStepResult := <-outputChanel:
			results[historyStepResult.Index] = historyStepResult
			completedResults++

			if (len(results) <= completedResults) {
				goto endCalcPopulation
			}
		default:
			if (resultIndex < len(results)) {
				inputChanel <- results[resultIndex]
				resultIndex++
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	endCalcPopulation:


	//chanel := make(chan *HistoryStepResult)
	//
	//for _, individual := range this.individuals {
	//	historyStepResult := &HistoryStepResult{
	//		Individual: individual.Clone(),
	//		Rating: 0,
	//	}
	//
	//	go calcIndividual(historyStepResult, individual, chanel)
	//}
	//
	//for indIndex := 0; indIndex < indLen; {
	//	results[indIndex] = <- chanel
	//	indIndex++
	//}

	sort.Sort(results)
	SaveHistoryStep(results[:1])

	//Shauffel
	//for resultIndex := 1; resultIndex < len(results); resultIndex++ {
	//	results[resultIndex].Rating = results[resultIndex].Rating * rand.Float64()
	//}
	//
	//sort.Sort(results)

	//oldPopulation := this.individuals
	//topIndividual := results[:1][0].Individual
	//saveCount := int(core.Round(config.SurvivalPercent * float64(len(results))))
	//this.individuals = make([]*Individual, saveCount)
	//this.individuals[0] = topIndividual.Clone()
	//
	//countSaved := 1
	//indexes := []int{0}
	//
	//for countSaved < saveCount {
	//	candidateIndex := core.RandomInt(0, len(oldPopulation) - 1)
	//
	//	if (!core.SliceContainsInt(indexes, candidateIndex)) {
	//		indexes = append(indexes, candidateIndex)
	//		this.individuals[countSaved] = oldPopulation[candidateIndex].Clone()
	//		countSaved++
	//	}
	//}

	results = results[:int(core.Round(config.SurvivalPercent * float64(len(results))))]

	this.individuals = make([]*Individual, len(results))

	for indIndex, result := range results {
		this.individuals[indIndex] = result.Individual
	}
}

func (this *Population) Mutation (config *Config) {
	//start := time.Now()
	for (len(this.individuals) < config.Count) {
		individual := this.individuals[core.RandomInt(0, len(this.individuals) - 1)]
		mutant := individual.Mutation()
		this.individuals = append(this.individuals, mutant)
		//if (!this.hasEqual(mutant)) {
		//	this.individuals = append(this.individuals, mutant)
		//}
	}
	//log.Printf("Time mutation: %s", time.Now().Sub(start))
}

func (this *Population) Dump (config *Config) {
	step := GetTick()
	individuals := make([]string, len(this.individuals))

	for indIndex, indInPop := range this.individuals {
		individuals[indIndex] = indInPop.ToString()
	}
	d1 := []byte(strings.Join(individuals, ";"))
	ioutil.WriteFile("./result/" + core.IntToString(step) + ".pop", d1, 0644)
}

func (this *Population) hasEqual (individual *Individual) bool {
	for _, indInPopulation := range this.individuals {
		if (individual.IsEqual(indInPopulation)) {
			return true
		}
	}

	return false
}

type caclIndividualChanelItem struct {
	indIndex int
	result float64
}

func calcIndividual (historyStep *HistoryStepResult, individual *Individual, chanel chan<- *HistoryStepResult) {
	historyStep.Rating = Test(individual)
	chanel <- historyStep
}

func getPopulationFromFile (config *Config) (individuals []*Individual, ok bool) {
	individuals = []*Individual{}
	ok = false

	lastStepFile := 0

	// walk all files in directory
	filepath.Walk("./result/", func(path string, info os.FileInfo, err error) error {
		if (!info.IsDir()) {
			names := strings.Split(info.Name(), ".")

			if (names[1] == "pop") {
				name := core.StringToInt(names[0])

				if (lastStepFile < name) {
					lastStepFile = name
				}
			}
		}
		return nil
	})

	if (0 < lastStepFile) {
		file, _ := ioutil.ReadFile("./result/" + core.IntToString(lastStepFile) + ".pop")
		dataAsString := string(file)
		indStrings := strings.Split(dataAsString, ";")
		individuals = make([]*Individual, len(indStrings))

		for indIndex, individualString := range indStrings {
			individuals[indIndex] = StringToIndividual(individualString)
		}

		SetTick(lastStepFile)

		ok = true
		return individuals, ok
	} else {

		return individuals, ok
	}
}

func targetFunctionProcessor() {
	for {
		historyStep := <- inputChanel
		historyStep.Rating = Test(historyStep.Individual)
		outputChanel <- historyStep
	}
}

func init () {
	for i := 0; i < 8; i++ {
		go targetFunctionProcessor()
	}
}
