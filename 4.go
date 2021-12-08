package main

import (
	"fmt"
	"strings"
)

func checkWin(board board) bool {
	//check rows
	for i := 0; i < 5; i++ {
		good := true
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				good = false
				break
			}
		}
		if good {
			return true
		}
	}
	for j := 0; j < 5; j++ {
		good := true
		for i := 0; i < 5; i++ {
			if board[i][j] != -1 {
				good = false
				break
			}
		}
		if good {
			return true
		}
	}
	return false
}

func totalPositive(board board) int {
	total := 0
	for _,row := range board {
		for _, val := range row {
			if val > 0 {
				total += val
			}
		}
	}
	return total
}

type board [5][5]int

func adventDay4A(path string) {
	lines := readLines(path)
	drawn := toInts(strings.Split(lines[0], ","))

	lines = lines[1:]

	boards := make([]board, len(lines)/5)
	for i := range boards {
		for k := 0; k < 5; k++ {
			a := toInts(strings.Fields(lines[0]))
			for j := 0; j < 5; j++ {
				boards[i][k][j] = a[j]
			}
			lines = lines[1:]
		}
		//fmt.Printf("scanned board:\n%+v\n", boards[i])
	}

	for _,draw := range drawn {
		for i := range boards {
			for j,row := range boards[i] {
				for k,val := range row {
					if val == draw {
						boards[i][j][k] = -1
						if checkWin(boards[i]) {
							fmt.Printf("board %d wins with a score of %d\n", i, totalPositive(boards[i])*draw)
							return
						}
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", drawn[0])

}

func adventDay4B(path string) {
	lines := readLines(path)
	drawn := toInts(strings.Split(lines[0], ","))

	lines = lines[1:]

	boards := make([]board, len(lines)/5)
	for i := range boards {
		for k := 0; k < 5; k++ {
			a := toInts(strings.Fields(lines[0]))
			for j := 0; j < 5; j++ {
				boards[i][k][j] = a[j]
			}
			lines = lines[1:]
		}
		//fmt.Printf("scanned board:\n%+v\n", boards[i])
	}

	boardsWon := make([]bool, len(boards))
	boardsLeft := len(boards)
	for _,draw := range drawn {
		for i := range boards {
			if boardsWon[i] {
				continue
			}
			for j,row := range boards[i] {
				for k,val := range row {
					if val == draw {
						boards[i][j][k] = -1
						if checkWin(boards[i]) {
							boardsWon[i] = true
							boardsLeft--
							if boardsLeft == 0 {
								fmt.Printf("board %d wins last with a score of %d\n", i, totalPositive(boards[i])*draw)
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("%d\n", drawn[0])

}