package main

import (
	"fmt"
)

func main() {
	dices := map[string]int{
		"d4":  0,
		"d6":  0,
		"d8":  0,
		"d10": 1,
		"d20": 0,
	}

	var chance float64 = 0.65
	shots, bonus := 3, 0
	ans := getDistribution(dices, chance, shots, bonus)
	for i := range ans {
		fmt.Println(i, ": ", ans[i])
	}

}
