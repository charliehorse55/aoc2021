package main

import (
	"fmt"
	"strings"
)

func parseCubeLine(line string) (on bool, xMin,xMax,yMin,yMax,zMin,zMax int) {
	parts := strings.Split(line, " ")
	if parts[0] == "on" {
		on = true
	}

	ranges := strings.Split(parts[1], ",")
	xMin,xMax = parseRange(ranges[0][2:])
	yMin,yMax = parseRange(ranges[1][2:])
	zMin,zMax = parseRange(ranges[2][2:])
	return
}



func adventDay22A(path string) {
	lines := readLines(path)

	grid := make([]bool, 101*101*101)
	for _,line := range lines {
		on, xMin,xMax,yMin,yMax,zMin,zMax := parseCubeLine(line)
		xMin = iMax(xMin, -50)
		xMax = iMin(xMax, 50)

		yMin = iMax(yMin, -50)
		yMax = iMin(yMax, 50)

		zMin = iMax(zMin, -50)
		zMax = iMin(zMax, 50)

		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				for k := zMin; k <= zMax; k++ {
					grid[(i+50)*101*101 + (j+50)*101 + (k+50)] = on
				}
			}
		}
	}

	total := 0
	for _,val := range grid {
		if val {
			total++
		}
	}
	fmt.Printf("total %d on\n", total)
}

type cuboid struct {
	xMin, xMax int
	yMin, yMax int
	zMin, zMax int
}

func (c cuboid)volume() int {
	dx := c.xMax - c.xMin + 1
	dy := c.yMax - c.yMin + 1
	dz := c.zMax - c.zMin + 1
	return dx * dy * dz
}

func overlapRange(x1, x2, y1, y2 int) bool {
	return x1 <= y2 && y1 <= x2
}

func RectsOverlap(a, b cuboid) bool {
	x := overlapRange(a.xMin, a.xMax, b.xMin, b.xMax)
	y := overlapRange(a.yMin, a.yMax, b.yMin, b.yMax)
	z := overlapRange(a.zMin, a.zMax, b.zMin, b.zMax)
	return x && y && z
}


func adventDay22B(path string) {
	lines := readLines(path)


	var currCuboids []cuboid
	for _, line := range lines {
		on, xMin, xMax, yMin, yMax, zMin, zMax := parseCubeLine(line)
		toAdd := cuboid{
			xMin: xMin,
			xMax: xMax,
			yMin: yMin,
			yMax: yMax,
			zMin: zMin,
			zMax: zMax,
		}
		var nextCuboids []cuboid
		for _,existing := range currCuboids {
			if RectsOverlap(existing, toAdd) {
				//can add top?
				top := existing
				top.yMin = toAdd.yMax+1
				if top.volume() > 0 {
					nextCuboids = append(nextCuboids, top)
				}
				bottom := existing
				bottom.yMax = toAdd.yMin-1
				if bottom.volume() > 0 {
					nextCuboids = append(nextCuboids, bottom)
				}
				//
				leftSide := existing
				leftSide.yMax = iMin(toAdd.yMax, existing.yMax)
				leftSide.yMin = iMax(toAdd.yMin, existing.yMin)
				leftSide.xMax = toAdd.xMin-1
				if leftSide.volume() > 0 {
					nextCuboids = append(nextCuboids, leftSide)
				}


				rightSide := existing
				rightSide.yMax = iMin(toAdd.yMax, existing.yMax)
				rightSide.yMin = iMax(toAdd.yMin, existing.yMin)
				rightSide.xMin = toAdd.xMax+1
				if rightSide.volume() > 0 {
					nextCuboids = append(nextCuboids, rightSide)
				}

				front := existing
				front.yMax = iMin(toAdd.yMax, existing.yMax)
				front.yMin = iMax(toAdd.yMin, existing.yMin)
				front.xMax = iMin(toAdd.xMax, existing.xMax)
				front.xMin = iMax(toAdd.xMin, existing.xMin)
				front.zMin = toAdd.zMax+1
				if front.volume() > 0 {
					nextCuboids = append(nextCuboids, front)
				}

				back := front
				back.zMin = existing.zMin
				back.zMax = toAdd.zMin-1
				if back.volume() > 0 {
					nextCuboids = append(nextCuboids, back)
				}
			} else {
				nextCuboids = append(nextCuboids, existing)
			}

		}
		if on == true {
			nextCuboids = append(nextCuboids, toAdd)
		}
		currCuboids = nextCuboids
	}
	total := 0
	for _,rect := range currCuboids {
		total += rect.volume()
	}
	fmt.Printf("total volume is %d\n", total)
}
