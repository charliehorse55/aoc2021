package main

import "fmt"

func adventDay20A(path string) {
	lines := readLines(path)

	lookup := make([]uint8, 512)
	for i,r := range lines[0] {
		val := uint8(0)
		if r == '#' {
			val = 1
		}
		lookup[i] = val
	}

	const offset = 1000

	width := len(lines[1])+offset*2
	height := len(lines)-1+offset*2

	state := make([]uint8, width*height)

	for i,line := range lines[1:] {
		for j,r := range line {
			val := uint8(0)
			if r == '#' {
				val = 1
			}
			state[(i+offset)*width + j + offset] = val
		}
	}

	getIndex := func (i,j int) uint8 {
		if i < 0 || i >= height || j < 0 || j >= width {
			return 0
		} else {
			return state[i*width + j]
		}
	}

	printGrid := func() {
		//for i := 0; i < height; i++ {
		//	for j := 0; j < width; j++ {
		//		if state[i*width + j] == 0 {
		//			fmt.Printf(".")
		//		} else {
		//			fmt.Printf("#")
		//		}
		//	}
		//	fmt.Printf("\n")
		//}
		//fmt.Printf("\n")
	}
	printGrid()

	const rounds = 50
	nextState := make([]uint8, width*height)
	for round := 0; round < rounds; round++ {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				index := 0
				for k := i-1; k <= i+1; k++ {
					for w := j-1; w <= j+1; w++ {
						index <<= 1
						index |= int(getIndex(k, w))
					}
				}
				nextState[i*width + j] = lookup[index]
			}
		}
		state, nextState = nextState, state
		printGrid()
	}
	total := 0
	for _,v := range state {
		total += int(v)
	}

	//const smallOffset = 100
	//smallWidth := len(lines[1])+smallOffset*2
	//smallHeight := len(lines)-1+smallOffset*2

	totalInner := 0
	for i := 500; i < height-500; i++ {
		for j := 500; j < width-500; j++ {
			totalInner += int(state[i*width + j])
		}
	}

	fmt.Printf("after %d rounds, %d lit pixels (%d inner)\n", rounds, total, totalInner)
}

func adventDay20B(path string) {

}