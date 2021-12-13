package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func fold(path string) map[[2]int]bool {
	lines := readLines(path)
	points := make(map[[2]int]bool)
	for _,line := range lines {
		if strings.ContainsRune(line, ',') {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points[[2]int{x,y}] = true
		} else {
			parts := strings.Split(line, "=")
			threshold,_ := strconv.Atoi(parts[1])
			var index int
			if strings.HasSuffix(parts[0], "x") {
				index = 0
			} else {
				index = 1
			}
			afterFold := make(map[[2]int]bool)
			for key := range points {
				if key[index] > threshold {
					key[index] = threshold - (key[index]-threshold)
				}
				afterFold[key] = true
			}
			points = afterFold
			fmt.Printf("%d dots remaining\n", len(points))
		}
	}
	return points
}

func adventDay13A(path string) {
	fold(path)
}

func adventDay13B(path string) {
	points := fold(path)
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for point := range points {
		if point[0] < minX {
			minX = point[0]
		}
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] < minY {
			minY = point[1]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
	}
	fmt.Printf("%d->%d, %d-%d\n", minX, maxX, minY, maxY)
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			char := '.'
			_,ok := points[[2]int{j,i}]
			if ok {
				char = '#'
			}
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}

}

