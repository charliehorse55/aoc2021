package main

import (
	"fmt"
	"strconv"
)

func iMax(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func iMin(a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func flash(grid []int, width, height, i, j int) {
	//mark as flashed
	grid[i*width + j] = -1
	//fmt.Printf("flashing at %2d,%2d\n", i, j)
	//fmt.Printf("looping from %d to %di, %d to %d j\n", iMax(0,i-1), iMin(height-1, i+1),  iMax(0,j-1), iMin(width-1, j+1))
	for k := iMax(0,i-1); k <= iMin(height-1, i+1); k++ {
		for w := iMax(0,j-1); w <= iMin(width-1, j+1); w++ {
			if grid[k*width+w] >= 0 {
				grid[k*width + w]++
				if grid[k*width + w] > 9 {
					flash(grid, width, height, k, w)
				}
			}
		}
	}
}

func adventDay11A(path string) {
	lines := readLines(path)

	width := len(lines[0])
	height := len(lines)
	state := make([]int, width*height)
	for i,line := range lines {
		for j,r := range line {
			state[i*width + j],_ = strconv.Atoi(string(r))
		}
	}

	//for i := 0; i < height; i++ {
	//	fmt.Printf("%v\n", state[i*width:(i+1)*width])
	//}
	//fmt.Printf("\n")


	flashes := 0
	const steps = 100
	for step := 0; step < steps; step++ {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				state[i*width+j]++
			}
		}
		//for i := 0; i < height; i++ {
		//	fmt.Printf("%2v\n", state[i*width:(i+1)*width])
		//}
		//fmt.Printf("\n")

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if state[i*width + j] > 9 {
					//flash it
					flash(state, width, height, i, j)
				}
			}
		}
		//for i := 0; i < height; i++ {
		//	fmt.Printf("%2v\n", state[i*width:(i+1)*width])
		//}
		//fmt.Printf("\n")


		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if state[i*width + j] == -1 {
					flashes++
					state[i*width+j] = 0
				}
			}
		}
		//for i := 0; i < height; i++ {
		//	fmt.Printf("%v\n", state[i*width:(i+1)*width])
		//}
		//fmt.Printf("\n")
	}
	fmt.Printf("%d flashes after %d steps\n", flashes, steps)
}

func adventDay11B(path string) {
	lines := readLines(path)

	width := len(lines[0])
	height := len(lines)
	state := make([]int, width*height)
	for i,line := range lines {
		for j,r := range line {
			state[i*width + j],_ = strconv.Atoi(string(r))
		}
	}

	//for i := 0; i < height; i++ {
	//	fmt.Printf("%v\n", state[i*width:(i+1)*width])
	//}
	//fmt.Printf("\n")


	const steps = 10000
	for step := 0; step < steps; step++ {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				state[i*width+j]++
			}
		}

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if state[i*width + j] > 9 {
					//flash it
					flash(state, width, height, i, j)
				}
			}
		}

		flashes := 0
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if state[i*width + j] == -1 {
					flashes++
					state[i*width+j] = 0
				}
			}
		}
		if flashes == width*height {
			fmt.Printf("after %d steps, converged\n", step+1)
			break
		}
	}

}