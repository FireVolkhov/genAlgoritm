package modules

type Config struct {
	Count int
	CountOnNode int
	EnterGens int

	SurvivalPercent float64
	StartGenCountMin int
	StartGenCountMax int

	MutationShockSteps int
	MutationShockRatio float64

	FinishSteps int
}

func GetAppConfig() *Config {
	return &config
}

var config Config

func init () {
	config = Config{
		Count: 8000,
		CountOnNode: 2000,
		EnterGens: 7,

		SurvivalPercent: 0.25,
		StartGenCountMin: 1,
		StartGenCountMax: 100,

		MutationShockSteps: 5,
		MutationShockRatio: 1.01,

		FinishSteps: 1000,
	}
}
