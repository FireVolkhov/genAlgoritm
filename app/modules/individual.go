package modules

import (
	"../../app/genome"
	"../../app/core"
	"strings"
	"strconv"
)

type Individual struct {
	body []*genType
	calcCache map[int]float64
}

func NewIndividual (enterGensCount, gensCount int) *Individual {
	individual := &Individual{
		body: make([]*genType, enterGensCount),
	}

	for genIndex := range individual.body {
		individual.body[genIndex] = &genType{
			Name: "ENTER",
			Args: make([]int, 0),
			IsEnterGen: true,
		}
	}

	for i := 0; i < gensCount; i++ {
		individual.AddGen(genome.GetRandomGen())
	}

	return individual
}

func (this *Individual) GetGensCount() int {
	var result int = 0

	for _, gen := range this.body {
		if (!gen.IsEnterGen) {
			result++
		}
	}

	return result
}

func (this *Individual) AddGen (gen *genome.Gen) {
	genAsClassFormat := genToGenType(gen)

	for argIndex := range genAsClassFormat.Args {
		genAsClassFormat.Args[argIndex] = core.RandomInt(0, len(this.body) - 1)
	}

	this.body = append(this.body, genAsClassFormat)
}

func (this *Individual) Mutation () *Individual {
	clone := this.clone()
	mutationCount := core.RandomInt(1, clone.GetGensCount())
	mutationRule := core.RandomInt(0, 2)

	for i := 0; i < mutationCount; i++ {
		switch mutationRule {
		case 0:
			clone.mutationAddGen()
		case 1:
			clone.mutationModifyGen()
		case 2:
			clone.mutationRemoveGen()
		default:
			panic("Not have exits case")
		}
	}

	return clone
}

func (this *Individual) Run (args []float64) float64 {
	this.calcCache = make(map[int]float64)

	for argIndex, arg := range args {
		this.calcCache[argIndex] = arg
	}

	genIndex := len(this.body) - 1;
	result := this.calcGen(genIndex, this.body[genIndex]);

	this.calcCache = nil

	return result
}

func (this *Individual) ToString () string {
	arrayLen := len(this.body)

	if (0 < arrayLen) {
		result := make([]string, arrayLen)

		for genIndex, gen := range this.body {
			line := make([]string, 1 + len(gen.Args))
			line[0] = strconv.Itoa(genIndex) + " " + gen.Name

			for argIndex, arg := range gen.Args {
				line[argIndex + 1] = strconv.Itoa(arg)
			}

			result[genIndex] = strings.Join(line, " ")
		}

		return "\n" + strings.Join(result, "\n")

	} else {
		return ""
	}
}

type genType struct {
	Name string
	Args []int

	IsEnterGen bool
}

func genToGenType (gen *genome.Gen) *genType {
	return &genType{
		Name: gen.Name,
		Args: make([]int, gen.EnterGens),
		IsEnterGen: false,
	}
}

func (this *Individual) calcGen (genIndex int, gen *genType) float64 {
	valueFromCache, ok := this.calcCache[genIndex]

	if (ok) {
		return valueFromCache

	} else {
		genArgsResult := make([]float64, len(gen.Args))

		for genArgIndex, targetArgGenIndex := range gen.Args {
			argGen := this.body[targetArgGenIndex]
			genArgsResult[genArgIndex] = this.calcGen(targetArgGenIndex, argGen)
		}

		result := genome.Calc(gen.Name, genArgsResult);
		this.calcCache[genIndex] = result
		return result
	}
}

func (this *Individual) clone () *Individual {
	clone := &Individual{
		body: make([]*genType, len(this.body)),
		calcCache: make(map[int]float64),
	}

	for genIndex, gen := range this.body {
		clone.body[genIndex] = &genType{
			Name: gen.Name,
			Args: append([]int(nil), gen.Args...),
			IsEnterGen: gen.IsEnterGen,
		}
	}

	return clone
}

func (this *Individual) countEnterGens () int {
	return len(this.body) - this.GetGensCount()
}

func (this *Individual) mutationAddGen () {
	newGen := genToGenType(genome.GetRandomGen())
	countEnterGens := this.countEnterGens()
	position := core.RandomInt(countEnterGens, len(this.body) - 1);

	for argIndex := range newGen.Args {
		newGen.Args[argIndex] = core.RandomInt(0, position - 1)
	}

	// Insert array
	this.body = append(this.body, newGen)
	copy(this.body[position + 1:], this.body[position:])
	this.body[position] = newGen

	// Filter array
	body := make([]*genType, len(this.body))
	copy(body, this.body)
	body = body[position + 1:]

	filteredBody := make([]*genType, 0)
	for _, gen := range body {
		if (0 < len(gen.Args)) {
			filteredBody = append(filteredBody, gen)
		}
	}

	if (0 < len(filteredBody)) {
		target := filteredBody[core.RandomInt(0, len(filteredBody) - 1)]
		targetPos := core.RandomInt(0, len(target.Args) - 1)
		target.Args[targetPos] = position
	}
}

func (this *Individual) mutationModifyGen () {
	changeType := core.RandomBool()
	changeArgs := core.RandomBool()

	countEnterGens := this.countEnterGens()
	position := core.RandomInt(countEnterGens, len(this.body) - 1)
	target := this.body[position]

	if (!changeType && !changeArgs) {
		changeType = true
	}

	if (changeType) {
		newGen := genome.GetRandomGen()
		needGens := newGen.EnterGens
		hasGens := len(target.Args)

		target.Name = newGen.Name;

		if (needGens < hasGens) {
			target.Args = target.Args[:needGens]

		} else if (hasGens < needGens) {
			for len(target.Args) < needGens {
				target.Args = append(target.Args, core.RandomInt(0, position - 1))
			}
		} else {
			// Do nothing
		}
	}

	if (changeArgs && 0 < len(target.Args)) {
		targetPos := core.RandomInt(0, len(target.Args) - 1);
		newArg := core.RandomInt(0, position - 1);

		target.Args[targetPos] = newArg;
	}
}

func (this *Individual) mutationRemoveGen () {
	countEnterGens := this.countEnterGens()
	position := core.RandomInt(countEnterGens, len(this.body) - 1)

	if (len(this.body) > (countEnterGens + 1)) {
		copy(this.body[position:], this.body[position + 1:])
		this.body = this.body[:len(this.body) - 1]

		for genIndex, gen := range this.body {
			if (position <= genIndex) {
				for argIndex, arg := range gen.Args {
					gen.Args[argIndex] = core.MaxInt(0, arg - 1)
				}
			}
		}
	}
}
