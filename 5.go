package main

import (
	"fmt"
	"strconv"
	"strings"
)

const gridSize = 1000

func parsePoint(in string) (x, y int) {
	strs := strings.Split(in, ",")
	xI, _ := strconv.ParseInt(strs[0], 10, 64)
	yI, _ := strconv.ParseInt(strs[1], 10, 64)
	return int(xI), int(yI)
}

func adventDay5A(path string) {
	lines := readLines(path)

	grid := make([]int, gridSize*gridSize)

	total := 0

	for _,line := range lines {
		parts := strings.Split(line, " -> ")
		x1,y1 := parsePoint(parts[0])
		x2,y2 := parsePoint(parts[1])

		if x1 != x2 && y1 != y2 {
			//fmt.Printf("skipping (%d,%d)->(%d,%d)\n", x1, y1, x2, y2)
			continue
		}

		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				index := i + j*gridSize
				grid[index]++
				if grid[index] == 2 {
					total++
				}
			}
		}
	}
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v\n", grid[gridSize*i:gridSize*i+10])
	//}
	fmt.Printf("total is %d points\n", total)
}
func adventDay5B(path string) {
	lines := readLines(path)

	grid := make([]int, gridSize*gridSize)

	total := 0

	for _,line := range lines {
		parts := strings.Split(line, " -> ")
		x1,y1 := parsePoint(parts[0])
		x2,y2 := parsePoint(parts[1])

		xInc := 0
		if x1 > x2 {
			xInc = -1
		} else if x1 < x2 {
			xInc = 1
		}

		yInc := 0
		if y1 > y2 {
			yInc = -1
		} else if y1 < y2 {
			yInc = 1
		}
		length := intABS(x2-x1)
		if length == 0 {
			length = intABS(y2-y1)
		}
		length++

		for i := 0; i < length; i++ {
			x := x1 + xInc*i
			y := y1 + yInc*i
			index := x + y*gridSize
			grid[index]++
			if grid[index] == 2 {
				total++
			}
		}
	}
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v\n", grid[gridSize*i:gridSize*i+10])
	//}
	fmt.Printf("total is %d points\n", total)
}
