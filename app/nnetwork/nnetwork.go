package nnetwork

import "../../app/core"
import "../../app/test/data"
import "math"
import "log"

type Config struct {
	LearningRate float64
	RegularizationFactor float64
	BatchSize int
	MaxEpoches int
	MinError float64
	MinErrorChange float64
	TestDataPercent float64
}

type TrainData struct {
	Input []float64
	Output []float64
	Step data.TableRow
}

type NNetwork struct {
	Layers []layer
}

func NewNNetwork (inputNeurons, hiddenNeurons, countLayers, outputNeurons int) *NNetwork {
	return newNNetwork(inputNeurons, hiddenNeurons, countLayers, outputNeurons)
}

func (this *NNetwork) Run (input []float64) []float64 {
	return this.run(input)
}

func (this *NNetwork) Train (data []TrainData, config *Config) {
	this.train(data, config)
}

func (this *NNetwork) ToBytes () []byte {
	return this.toBytes()
}

func NNetworkFromBytes (bytes []byte) NNetwork {
	return fromBytes(bytes)
}


// --- PRIVATE ---------------------------------------------------------------------------------------------------------
func newNNetwork (inputNeurons, hiddenNeurons, countLayers, outputNeurons int) *NNetwork {
	nnetwork := NNetwork{
		Layers: make([]layer, countLayers + 2),
	}

	for layerIndex := range nnetwork.Layers {
		switch layerIndex {

		// Input layer
		case 0:
			nnetwork.Layers[layerIndex] = getRandomLayer(inputNeurons, 0)
			break

		// Second layer
		case 1:
			nnetwork.Layers[layerIndex] = getRandomLayer(hiddenNeurons, inputNeurons)
			break

		// Last layer
		case len(nnetwork.Layers) - 1:
			nnetwork.Layers[layerIndex] = getRandomLayer(outputNeurons, hiddenNeurons)
			break

		// layer
		default:
			nnetwork.Layers[layerIndex] = getRandomLayer(hiddenNeurons, hiddenNeurons)
		}
	}

	return &nnetwork
}

func (this *NNetwork) run (input []float64) []float64 {
	this.clearState()

//	normalizedInput := normalizeInput(input)
	inputLayer := 0
	outputLayer := len(this.Layers) - 1
	result := make([]float64, len(this.Layers[outputLayer].Neurons))

	for neuronIndex, value := range input {
		this.Layers[inputLayer].Neurons[neuronIndex].Output = value
	}

	for layerIndex, layer := range this.Layers {
		if (inputLayer < layerIndex) {
			this.Layers[layerIndex] = runLayer(layer, this.Layers[layerIndex - 1])
		}
	}

	for neuronIndex, neuron := range this.Layers[outputLayer].Neurons {
		result[neuronIndex] = neuron.Output
	}

	return result
}

func (this *NNetwork) clearState() {
	for layerIndex := range this.Layers {
		for neuronIndex := range this.Layers[layerIndex].Neurons {
			this.Layers[layerIndex].Neurons[neuronIndex].Output = float64(0)
		}
	}
}

func getRandomLayer (countNeurons, countWeight int) layer {
	return layer{
		Neurons: getRandomNeurons(countNeurons, countWeight),
	}
}

func getRandomNeurons (countNeurons, countWeight int) []neuron {
	neurons := make([]neuron, countNeurons)

	for neuronIndex := range neurons {
		neurons[neuronIndex] = neuron{
			Weights: getRandomWeights(countWeight),
		}
	}

	return neurons
}

func getRandomWeights (length int) []float64 {
	slice := make([]float64, length)

	for weightIndex := range slice {
		slice[weightIndex] = (core.RandomFloat64() * 2) - 1
	}

	return slice
}

func NormalizeInput (input []float64) []float64 {
	for index, value := range input {
		input[index] = 1 / value;

		if (1 < input[index]) {
			input[index] = 1
		}

		if (input[index] < -1) {
			input[index] = -1
		}

		if (math.IsInf(input[index], 0)) {
			log.Panic("normalizeInput Detect >", input[index], value)
		}
	}


	return input
}
