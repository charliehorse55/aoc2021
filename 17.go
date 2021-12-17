package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRange(str string) (int, int) {
	parts := strings.Split(str, "..")
	min,_ := strconv.Atoi(parts[0])
	max,_ := strconv.Atoi(parts[1])
	return min, max
}

func parseRanges(str string) (xMin, xMax, yMin, yMax int) {
	parts := strings.Split(str, ",")
	xMin,xMax = parseRange(strings.TrimPrefix(parts[0], "target area: x="))
	yMin,yMax = parseRange(strings.TrimPrefix(parts[1], " y="))
	return
}

func adventDay17A(path string) {
	lines := readLines(path)
	xMin, xMax, yMin, yMax := parseRanges(lines[0])
	fmt.Printf("%d->%d, %d->%d\n", xMin, xMax, yMin, yMax)

	vYMax := (-yMin)-1
	fmt.Printf("max y is %d\n", (vYMax)*(vYMax+1)/2)
}

func adventDay17B(path string) {
	lines := readLines(path)
	xMin, xMax, yMin, yMax := parseRanges(lines[0])
	fmt.Printf("%d->%d, %d->%d\n", xMin, xMax, yMin, yMax)

	vYMax := (-yMin)-1

	worked := 0
	for vXstart := 1; vXstart <= xMax; vXstart++ {
		for vYstart := yMin; vYstart <= vYMax; vYstart++ {
			vX := vXstart
			vY := vYstart

			yPos := 0
			xPos := 0
			for yPos > yMin {
				xPos += vX
				yPos += vY
				if xPos >= xMin && xPos <= xMax && yPos >= yMin && yPos <= yMax {
					worked++
					//fmt.Printf("(%d,%d) worked!\n", vXstart, vYstart)
					break
				}
				if vX > 0 {
					vX--
				}
				vY--
			}

		}
	}

	fmt.Printf("worked %d times\n", worked)

}