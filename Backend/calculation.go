package main

import (
	"fmt"
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

func updateValuesD4(dices map[string]int, values []int, bonus int) []int {
	for i := 0; i < dices["d4"]; i++ {
		temp := values
		for j := 1; j < 5; j++ {
			for k := 0; k < len(temp); k++ {
				value := max(j+temp[k]+bonus, 0)
				values = append(values, value)
			}
		}
		values = values[len(temp):]
	}
	return values
}
func updateValuesD6(dices map[string]int, values []int, bonus int) []int {
	for i := 0; i < dices["d6"]; i++ {
		temp := values
		for j := 1; j < 7; j++ {
			for k := 0; k < len(temp); k++ {
				value := max(j+temp[k]+bonus, 0)
				values = append(values, value)
			}
		}
		values = values[len(temp):]
	}
	return values
}
func updateValuesD8(dices map[string]int, values []int, bonus int) []int {
	for i := 0; i < dices["d8"]; i++ {
		temp := values
		for j := 1; j < 9; j++ {
			for k := 0; k < len(temp); k++ {
				value := max(j+temp[k]+bonus, 0)
				values = append(values, value)
			}
		}
		values = values[len(temp):]
	}
	return values
}
func updateValuesD10(dices map[string]int, values []int, bonus int) []int {
	for i := 0; i < dices["d10"]; i++ {
		temp := values
		for j := 1; j < 11; j++ {
			for k := 0; k < len(temp); k++ {
				value := max(j+temp[k]+bonus, 0)
				values = append(values, value)
			}
		}
		values = values[len(temp):]
	}
	return values
}
func updateValuesD12(dices map[string]int, values []int, bonus int) []int {
	for i := 0; i < dices["d12"]; i++ {
		temp := values
		for j := 1; j < 12; j++ {
			for k := 0; k < len(temp); k++ {
				value := max(j+temp[k]+bonus, 0)
				values = append(values, value)
			}
		}
		values = values[len(temp):]
	}
	return values
}
func updateValuesD20(dices map[string]int, values []int, bonus int) []int {
	for i := 0; i < dices["d20"]; i++ {
		temp := values
		for j := 1; j < 21; j++ {
			for k := 0; k < len(temp); k++ {
				value := max(j+temp[k]+bonus, 0)
				values = append(values, value)
			}
		}
		values = values[len(temp):]
		fmt.Println(i)
	}
	return values
}

func getDistribution(dices map[string]int, chance float64, shots int, bonus int) map[int]float64 {
	distr := make(map[int]float64)
	values := []int{0}

	distr[0] = math.Pow(1-chance, float64(shots))

	values = updateValuesD4(dices, values, bonus)
	values = updateValuesD6(dices, values, bonus)
	values = updateValuesD8(dices, values, bonus)
	values = updateValuesD10(dices, values, bonus)
	values = updateValuesD12(dices, values, bonus)
	values = updateValuesD20(dices, values, bonus)

	n := len(values)
	for i := 0; i < n; i++ {
		distr[values[i]] += float64(1) / float64(n) * chance
	}

	return distr
}
