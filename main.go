package main

import (
	"fmt"

	"github.com/saracha-06422/band-protocal-interview/problem"
)

func main() {
	fmt.Println("band protocal interview")

	// Example test cases
	testCases := []string{
		"SRSSRRR",  // Output: Good boy
		"RSSRR",    // Output: Bad boy
		"SSSRRRRS", // Output: Bad boy
		"SRRSSR",   // Output: Bad boy
		"SSRSRRR",  // Output: Good boy
	}

	for _, test := range testCases {
		fmt.Println(problem.BossBabyRevenge(test))
	}
}
