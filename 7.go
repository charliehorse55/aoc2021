package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func intABS(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gas(crabs []int, pos int) int {
	g := 0
	for _,crab := range crabs {
		g += intABS(crab-pos)
	}
	return g
}

func median(x []int) int {
	y := make([]int, len(x))
	copy(y, x)
	sort.Ints(y)
	return y[len(y)/2]
}

func adventDay7A(path string) {
	lines := readLines(path)
	crabs := toInts(strings.Split(lines[0], ","))
	total := 0
	for _,crab := range crabs {
		total += crab
	}
	//avgf := float64(total)/float64(len(crabs))
	avg := median(crabs)



	fmt.Printf("avg is %d, requires %d gas (-1 = %d, +1 = %d)\n", avg, gas(crabs, avg), gas(crabs, avg-1), gas(crabs, avg+1))
}

func adventDay7B(path string) {
	lines := readLines(path)
	crabs := toInts(strings.Split(lines[0], ","))
	sort.Ints(crabs)

	low := crabs[0]
	high := crabs[len(crabs)-1]

	fmt.Printf("low = %d, high = %d\n", low, high)

	dist := high-low+1
	rightMove := make([]int, dist)
	leftMove := make([]int, dist)

	n := 0
	stepCost := 0
	totalCost := 0
	pos := 0
	for i := low; i <= high; i++ {
		totalCost += stepCost
		rightMove[i-low] = totalCost
		for pos < len(crabs) && crabs[pos] == i {
			n++
			pos++
		}
		stepCost += n
	}

	n2 := 0
	stepCost2 := 0
	totalCost2 := 0
	pos2 := len(crabs)-1
	for i := high; i >= low; i-- {
		totalCost2 += stepCost2
		leftMove[low+i] = totalCost2
		for pos2 > 0 && crabs[pos2] == i {
			n2++
			pos2--
		}
		stepCost2 += n2
	}

	fmt.Printf("right: %v\n", rightMove)
	fmt.Printf("left: %v\n", leftMove)

	best := math.MaxInt
	for i := range rightMove {
		sum := rightMove[i] + leftMove[i]
		if sum < best {
			best = sum
		}
	}
	fmt.Printf("best costs %d gas\n", best)
}