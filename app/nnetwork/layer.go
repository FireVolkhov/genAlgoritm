package nnetwork

type layer struct {
	Neurons []neuron
}

func runLayer (targetLayer, preLayer layer) layer {
	for neuronIndex, neuron := range targetLayer.Neurons {
		targetLayer.Neurons[neuronIndex] = runNeuron(neuron, preLayer)
	}

	return targetLayer
}
