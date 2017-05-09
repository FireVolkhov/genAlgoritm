package nnetwork

import "math"
import "log"

type neuron struct {
	Output float64
	Weights []float64
	LastInput float64

	Bias float64
	DEdz float64
}

func runNeuron(targetNeuron neuron, preLayer layer) neuron {
	sum := float64(0)

	for weightIndex, weight := range targetNeuron.Weights {
		sum += preLayer.Neurons[weightIndex].Output * weight
	}

	targetNeuron.LastInput = sum
	targetNeuron.Output = hyperbolicTangens(sum)

	if (math.IsNaN(targetNeuron.Output) || math.IsInf(targetNeuron.Output, 0)) {
		log.Panic("runNeuron Detect >", targetNeuron.Output)
	}

	return targetNeuron
}
