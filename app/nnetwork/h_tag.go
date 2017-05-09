package nnetwork

import "math"
import "log"

var hyperbolicTangensAlpha float64 = 1;

func hyperbolicTangens(x float64) float64 {
	result := math.Tanh(hyperbolicTangensAlpha * x)

	if (math.IsNaN(result) || math.IsInf(result, 0)) {
		log.Panic("hyperbolicTangens Detect >", x, result)
	}

	return result
}

func hyperbolicTangensFirstDerivative(x float64) float64 {
	tan := math.Tanh(hyperbolicTangensAlpha * x)
	result := hyperbolicTangensAlpha * (1 - tan * tan);

	if (math.IsNaN(result) || math.IsInf(result, 0)) {
		log.Panic("hyperbolicTangensFirstDerivative Detect >", x, result)
	}

	return result
}
