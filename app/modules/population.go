package modules

import (
	"../../app/core"
	"sort"
	"reflect"
	"math/rand"
	"time"
	"log"
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
	//chanel := make(chan *HistoryStepResult)

	for indIndex, individual := range this.individuals {
		results[indIndex] = &HistoryStepResult{
			Individual: individual.Clone(),
			Rating: Test(individual),
		}

		//go calcIndividual(historyStepResult, individual, chanel)
	}

	//for indIndex := 0; indIndex < indLen; {
	//	results[indIndex] = <- chanel
	//	indIndex++
	//}

	sort.Sort(results)
	SaveHistoryStep(results[:1])

	for resultIndex := 1; resultIndex < len(results); resultIndex++ {
		results[resultIndex].Rating = results[resultIndex].Rating * rand.Float64()
	}

	sort.Sort(results)

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
	start := time.Now()
	for (len(this.individuals) < config.Count) {
		individual := this.individuals[core.RandomInt(0, len(this.individuals) - 1)]
		mutant := individual.Mutation()
		this.individuals = append(this.individuals, mutant)
		//if (!this.hasEqual(mutant)) {
		//	this.individuals = append(this.individuals, mutant)
		//}
	}
	log.Printf("Time mutation: %s", time.Now().Sub(start))
}

func (this *Population) hasEqual (individual *Individual) bool {
	for _, indInPopulation := range this.individuals {
		if (reflect.DeepEqual(individual, indInPopulation)) {
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
