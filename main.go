package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	return lines
}

func toInts(strs []string) []int {
	out := make([]int, len(strs))
	for i,str := range strs {
		val,err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		out[i] = int(val)
	}
	return out
}


func checkSlope(isTree []bool, width, height, slopeX,slopeY int) int {
	xPos := 0
	treesHit := 0
	for yPos := 0; yPos < height; yPos += slopeY {
		if isTree[yPos*width + (xPos % width)] {
			treesHit++
		}
		xPos += slopeX
	}
	return treesHit
}



// readLines reads a whole file into memory
// and returns a slice of its lines.
func readPassports(path string) ([]string) {
	file, err := os.Open(path)
	if err != nil {
		 panic(err)
	}
	defer file.Close()

	var passports []string
	scanner := bufio.NewScanner(file)
	passport := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			passport += line + "\n"
		} else {
			passports = append(passports, passport)
			passport = ""
		}
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return passports
}

func skipped(path string) {
	fmt.Printf("skipped this one\n")
}

var days = []func(path string){
	skipped, skipped,
	skipped, skipped,
	skipped, skipped,
	adventDay4A, adventDay4B,
	adventDay5A, adventDay5B,
	adventDay6A, adventDay6B,
	adventDay7A, adventDay7B,
}

func usage() {
	fmt.Printf("usage:\n\t%s <day number OR filename starting with day number>\n", os.Args[0])
}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		usage()
		return
	}

	var inputs []string

	dayToRun, err := strconv.ParseInt(flag.Args()[0], 10, 64)
	if err == nil {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			filename := file.Name()
			res := strings.Split(filename, "_")
			if len(res) > 1 {
				val, err :=  strconv.ParseInt(res[0], 10, 64)
				if err != nil {
					continue
				}
				if val == dayToRun {
					inputs = append(inputs, filename)
				}
			}
		}
	} else {
		filename := flag.Args()[0]
		dayToRun, err = strconv.ParseInt(strings.Split(filename, "_")[0], 10, 64)
		if err != nil {
			usage()
			return
		}
		inputs = []string{filename}
	}


	aIndex := int(dayToRun-1)*2
	if aIndex >= len(days) {
		fmt.Printf("Haven't written this yet!\n")
		return
	}

	fmt.Printf("Part A\n=====================\n")
	for _,filename := range inputs {
		fmt.Printf("%s:\n", strings.SplitN(filename, "_", 2)[1])
		days[aIndex](filename)
		fmt.Printf("\n")
	}

	bIndex := int(dayToRun-1)*2 + 1
	if bIndex >= len(days) {
		//fmt.Printf("Haven't written this yet!\n")
		return
	}

	fmt.Printf("\nPart B\n=====================\n")
	for _,filename := range inputs {
		fmt.Printf("%s:\n", strings.SplitN(filename, "_", 2)[1])
		days[bIndex](filename)
		fmt.Printf("\n")
	}




}


