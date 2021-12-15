package main

import (
	"fmt"
	"math"
	"strconv"
)

func adventDay15A(path string) {
	lines := readLines(path)

	height := len(lines)
	width := len(lines[0])

	grid := make([]int, width*height)
	for i,line := range lines {
		for j,r := range line {
			num, _ := strconv.Atoi(string(r))
			grid[i*width + j] = num
		}
	}

	distances := make([]int, width*height)
	distances[0] = 0
	for i := 1; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}

	nchanged := 1
	for nchanged > 0 {
		nchanged = 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if x == 0 && y == 0 {
					distances[0] = 0
				} else {
					yDist := math.MaxInt
					if y > 0 {
						yDist = distances[(y-1)*width + x]
					}
					xDist := math.MaxInt
					if x > 0 {
						xDist = distances[y*width + x-1]
					}
					min1 := iMin(xDist, yDist)

					y2Dist := math.MaxInt
					if y < height-1 {
						y2Dist = distances[(y+1)*width + x]
					}
					x2Dist := math.MaxInt
					if x < width-1 {
						x2Dist = distances[y*width + x+1]
					}
					min2 := iMin(y2Dist, x2Dist)
					newDist := iMin(min1, min2) + grid[y*width + x]
					oldDist := distances[y*width + x]
					if newDist < oldDist {
						distances[y*width + x] = newDist
						nchanged++
					}
				}
			}
		}

	}
	fmt.Printf("distance to corner is %d\n", distances[len(distances)-1])
	//x := 0
	//y := 0
	//for x < width-1 && y < height-1 {
	//	this := distances[y*width + x]
	//	if x > 0 {
	//		index := y*width + (x-1)
	//		newDist := this + grid[index]
	//		if newDist < distances[index] {
	//			distances[index] = newDist
	//		}
	//	}
	//}

}

func adventDay15B(path string) {
	lines := readLines(path)

	origHeight := len(lines)
	origWidth := len(lines[0])

	gridOrig := make([]int, origWidth*origHeight)
	for i,line := range lines {
		for j,r := range line {
			num, _ := strconv.Atoi(string(r))
			gridOrig[i*origWidth + j] = num
		}
	}

	width := origWidth*5
	height := origHeight*5
	grid := make([]int, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			val := gridOrig[(y%origHeight)*origWidth + x%origWidth]
			val += y/origHeight
			val += x/origWidth
			toSub := val/9
			val -= toSub*9
			if val == 0 {
				val = 9
			}
			grid[y*width + x] = val
		}
	}
	//fmt.Printf("%v\n", grid[3*width:(3*width)+origWidth*2])

	distances := make([]int, width*height)
	distances[0] = 0
	for i := 1; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}

	//brute force, we dont need no djikstras
	nchanged := 1
	for nchanged > 0 {
		nchanged = 0
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if x == 0 && y == 0 {
					distances[0] = 0
				} else {
					yDist := math.MaxInt
					if y > 0 {
						yDist = distances[(y-1)*width + x]
					}
					xDist := math.MaxInt
					if x > 0 {
						xDist = distances[y*width + x-1]
					}
					min1 := iMin(xDist, yDist)

					y2Dist := math.MaxInt
					if y < height-1 {
						y2Dist = distances[(y+1)*width + x]
					}
					x2Dist := math.MaxInt
					if x < width-1 {
						x2Dist = distances[y*width + x+1]
					}
					min2 := iMin(y2Dist, x2Dist)
					newDist := iMin(min1, min2) + grid[y*width + x]
					oldDist := distances[y*width + x]
					if newDist < oldDist {
						distances[y*width + x] = newDist
						nchanged++
					}
				}
			}
		}

	}
	fmt.Printf("distance to corner is %d\n", distances[len(distances)-1])


}