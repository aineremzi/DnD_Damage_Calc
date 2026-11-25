package main

import (
	"math"
)

func getMean(distr []float64) float64 {
	M := 0.0
	for i := range distr {
		M += float64(i) * distr[i]
	}
	return M
}

func getDispersion(distr []float64) float64 {
	M := getMean(distr)
	M2 := 0.0
	for i := range distr {
		M2 += math.Pow(float64(i), 2) * distr[i]
	}
	return M2 - math.Pow(M, 2)
}

func getStdDeviation(distr []float64) float64 {
	return math.Sqrt(getDispersion(distr))
}

func calculateDice(distr []float64, n int, dice int, maxDmg int) []float64 {
	for i := 0; i < n; i++ {
		next := make([]float64, maxDmg+1)
		for s := 0; s <= maxDmg; s++ {
			if distr[s] == 0 {
				continue
			}
			for f := 1; f <= dice; f++ {
				if s+f <= maxDmg {
					next[s+f] += distr[s] / float64(dice)
				}
			}
		}
		distr = next
	}
	return distr
}

func calculateShots(distr []float64, n int, chance float64) []float64 {
	tempDistr := make([]float64, cap(distr)*n-n+1)
	next := make([]float64, cap(tempDistr))
	tempDistr[0] = 1
	for i := 0; i < n; i++ {
		for j := range next {
			next[j] = 0
		}
		for j := 0; j < len(next); j++ {
			next[j] += (1 - chance) * tempDistr[j]
			for k := 0; k < len(distr); k++ {
				if j+k < cap(tempDistr) {
					next[j+k] += chance * distr[k] * tempDistr[j]
				}
			}
		}
		copy(tempDistr, next)
	}
	return tempDistr
}

func getDistribution(dices map[string]int, chance float64, shots int, bonus int) []float64 {
	c := 0
	for _, n := range dices {
		if n == 0 {
			c += 1
		}
	}
	if c == len(dices) || shots == 0 || chance == 0 {
		return []float64{1}
	}

	var maxDmg int
	for dice, number := range dices {
		switch dice {
		case "d4":
			maxDmg += 4 * number
		case "d6":
			maxDmg += 6 * number
		case "d8":
			maxDmg += 8 * number
		case "d10":
			maxDmg += 10 * number
		case "d12":
			maxDmg += 12 * number
		case "d20":
			maxDmg += 20 * number
		case "d100":
			maxDmg += 100 * number
		}
	}

	var distr []float64
	if bonus > 0 {
		distr = make([]float64, maxDmg+bonus+1)
	} else {
		distr = make([]float64, maxDmg+1)
	}
	distr[0] = 1.0

	for dice, number := range dices {
		switch dice {
		case "d4":
			distr = calculateDice(distr, number, 4, maxDmg)
		case "d6":
			distr = calculateDice(distr, number, 6, maxDmg)
		case "d8":
			distr = calculateDice(distr, number, 8, maxDmg)
		case "d10":
			distr = calculateDice(distr, number, 10, maxDmg)
		case "d12":
			distr = calculateDice(distr, number, 12, maxDmg)
		case "d20":
			distr = calculateDice(distr, number, 20, maxDmg)
		case "d100":
			distr = calculateDice(distr, number, 100, maxDmg)
		}
	}

	if bonus < 0 {
		for i := 0; i < -bonus; i++ {
			distr[i+1] += distr[i]
		}
		distr = distr[-bonus:]
	} else {
		copy(distr[bonus:], distr)
	}

	return calculateShots(distr, shots, chance)
}
