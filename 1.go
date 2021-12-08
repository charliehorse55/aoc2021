package main

import "fmt"

func adventDay1A(path string) {
	depths := toInts(readLines(path))

	lastDepth := depths[0]
	increases := 0
	for _,depth := range depths[1:] {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}
	fmt.Printf("increased %d times\n", increases)
}

func adventDay1B(path string) {
	depths := toInts(readLines(path))

	var lastDepths [3]int
	copy(lastDepths[:], depths[:3])
	increases := 0
	for _,depth := range depths[3:] {
		oldSum := lastDepths[0] + lastDepths[1] + lastDepths[2]
		copy(lastDepths[:], lastDepths[1:])
		lastDepths[2] = depth
		newSum := lastDepths[0] + lastDepths[1] + lastDepths[2]
		if newSum > oldSum {
			increases++
		}
	}
	fmt.Printf("increased %d times\n", increases)
}
