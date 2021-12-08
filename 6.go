package main

import (
	"fmt"
	"strings"
)

func simulateDays(state [9]int, days int) int {
	for day := 0; day < days; day++ {
		resetting := state[0]
		for i := 0; i < 8; i++ {
			state[i] = state[i+1]
		}
		state[8] = resetting
		state[6] += resetting
	}
	total := 0
	for _,n := range state {
		total += n
	}
	return total
}

func adventDay6A(path string) {
	lines := readLines(path)
	lanterns := toInts(strings.Split(lines[0], ","))


	var state [9]int
	for _,lantern := range lanterns {
		state[lantern]++
	}

	const days = 80
	total := simulateDays(state, days)
	fmt.Printf("After %d days there are %d fish\n", days, total)

}

func adventDay6B(path string) {
	lines := readLines(path)
	lanterns := toInts(strings.Split(lines[0], ","))


	var state [9]int
	for _,lantern := range lanterns {
		state[lantern]++
	}

	const days = 256
	total := simulateDays(state, days)
	fmt.Printf("After %d days there are %d fish\n", days, total)

}
