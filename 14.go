package main

import (
	"fmt"
	"math"
	"strings"
)

func adventDay14A(path string) {
	lines := readLines(path)
	state := lines[0]


	rules := make(map[string]string)
	for _,line := range lines[1:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	const steps = 10

	var nextString strings.Builder
	for step := 0; step < steps; step++ {
		for i := 0; i < len(state)-1; i++ {
			nextString.WriteString(state[i:i+1])
			rule, ok := rules[state[i:i+2]]
			if ok {
				nextString.WriteString(rule)
			}
		}
		nextString.WriteString(state[len(state)-1:])
		state = nextString.String()
		nextString.Reset()

		//fmt.Printf("string is %s after %d steps\n", state, step+1)
		//fmt.Printf("%+v\n", stateFromLine(state))
	}
	freq := make(map[rune]int)
	for _,r := range state {
		freq[r]++
	}
	max := 0
	min := math.MaxInt
	for _,val := range freq {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Printf("max - min = %d\n", max-min)
}

func minMax(state map[string]int, lastChar rune) int {
	freq := make(map[rune]int)
	for pattern,n := range state {
		freq[rune(pattern[0])] += n
		//freq[rune(pattern[1])] += n
	}
	freq[lastChar]++
	//for r, n := range freq {
	//	fmt.Printf("%c:%d ", r, n)
	//}
	//fmt.Printf("\n")
	max := 0
	min := math.MaxInt
	for _,val := range freq {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max-min
}

func stateFromLine(line string) map[string]int {
	state := make(map[string]int)
	for i := 0; i < len(line)-1; i++ {
		state[line[i:i+2]] += 1
	}
	return state
}

func adventDay14B(path string) {
	lines := readLines(path)
	firstLine := lines[0]


	rules := make(map[string]string)
	for _,line := range lines[1:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	const steps = 40
	state := stateFromLine(firstLine)

	for step := 0; step < steps; step++ {
		nextState := make(map[string]int)
		for pattern,n := range state {
			rule  := rules[pattern]
			nextState[pattern[0:1]+rule] += n
			nextState[rule+pattern[1:2]] += n
		}
		state = nextState
		//fmt.Printf("%+v\n", state)
		//fmt.Printf("max - min = %d\n", minMax(state))
		//fmt.Printf("string is %d long after %d steps\n", len(state), step+1)
	}
	fmt.Printf("max - min = %d\n", minMax(state, rune(firstLine[len(firstLine)-1])))

}