package main

import (
	"fmt"
	"sort"
)

func adventDay10A(path string) {
	lines := readLines(path)

	stack := make([]rune, 1000)

	//charScores := map[rune]int {
	//	')': 3,
	//	']': 57,
	//	'}': 1197,
	//	'>': 25137,
	//}

	score := 0
	for _,line := range lines {
		pos := 0
		out:
		for _,r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack[pos] = r
				pos++
			case ')':
				if stack[pos-1] != '(' {
					score += 3
					break out
				}
				pos--
			case ']':
				if stack[pos-1] != '[' {
					score += 57
					break out
				}
				pos--
			case '}':
				if stack[pos-1] != '{' {
					score += 1197
					break out
				}
				pos--
			case '>':
				if stack[pos-1] != '<' {
					score += 25137
					break out
				}
				pos--
			}
		}
	}
	fmt.Printf("score is %d\n", score)
}

func adventDay10B(path string) {
	lines := readLines(path)

	stack := make([]rune, 1000)

	charScores := map[rune]int {
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}


	var lineScores []int
	for _,line := range lines {
		pos := 0
		good := true
	out:
		for _,r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack[pos] = r
				pos++
			case ')':
				if stack[pos-1] != '(' {
					good = false
					break out
				}
				pos--
			case ']':
				if stack[pos-1] != '[' {
					good = false
					break out
				}
				pos--
			case '}':
				if stack[pos-1] != '{' {
					good = false
					break out
				}
				pos--
			case '>':
				if stack[pos-1] != '<' {
					good = false
					break out
				}
				pos--
			}
		}
		if good {
			lineScore := 0
			for i := pos-1; i>=0; i-- {
				lineScore *= 5
				lineScore += charScores[stack[i]]
			}
			lineScores = append(lineScores, lineScore)
		}
	}
	sort.Ints(lineScores)
	fmt.Printf("score is %d\n", lineScores[len(lineScores)/2])

}

