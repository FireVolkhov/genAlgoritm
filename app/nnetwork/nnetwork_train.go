package nnetwork

import "math"
import "log"
import "fmt"
import "time"
import "../../app/core"
import "io/ioutil"

func (this *NNetwork) train (data []TrainData, config *Config) {
	dataLen := core.RoundToInt(config.TestDataPercent * float64(len(data)))
	testData := data[dataLen:]
	data = data[:dataLen]

//	dataLen := len(data)
	layerLen := len(this.Layers)

	currentError := math.MaxFloat64
	lastError := float64(0)
	epochNumber := int(0)

	if (config.BatchSize < 1 || config.BatchSize > dataLen) {
		config.BatchSize = dataLen
	}

	log.Println("Start learning...")

	for {
		lastError = currentError
		dtStart := time.Now()

		//process data set
		currentIndex := 0

		for currentIndex < dataLen {
			//accumulated error for batch, for weights and biases
			nablaWeights := make([][][]float64, len(this.Layers))
			nablaBiases := make([][]float64, len(this.Layers))

			//process one batch
			for inBatchIndex := currentIndex; inBatchIndex < currentIndex + config.BatchSize && inBatchIndex < dataLen; inBatchIndex++ {
				//forward pass
				realOutput := this.Run(data[inBatchIndex].Input)

				//backward pass, error propagation
				//last layer
				lastLayerIndex := layerLen - 1
				lastLayerNeuronsLen := len(this.Layers[lastLayerIndex].Neurons)
				nablaWeights[lastLayerIndex] = make([][]float64, lastLayerNeuronsLen)
				nablaBiases[lastLayerIndex] = make([]float64, lastLayerNeuronsLen)

				for neuronIndex, neuron := range this.Layers[lastLayerIndex].Neurons {
					dEdz := halfSquaredEuclidianDistancePartialDerivaitveByV2Index(data[inBatchIndex].Output, realOutput, neuronIndex) * hyperbolicTangensFirstDerivative(neuron.LastInput)

					this.Layers[lastLayerIndex].Neurons[neuronIndex].DEdz = dEdz
					nablaBiases[lastLayerIndex][neuronIndex] = config.LearningRate * dEdz;
					nablaWeights[lastLayerIndex][neuronIndex] = make([]float64, len(neuron.Weights))

					for weightIndex, weight := range neuron.Weights {
						nablaWeights[lastLayerIndex][neuronIndex][weightIndex] = calcNablaWeights(dEdz, this.Layers[lastLayerIndex - 1].Neurons[weightIndex].Output, weight, dataLen, config)
					}
				}

				//hidden layers
				for hiddenLayerIndex := layerLen - 2; hiddenLayerIndex >= 0; hiddenLayerIndex-- {
					hiddenLayer := this.Layers[hiddenLayerIndex]
					hiddenLayerNeuronsLen := len(hiddenLayer.Neurons)

					nablaWeights[hiddenLayerIndex] = make([][]float64, hiddenLayerNeuronsLen)
					nablaBiases[hiddenLayerIndex] = make([]float64, hiddenLayerNeuronsLen)

					for neuronIndex, neuron := range hiddenLayer.Neurons {
						dEdz := float64(0)

						for _, preNeuron := range this.Layers[hiddenLayerIndex + 1].Neurons {
							dEdz += preNeuron.Weights[neuronIndex] * preNeuron.DEdz
						}

						dEdz *= hyperbolicTangensFirstDerivative(neuron.LastInput)
						nablaBiases[hiddenLayerIndex][neuronIndex] = config.LearningRate * dEdz
						nablaWeights[hiddenLayerIndex][neuronIndex] = make([]float64, len(neuron.Weights))

						for weightIndex, weight := range neuron.Weights {
							nablaWeights[hiddenLayerIndex][neuronIndex][weightIndex] = calcNablaWeights(dEdz, this.Layers[hiddenLayerIndex - 1].Neurons[weightIndex].Output, weight, dataLen, config)
						}

						this.Layers[hiddenLayerIndex].Neurons[neuronIndex].DEdz = dEdz
					}
				}
			}

			//update Weights and bias
			for layerIndex := 0; layerIndex < layerLen; layerIndex++ {
				for neuronIndex := 0; neuronIndex < len(this.Layers[layerIndex].Neurons); neuronIndex++ {
					this.Layers[layerIndex].Neurons[neuronIndex].Bias -= nablaBiases[layerIndex][neuronIndex];

					for weightIndex := 0; weightIndex < len(this.Layers[layerIndex].Neurons[neuronIndex].Weights); weightIndex++ {
						this.Layers[layerIndex].Neurons[neuronIndex].Weights[weightIndex] -= nablaWeights[layerIndex][neuronIndex][weightIndex];
					}
				}
			}

			currentIndex += config.BatchSize
		}

		//recalculating error on all data
		currentError = 0
		maxMouseGoodPercent := float64(0)
		mouseGoodPercent := float64(0)
		mouseGoonSteps := 0
		fullResult := float64(1)

		for _, step := range data {
			realOutput := this.Run(step.Input)
			currentError += halfSquaredEuclidianDistance(step.Output, realOutput)

			futureDay := step.Step
			futureDayDif := futureDay.Close - futureDay.Open
			futureDayPercent := math.Abs(futureDayDif) / futureDay.Open

			marketBay := 0 < futureDayDif
			marketWait := futureDayDif == 0
			marketSell := futureDayDif < 0

			mouseBay := 0 < realOutput[0]
			//mouseWait := realOutput == 0
			mouseSell := realOutput[0] < 0

			if ((mouseBay && marketBay) ||
				(mouseSell && marketSell)) {
				fullResult = fullResult * (1 + futureDayPercent)
				mouseGoonSteps++
			} else if (
				(mouseBay && (marketSell || marketWait)) ||
				(mouseSell && (marketBay || marketWait))) {
				fullResult = fullResult * (1 - futureDayPercent)
			}
		}

		mouseGoodPercent = float64(mouseGoonSteps) / float64(dataLen)
		percentYear := math.Pow(fullResult, 1 / (float64(dataLen) / 365)) - 1
		currentError *= 1 / float64(dataLen)
		epochNumber++


		// Test
		testCurrentError := float64(0)
		testMouseGoodPercent := float64(0)
		testMouseGoonSteps := 0
		testFullResult := float64(1)

		for _, step := range testData {
			realOutput := this.Run(step.Input)
			testCurrentError += halfSquaredEuclidianDistance(step.Output, realOutput)

			futureDay := step.Step
			futureDayDif := futureDay.Close - futureDay.Open
			futureDayPercent := math.Abs(futureDayDif) / futureDay.Open

			marketBay := 0 < futureDayDif
			marketWait := futureDayDif == 0
			marketSell := futureDayDif < 0

			mouseBay := 0 < realOutput[0]
			//mouseWait := realOutput == 0
			mouseSell := realOutput[0] < 0

			if ((mouseBay && marketBay) ||
			(mouseSell && marketSell)) {
				testFullResult = testFullResult * (1 + futureDayPercent)
				testMouseGoonSteps++
			} else if (
			(mouseBay && (marketSell || marketWait)) ||
			(mouseSell && (marketBay || marketWait))) {
				testFullResult = testFullResult * (1 - futureDayPercent)
			}
		}

		testMouseGoodPercent = float64(testMouseGoonSteps) / float64(len(testData))

		//logs
		if (epochNumber % 10 == 0) {
			if (maxMouseGoodPercent < mouseGoodPercent) {
				maxMouseGoodPercent = mouseGoodPercent

				network := this.ToBytes()
				ioutil.WriteFile("./result/" + fmt.Sprintf("%d-%2.2f%s", epochNumber, maxMouseGoodPercent * 100, "%") + ".net", network, 0644)
			}

			log.Println(
				fmt.Sprintf("%4.0f ", float64(epochNumber)) +
				fmt.Sprintf("E: %4.6f ", currentError) +
				fmt.Sprintf("FR: %10.0f%s ", fullResult * 100, "%") +
				fmt.Sprintf("YR: %2.2f%s ", percentYear * 100, "%") +
				fmt.Sprintf("GS: %2.6f%s ", mouseGoodPercent * 100, "%") +
				"TEST >>> " +
				fmt.Sprintf("E: %4.6f ", testCurrentError) +
				fmt.Sprintf("TFR: %10.0f%s ", testFullResult * 100, "%") +
				fmt.Sprintf("TGS: %2.6f%s ", testMouseGoodPercent * 100, "%") +
				fmt.Sprintf("%s", time.Now().Sub(dtStart)))
		}

		if (!(epochNumber < config.MaxEpoches &&
			currentError > config.MinError &&
			math.Abs(currentError - lastError) > config.MinErrorChange)) {
			break
		}
	}

	log.Println("End")
}

func calcNablaWeights (dEdz, output, weight float64, dataLen int, config *Config) float64 {
	return config.LearningRate * (dEdz * output + config.RegularizationFactor * weight / float64(dataLen))
}
