package calculation

import (
	"math"
)

func getMean(distr map[int]float64) float64 {
	return 0
}

func getDispersion(distr map[int]float64) float64 {

	return 0
}

func getStdDeviation(distr map[int]float64) float64 {
	return 0
}

func getDistribution(dices map[string]int, chance float64, shots int, bonus int) map[int]float64 {
	distr := make(map[int]float64)

	if chance == 1 {
		distr[0] = float64(1)
	} else if chance < 1 {
		distr[0] = math.Pow(chance, float64(shots))
	}

	return distr
}
