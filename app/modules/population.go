package modules

import (
	"../../app/core"
	"sort"
)

type Population struct {
	individuals []*Individual
}

func NewPopulation (config *Config) *Population {
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

func (this *Population) Selection (config *Config) {
	indLen := len(this.individuals)
	results := make(HistoryStepResults, indLen)
	chanel := make(chan *HistoryStepResult)

	for _, individual := range this.individuals {
		historyStepResult := &HistoryStepResult{
			Individual: &*individual,
			Rating: 0,
		}

		go calcIndividual(historyStepResult, individual, chanel)
	}

	for indIndex := 0; indIndex < indLen; {
		results[indIndex] = <- chanel
		indIndex++
	}




	//for indIndex := range results {
	//	individual := this.individuals[indIndex]
	//
	//	results[indIndex] = HistoryStepResult{
	//		Individual: &*individual,
	//		Rating: Test(individual),
	//	}
	//}

	sort.Sort(results)
	SaveHistoryStep(results[:1])

	results = results[:int(core.Round(config.SurvivalPercent * float64(len(results))))]

	this.individuals = make([]*Individual, len(results))

	for indIndex, result := range results {
		this.individuals[indIndex] = result.Individual
	}
}

func (this *Population) Mutation (config *Config) {
	for (len(this.individuals) < config.Count) {
		individual := this.individuals[core.RandomInt(0, len(this.individuals) - 1)]
		this.individuals = append(this.individuals, individual.Mutation())
	}
}

type caclIndividualChanelItem struct {
	indIndex int
	result float64
}

func calcIndividual (historyStep *HistoryStepResult, individual *Individual, chanel chan<- *HistoryStepResult) {
	historyStep.Rating = Test(individual)
	chanel <- historyStep
}
