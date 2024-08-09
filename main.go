package main

import (
	"fmt"

	"github.com/saracha-06422/band-protocal-interview/problem"
)

func main() {
	fmt.Println("Problem 1: Boss Baby's Revenge")

	testCasesP1 := []string{
		"SRSSRRR",  // Output: Good boy
		"RSSRR",    // Output: Bad boy
		"SSSRRRRS", // Output: Bad boy
		"SRRSSR",   // Output: Bad boy
		"SSRSRRR",  // Output: Good boy
	}

	for _, test := range testCasesP1 {
		fmt.Println(problem.BossBabyRevenge(test))
	}

	fmt.Println("Problem 2: Superman's Chicken Rescue")
	type SupermanChickenRescue struct {
		NumOfChicken      int
		Distance          int
		ChickensPositions []int
	}

	var testCasesP2 = []SupermanChickenRescue{
		{
			NumOfChicken:      5,
			Distance:          5,
			ChickensPositions: []int{2, 5, 10, 12, 15},
		},
		{
			NumOfChicken:      6,
			Distance:          10,
			ChickensPositions: []int{1, 11, 30, 34, 35, 37},
		},
	}

	for _, test := range testCasesP2 {
		fmt.Println(problem.MaxChickensCovered(test.ChickensPositions, test.NumOfChicken, test.Distance))
	}

}
