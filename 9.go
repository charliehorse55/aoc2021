package main

import (
	"fmt"
	"sort"
	"strconv"
)

func adventDay9A(path string) {
	lines := readLines(path)
	width := len(lines[0])
	height := len(lines)
	grid := make([]int, width*height)
	for i,line := range lines {
		for j,r := range line {
			v, _ := strconv.Atoi(string(r))
			grid[i*width + j] = v
		}
	}
	score := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			at := grid[i*width + j]

			if i > 0 && at >= grid[(i-1)*width + j] {
				continue
			}
			if j < width-1 && at >=  grid[(i)*width + j+1] {
				continue
			}
			if i < height-1 && at >= grid[(i+1)*width + j] {
				continue
			}
			if j > 0 && at >= grid[(i)*width + j-1]{
				continue
			}
			score += at + 1
		}
	}
	//for i := 0; i < height; i++ {
	//	fmt.Printf("%v\n", grid[i*width:(i+1)*width])
	//}

	fmt.Printf("score is %d\n", score)
}

func turninto9s(grid []int, width, height,x,y int) int {
	if grid[x*width + y] == 9 {
		return 0
	}
	grid[x*width + y] = 9
	total := 1

	if x > 0 {
		total += turninto9s(grid, width, height, x-1, y)
	}
	if x < height-1 {
		total += turninto9s(grid, width, height, x+1, y)
	}
	if y > 0 {
		total += turninto9s(grid, width, height, x, y-1)
	}
	if y < width-1 {
		total += turninto9s(grid, width, height, x, y+1)
	}
	return total
}

func adventDay9B(path string) {
	lines := readLines(path)
	width := len(lines[0])
	height := len(lines)
	grid := make([]int, width*height)
	for i, line := range lines {
		for j, r := range line {
			v, _ := strconv.Atoi(string(r))
			grid[i*width+j] = v
		}
	}

	var basins []int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i*width + j] == 9 {
				continue
			}
			size := turninto9s(grid, width, height, i,j)
			if size > 0 {
				basins = append(basins, size)
			}
		}
	}
	sort.Ints(basins)
	//fmt.Printf("%v\n", basins)
	b1 := basins[len(basins)-1]
	b2 := basins[len(basins)-2]
	b3 := basins[len(basins)-3]
	fmt.Printf("3 biggest are: %d, %d, %d, product is %d\n", b1, b2, b3, b1*b2*b3)
}
