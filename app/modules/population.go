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
	results := make(HistoryStepResults, len(this.individuals))

	for indIndex := range results {
		individual := this.individuals[indIndex]

		results[indIndex] = HistoryStepResult{
			Individual: &*individual,
			Rating: Test(individual),
		}
	}

	sort.Sort(results)
	SaveHistoryStep(results)

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
