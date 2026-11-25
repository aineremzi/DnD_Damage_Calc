package main

import (
	"fmt"
)

func main() {
	dices := map[string]int{
		"d4":  0,
		"d6":  0,
		"d8":  0,
		"d10": 0,
		"d20": 20,
	}

	var chance float64 = 1
	shots, bonus := 1, 0
	ans := getDistribution(dices, chance, shots, bonus)
	fmt.Println(ans)
}
