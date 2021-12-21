package main

import (
	"fmt"
)

const player1Start = 1
const player2Start = 5

type die struct {
	pos int
	n int
}

func (d *die)Roll() int {
	tmp := d.pos
	d.pos++
	d.pos %= 101
	if d.pos == 0 {
		d.pos = 1
	}
	d.n++
	return tmp
}

func (d *die)Roll3() int {
	return d.Roll() + d.Roll() + d.Roll()
}

func adventDay21A(path string) {
	dice := die{pos:1}
	player1Score := 0
	player2Score := 0
	p1Pos := player1Start-1
	p2Pos := player2Start-1

	for {
		p1Pos = (p1Pos+dice.Roll3()) % 10
		player1Score += p1Pos + 1
		if player1Score >= 1000 {
			score := player2Score * dice.n
			fmt.Printf("player 1 wins with a score of %d, %d points\n", player1Score, score)
			break
		}
		p2Pos = (p2Pos+dice.Roll3()) % 10
		player2Score += p2Pos + 1
		if player2Score >= 1000 {
			score := player1Score * dice.n
			fmt.Printf("player 2 wins with a score of %d, %d points\n", player2Score, score)
			break
		}
	}
}

type gameState struct {
	pos [2]uint8
	scores [2]int
	move   uint8
}

var memory = make(map[gameState][2]int)

func DiceWins(state gameState) [2]int {
	val, ok := memory[state]
	if ok {
		return val
	}

	var sums [2]int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				dieSum := uint8(i + j + k)
				ns := state
				ns.pos[ns.move] = (ns.pos[ns.move] + dieSum) % 10
				ns.scores[ns.move] += int(ns.pos[ns.move]) + 1
				if ns.scores[ns.move] >= 21 {
					sums[ns.move]++
				} else {
					ns.move = (ns.move + 1) % 2
					result := DiceWins(ns)
					sums[0] += result[0]
					sums[1] += result[1]
				}
			}

		}
	}
	memory[state] = sums
	return sums
}


func adventDay21B(path string) {

	state := gameState{
		pos:    [2]uint8{player1Start-1, player2Start-1},
		move:   0,
	}
	result := DiceWins(state)
	//if result[0] > result [1]{
	//	fmt.Printf("result: %d\n", result[0])
	//} else {
	//	fmt.Printf("result: %d\n", result[1])
	//}
	fmt.Printf("result: %v\n", result)
}

