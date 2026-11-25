package main

import (
	"fmt"
)

func main() {
	dices := map[string]int{
		"d4":   1,
		"d6":   0,
		"d8":   0,
		"d10":  0,
		"d20":  0,
		"d100": 0,
	}

	var chance float64 = 1
	shots, bonus := 1, -4
	ans := getDistribution(dices, chance, shots, bonus)
	for i := range ans {
		fmt.Println(i, ": ", ans[i]*100)
	}
	fmt.Println("Mean: ", getMean(ans))
	fmt.Println("Dispersion ", getDispersion(ans))
	fmt.Println("Standart deviation: ", getStdDeviation(ans))
}
