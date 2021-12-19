package main

import (
	"fmt"
	"strings"
)

type vector [3]int

func VectorDifference(a, b vector) vector {
	return vector{a[0]-b[0], a[1]-b[1], a[2]-b[2]}
}

var orientations = map[int][3]int{
	0: {0,1,2},
	1: {0,2,1},
	2: {1,0,2},
	3: {1,2,0},
	4: {2,1,0},
	5: {2,0,1},
}

func rotateScanner(scanner []vector, rot int) []vector {
	columns := orientations[rot/8]
	polarity := rot % 8

	xPol := 1
	if polarity & 0x1 != 0 {
		xPol = -1
	}
	yPol := 1
	if polarity & 0x2 != 0 {
		yPol = -1
	}
	zPol := 1
	if polarity & 0x4 != 0 {
		zPol = -1
	}

	dest := make([]vector, len(scanner))
	for i,vec := range scanner {
		dest[i] = vector{
			xPol*vec[columns[0]],
			yPol*vec[columns[1]],
			zPol*vec[columns[2]],
		}
	}
	return dest
}

func findPotentialCommon(a, b []vector) map[[3]int]int {
	result := make(map[[3]int]int)
	for _,vecA := range a {
		for _,vecB := range b {
			diff := VectorDifference(vecA, vecB)
			result[diff]++
		}
	}
	return result
}


func parseScanners(path string) [][]vector {
	lines := readLines(path)
	lineN := 0

	var scanners [][]vector
	for lineN < len(lines) {
		//skip the scanner name
		lineN++
		var scanner []vector
		for lineN < len(lines) && !strings.HasPrefix(lines[lineN], "--- scanner") {
			parts := strings.Split(lines[lineN], ",")
			ints := toInts(parts)
			scanner = append(scanner, vector{
				ints[0], ints[1], ints[2],
			})
			lineN++
		}
		scanners = append(scanners, scanner)
	}
	return scanners
}


func adventDay19A(path string) {
	scanners := parseScanners(path)

	fmt.Printf("%v\n", scanners[0][0])

	masterNodes := make(map[vector]bool)
	for _,vec := range scanners[0] {
		masterNodes[vec] = true
	}

	type base struct {
		offset vector
		rotation int
		b *base
	}
	validBases := make(map[int]*base)
	validBases[0] = &base{
		offset:   vector{0,0,0},
		rotation: 0,
		b: nil,
	}

	type toTryLater struct {
		rotated []vector
		offset vector
		i,j int
		rot int
	}

	var tryLater []toTryLater

	for i,topScan := range scanners {
		for rot := 0; rot < 48; rot++ {
			for j, scanner := range scanners {
				_,ok := validBases[j]
				if i == j || ok {
					continue
				}
				rotated := rotateScanner(scanner, rot)
				countsForOffset := findPotentialCommon(topScan, rotated)
				validOffsets := 0
				for offset,count := range countsForOffset {
					if count >= 12 {
						tryLater = append(tryLater, toTryLater{
							rotated: rotated,
							offset:  offset,
							i:i,
							j:j,
							rot:rot,
						})
						validOffsets++
					}
				}
				if validOffsets > 1 {
					panic("uh oh")
				}
			}
		}

	}

	for len(tryLater) > 0 {
		var onesLeft []toTryLater
		for _,a := range tryLater {
			b, ok := validBases[a.i]
			if !ok {
				onesLeft = append(onesLeft, a)
				continue
			}

			for k := range a.rotated {
				a.rotated[k][0] += a.offset[0]
				a.rotated[k][1] += a.offset[1]
				a.rotated[k][2] += a.offset[2]
			}

			validBases[a.j] = &base{
				offset:   a.offset,
				rotation: a.rot,
				b: validBases[a.i],
			}

			for b != nil {
				//fmt.Printf("doing rotation: %d, %d\n", i,j)
				//fmt.Printf("rotated[0] = %v\n", rotated[0])
				a.rotated = rotateScanner(a.rotated, b.rotation)
				//fmt.Printf("rotated[0] = %v\n", rotated[0])
				for k := range a.rotated {
					a.rotated[k][0] += b.offset[0]
					a.rotated[k][1] += b.offset[1]
					a.rotated[k][2] += b.offset[2]
				}
				b = b.b
			}
			for _,vec := range a.rotated {
				masterNodes[vec] = true
			}
			fmt.Printf("found valid offset of scanner %d with scanner %d is %v with rotation %d (%d unique beacons)\n", a.i, a.j, a.offset, a.rot, len(masterNodes))

		}
		tryLater = onesLeft
	}
	fmt.Printf("%d valid bases, %d scanners\n", len(validBases), len(scanners))
	scannerPos := make([]vector, len(scanners))
	scannerPos[0] = vector{0,0,0}
	for i := 1; i < len(scanners); i++ {
		b := validBases[i]
		rotated := []vector{{0,0,0}}
		for b != nil {
			//fmt.Printf("doing rotation: %d, %d\n", i,j)
			//fmt.Printf("rotated[0] = %v\n", rotated[0])
			rotated = rotateScanner(rotated, b.rotation)
			//fmt.Printf("rotated[0] = %v\n", rotated[0])
			for k := range rotated {
				rotated[k][0] += b.offset[0]
				rotated[k][1] += b.offset[1]
				rotated[k][2] += b.offset[2]
			}
			b = b.b
		}
		scannerPos[i] = rotated[0]
		fmt.Printf("scanner %d at %v\n", i, rotated[0])
	}

	maxDistance := 0
	for i,posA := range scannerPos {
		for j := i+1; j < len(scannerPos); j++ {
			delta := 0
			delta += intABS(posA[0]-scannerPos[j][0])
			delta += intABS(posA[1]-scannerPos[j][1])
			delta += intABS(posA[2]-scannerPos[j][2])
			if delta > maxDistance {
				maxDistance = delta
			}
		}
	}
	fmt.Printf("max distance is %d\n", maxDistance)
}


func adventDay19B(path string) {

}
