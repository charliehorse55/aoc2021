package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adventDay2A(path string) {
	directions := readLines(path)

	dist := 0

	depth := 0

	for _,line :=  range directions {
		parts := strings.Fields(line)
		dir := parts[0]
		val, _ := strconv.ParseInt(parts[1], 10, 64)
		amt := int(val)
		switch dir {
		case "forward":
			dist += amt
		case "down":
			depth += amt
		case "up":
			depth -= amt
		}
	}

	fmt.Printf("multiplied = %d\n", dist*depth)
}

func adventDay2B(path string) {
	directions := readLines(path)

	dist := 0
	aim := 0
	depth := 0

	for _,line :=  range directions {
		parts := strings.Fields(line)
		dir := parts[0]
		val, _ := strconv.ParseInt(parts[1], 10, 64)
		amt := int(val)
		switch dir {
		case "forward":
			dist += amt
			depth += aim*amt
		case "down":
			aim += amt
		case "up":
			aim -= amt
		}
	}

	fmt.Printf("multiplied = %d\n", dist*depth)

}