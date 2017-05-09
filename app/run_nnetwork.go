package main

import "../app/nnetwork"
import "../app/test/data"
//import "math"

func main () {
	config := nnetwork.Config{
		LearningRate: 0.00001,
		RegularizationFactor: 0.001,
		BatchSize: 1,
		MaxEpoches: 8000,
		MinError: 0,
		MinErrorChange: 0,
		TestDataPercent: 0.2,
	}
	windowSize := 30
	iterator := data.NewTableIterator(windowSize + 1)

	data := make([]nnetwork.TrainData, iterator.GetLen())

	stepIndex := 0;

	for (!iterator.IsFinished) {
		slice := iterator.Next()
		diffSlice := make([]float64, len(slice))

		for index, item := range slice {
			diffSlice[index] = item.Dif / item.Open
//			diffSlice[index] = item.Dif
		}

		input := diffSlice[:windowSize]

//		for index := range input {
//			input[index] = slice[index].Dif / slice[index].Open
//		}

		output := diffSlice[windowSize:]

//		output[0] = 1 / output[0]
//
//		if (1 < output[0]) {
//			output[0] = 1
//		}
//		if (output[0] < -1) {
//			output[0] = -1
//		}

//		output := make([]float64, 2)
//
//		output[0] = 1 / diffSlice[7]
//
//		if (0 < output[0]) {
//			output[1] = 0
//		} else {
//			output[1] = math.Abs(output[0])
//			output[0] = 0
//		}
//
//		output[0] = math.Min(output[0], 1)
//		output[1] = math.Min(output[1], 1)

//		data[stepIndex] = nnetwork.TrainData{
//			Input: nnetwork.NormalizeInput(input),
//			Output: nnetwork.NormalizeInput(output),
//			Step: *slice[windowSize],
//		}

		data[stepIndex] = nnetwork.TrainData{
			Input: input,
			Output: output,
			Step: *slice[windowSize],
		}

		stepIndex++
	}

	network := nnetwork.NewNNetwork(windowSize, windowSize * 2, 3, 1)
	network.Train(data, &config)
}
