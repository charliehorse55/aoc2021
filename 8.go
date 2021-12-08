package main

import (
	"fmt"
	"sort"
	"strings"
)




func parseday8(line string)(input, output []string) {
	parts := strings.Split(line, "|")
	input = strings.Fields(parts[0])[:10]
	output = strings.Fields(parts[1])[:4]
	return
}

func adventDay8A(path string) {
	lines := readLines(path)

	total := 0
	for _,line := range lines {
		_, output := parseday8(line)
		for _,str := range output {
			l := len(str)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				total++
			}
		}
	}

	fmt.Printf("a total of %d different 1s,4s,7s, and 8s were found\n", total)
}

func rfl(x []string, n int) string {
	for _,str := range x {
		if len(str) == n {
			return str
		}
	}
	return ""
}

func sub(x, y string) string {

	for _,r := range y {
		x = strings.Replace(x, string(r), "", -1)
	}
	return x
}

func nWith(x []string, incl rune) int {
	total := 0
	for _,str := range x {
		if strings.ContainsRune(str, incl) {
			total++
		}
	}
	return total
}

func filterStrs(x []string, incl rune, excl rune) (ret [2]string) {
	total := 0
	for _,str := range x {
		if strings.ContainsRune(str, incl) && !strings.ContainsRune(str, excl) {
			ret[total] = str
			total++
		}
	}
	return
}



type sevenSeg struct {
	top,mid,bot rune
	upLeft, upRight rune
	downLeft, downRight rune
}

func (s sevenSeg)Print() {
	fmt.Printf(" %c%c%c%c \n", s.top, s.top, s.top, s.top)
	fmt.Printf("%c    %c\n", s.upLeft, s.upRight)
	fmt.Printf("%c    %c\n", s.upLeft, s.upRight)
	fmt.Printf(" %c%c%c%c \n", s.mid, s.mid, s.mid, s.mid)
	fmt.Printf("%c    %c\n", s.downLeft, s.downRight)
	fmt.Printf("%c    %c\n", s.downLeft, s.downRight)
	fmt.Printf(" %c%c%c%c \n", s.bot, s.bot, s.bot, s.bot)
}

func (s sevenSeg)String() string {
	out := ""
	if s.top != 0 {
		out += string(s.top)
	}
	if s.mid != 0 {
		out += string(s.mid)
	}
	if s.bot != 0 {
		out += string(s.bot)
	}
	if s.upLeft != 0 {
		out += string(s.upLeft)
	}
	if s.upRight != 0 {
		out += string(s.upRight)
	}
	if s.downLeft != 0 {
		out += string(s.downLeft)
	}
	if s.downRight != 0 {
		out += string(s.downRight)
	}
	return out
}

func (s sevenSeg)Encode(x string) string {
	output := ""
	for _,r := range x {
		switch r {
		case 'a':
			output += string(s.top)
		case 'b':
			output += string(s.upLeft)
		case 'c':
			output += string(s.upRight)
		case 'd':
			output += string(s.mid)
		case 'e':
			output += string(s.downLeft)
		case 'f':
			output += string(s.downRight)
		case 'g':
			output += string(s.bot)
		}
	}
	return SortStringByCharacter(output)
}


func adventDay8B(path string) {

	lines := readLines(path)

	sum := 0

	for _,line := range lines {
		input, output := parseday8(line)

		var pattern sevenSeg

		letterFreq := make(map[rune]int)
		for _,str := range input {
			for _,r := range str {
				letterFreq[r]++
			}
		}
		for k,v := range letterFreq {
			if v == 6 {
				pattern.upLeft = k
			}
		}

		//find
		rightSide := rfl(input, 2)
		seven := rfl(input, 3)
		pattern.top = rune(sub(seven, rightSide)[0])

		if nWith(input, rune(rightSide[1])) == 9 {
			pattern.downRight = rune(rightSide[1])
			pattern.upRight = rune(rightSide[0])
		} else {
			pattern.downRight = rune(rightSide[0])
			pattern.upRight = rune(rightSide[1])
		}
		five6 := filterStrs(input, pattern.downRight, pattern.upRight)
		if len(five6[0]) > len(five6[1]) {
			five6[0], five6[1] = five6[1], five6[0]
		}
		pattern.downLeft = rune(sub(five6[1], five6[0])[0])
		//fmt.Printf("5 and 6 are %v\n", five6)

		for _,str := range input {
			if len(str) == 6 {
				if strings.ContainsRune(str, pattern.downLeft) && strings.ContainsRune(str, pattern.upRight){
					pattern.bot = rune(sub(str, pattern.String())[0])
					pattern.mid = rune(sub("abcdefg", pattern.String())[0])
				}
			}
		}

		//pattern.Print()
		//fmt.Printf("\n")

		//create a map
		lookup := make(map[string]int)
		lookup[pattern.Encode("abcefg")] = 0
		lookup[pattern.Encode("cf")] = 1
		lookup[pattern.Encode("acdeg")] = 2
		lookup[pattern.Encode("acdfg")] = 3
		lookup[pattern.Encode("bcdf")] = 4
		lookup[pattern.Encode("abdfg")] = 5
		lookup[pattern.Encode("abdefg")] = 6
		lookup[pattern.Encode("acf")] = 7
		lookup[pattern.Encode("abcdefg")] = 8
		lookup[pattern.Encode("abcdfg")] = 9

		num := 0
		for _,str := range output {
			num *= 10
			num += lookup[SortStringByCharacter(str)]
		}
		sum += num
	}
	fmt.Printf("sum is %d\n", sum)
}


type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}
