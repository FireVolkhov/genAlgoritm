package nnetwork

func halfSquaredEuclidianDistance (v1, v2 []float64) float64 {
	d := float64(0)

	for i := 0; i < len(v1); i++ {
		d += (v1[i] - v2[i]) * (v1[i] - v2[i])
	}

	return 0.5 * d
}

func halfSquaredEuclidianDistancePartialDerivaitveByV2Index (v1, v2 []float64, v2Index int) float64 {
	return v2[v2Index] - v1[v2Index]
}
