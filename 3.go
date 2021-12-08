package main

import (
	"fmt"
	"strconv"
)

func adventDay3A(path string) {
	lines := readLines(path)
	counts := make([]int, len(lines[0]))
	for _,line := range lines {
		for i,r := range line[:len(line)] {
			if r == '1' {
				counts[i]++
			}
		}
	}
	gamma := 0
	epsilon := 0
	for _,count := range counts {
		gamma *= 2
		epsilon *= 2
		if count > len(lines)/2 {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	fmt.Printf("gamma is %b\n", gamma)
	fmt.Printf("epsilon is %b\n", epsilon)
	fmt.Printf("power usage is %d\n", gamma*epsilon)
}

func adventDay3B(path string) {
	lines := readLines(path)
	linesBak := make([]string, len(lines))
	copy(linesBak, lines)

	oneStrs := make([]string, 0, len(lines))
	zeroStrs := make([]string, 0, len(lines))

	for i := 0; i < len(lines[0]); i++ {
		oneStrs = oneStrs[:0]
		zeroStrs = zeroStrs[:0]
		if len(lines) == 1 {
			break
		}
		ones := 0
		for _, line := range lines {
			if line[i] == '1' {
				ones++
				oneStrs = append(oneStrs, line)
			} else {
				zeroStrs = append(zeroStrs, line)
			}
		}

		if ones >= (len(lines)+1)/2 {
			copy(lines, oneStrs)
			lines = lines[:len(oneStrs)]
		} else {
			copy(lines, zeroStrs)
			lines = lines[:len(zeroStrs)]
		}
	}

	oxygenNum,_ := strconv.ParseInt(lines[0], 2, 64)

	lines = linesBak
	for i := 0; i < len(lines[0]); i++ {
		oneStrs = oneStrs[:0]
		zeroStrs = zeroStrs[:0]
		if len(lines) == 1 {
			break
		}
		ones := 0
		for _, line := range lines {
			if line[i] == '1' {
				ones++
				oneStrs = append(oneStrs, line)
			} else {
				zeroStrs = append(zeroStrs, line)
			}
		}
		if ones >= (len(lines)+1)/2 {
			copy(lines, zeroStrs)
			lines = lines[:len(zeroStrs)]
		} else {
			copy(lines, oneStrs)
			lines = lines[:len(oneStrs)]
		}
	}

	co2Num,_ := strconv.ParseInt(lines[0], 2, 64)

	fmt.Printf("life support rating = %d\n", oxygenNum*co2Num)



	//fmt.Printf("counts: %v\n", counts)
	//gamma := 0
	//epsilon := 0
	//for _,count := range counts {
	//	gamma *= 2
	//	epsilon *= 2
	//	if count >= len(lines)/2 {
	//		gamma += 1
	//	} else {
	//		epsilon += 1
	//	}
	//}
	//gammaStr := fmt.Sprintf("%b", gamma)
	//var oxygenGen string
	//out:
	//for {
	//	for _,line := range lines {
	//		if strings.HasPrefix(line, gammaStr) {
	//			oxygenGen = line
	//			break out
	//		}
	//	}
	//	gammaStr = gammaStr[:len(gammaStr)-1]
	//}
	//fmt.Printf("oxygen = %s\n", oxygenGen)
}