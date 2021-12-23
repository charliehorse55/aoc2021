package main

import (
	"fmt"
	"math"
)

type state23 struct {
	hallway [11]int
	rooms [4][2]int
	cost int
}


func (s state23)panicInvalid() {
	for _,room := range s.rooms {
		if room[0] != 0 && room[1] == 0 {
			panic("something bad happened!")
		}
	}
}

func (s state23)canLeave(room, depth int) bool {
	if s.rooms[room][depth] != room+1 {
		return true
	}
	if depth == 0 && s.rooms[room][1] != room+1 {
		return true
	}
	return false
}

func (s state23)haveWon() bool {
	for i,room := range s.rooms {
		for _,occupant := range room {
			if occupant != i+1 {
				return false
			}
		}
	}
	return true
}

func (s state23)String() string {
	cl := ".ABCD"
	return fmt.Sprintf("#############\n#%c%c%c%c%c%c%c%c%c%c%c#\n###%c#%c#%c#%c###\n  #%c#%c#%c#%c#\n  #########",
		cl[s.hallway[0]],
		cl[s.hallway[1]],
		cl[s.hallway[2]],
		cl[s.hallway[3]],
		cl[s.hallway[4]],
		cl[s.hallway[5]],
		cl[s.hallway[6]],
		cl[s.hallway[7]],
		cl[s.hallway[8]],
		cl[s.hallway[9]],
		cl[s.hallway[10]],
		cl[s.rooms[0][0]],
		cl[s.rooms[1][0]],
		cl[s.rooms[2][0]],
		cl[s.rooms[3][0]],
		cl[s.rooms[0][1]],
		cl[s.rooms[1][1]],
		cl[s.rooms[2][1]],
		cl[s.rooms[3][1]],
	)
}

func canMoveToHallway(hallway *[11]int, hallwaySpot, room int) bool {
	start := iMin(hallwaySpot, (room*2)+2)
	end := iMax(hallwaySpot, (room*2)+2)
	for i := start; i <= end; i++ {
		if hallway[i] != 0 {
			return false
		}
	}
	return true
}

var count = 0

var validHallSpots = [7]int{0,1,3,5,7,9,10}

var day23Memory = make(map[state23]int)

var didPrint = false

func day23Cost(state state23) int {
	state.panicInvalid()
	result, ok := day23Memory[state]
	if ok {
		return result
	}
	count++
	//if count == 288020 {
	//	fmt.Println(state)
	//}
	if state.haveWon() {
		if !didPrint {
			fmt.Println(state)
			didPrint = true
		}

		return state.cost
	}
	min := math.MaxInt
	for _,spot := range validHallSpots {
		//try moving it back into a room
		if state.hallway[spot] != 0 {
			room := state.hallway[spot]-1
			canMoveBack := true
			for _,occupant := range state.rooms[room] {
				if occupant != 0 && occupant != room+1 {
					canMoveBack = false
				}
			}
			state.hallway[spot] = 0
			hallwayClear := canMoveToHallway(&state.hallway, spot, room)
			state.hallway[spot] = room+1
			if canMoveBack && hallwayClear{
				nextState := state
				nextState.hallway[spot] = 0
				roomDepth := 1
				if nextState.rooms[room][1] != 0 {
					roomDepth = 0
				}
				nextState.rooms[room][roomDepth] = state.hallway[spot]
				//fmt.Printf("put 1 back!\n")
				distance := intABS(spot-((room*2)+2)) + roomDepth + 1
				nextState.cost += int(math.Pow10(state.hallway[spot]-1)) * distance
				cost := day23Cost(nextState)
				if cost < min {
					min = cost
				}
			}
			continue
		}

		//try moving things into the hallway
		for room := 0; room < len(state.rooms); room++ {
			if !canMoveToHallway(&state.hallway, spot, room) {
				continue
			}
			for i,occupant := range state.rooms[room] {
				if occupant > 0 {
					if !state.canLeave(room, i) {
						break
					}
					nextState := state
					nextState.hallway[spot] = occupant
					nextState.rooms[room][i] = 0
					distance := intABS(spot-((room*2)+2)) + i + 1
					nextState.cost += int(math.Pow10(occupant-1)) * distance
					cost := day23Cost(nextState)
					if cost < min {
						min = cost
					}
					break
				}
			}
		}
	}
	day23Memory[state] = min
	return min
}

func adventDay23A(path string) {
	state := state23{
		hallway: [11]int{},
		rooms:   [4][2]int{{4,2},{4,1},{3,1},{2,3}},
		cost:    0,
	}

	//test
	//state := state23{
	//	hallway: [11]int{},
	//	rooms:   [4][2]int{{2,1},{3,4},{2,3},{4,1}},
	//	cost:    0,
	//}

	fmt.Printf("%s\n", state)
	minCost := day23Cost(state)
	fmt.Printf("minimum cost is %d\n", minCost)
	fmt.Printf("called %d times\n", count)
}

type state23b struct {
	hallway [11]int
	rooms [4][4]int
	cost int
}


func (s state23b)canLeave(room, depth int) bool {
	if s.rooms[room][depth] != room+1 {
		return true
	}
	for i := depth+1; i < 4; i++ {
		if s.rooms[room][i] != room+1 {
			return true
		}
	}
	return false
}

func (s state23b)haveWon() bool {
	for i,room := range s.rooms {
		for _,occupant := range room {
			if occupant != i+1 {
				return false
			}
		}
	}
	return true
}

func (s state23b)String() string {
	cl := ".ABCD"
	return fmt.Sprintf("#############\n#%c%c%c%c%c%c%c%c%c%c%c#\n###%c#%c#%c#%c###\n  #%c#%c#%c#%c#\n  #%c#%c#%c#%c#\n  #%c#%c#%c#%c#\n  #########",
		cl[s.hallway[0]],
		cl[s.hallway[1]],
		cl[s.hallway[2]],
		cl[s.hallway[3]],
		cl[s.hallway[4]],
		cl[s.hallway[5]],
		cl[s.hallway[6]],
		cl[s.hallway[7]],
		cl[s.hallway[8]],
		cl[s.hallway[9]],
		cl[s.hallway[10]],
		cl[s.rooms[0][0]],
		cl[s.rooms[1][0]],
		cl[s.rooms[2][0]],
		cl[s.rooms[3][0]],

		cl[s.rooms[0][1]],
		cl[s.rooms[1][1]],
		cl[s.rooms[2][1]],
		cl[s.rooms[3][1]],

		cl[s.rooms[0][2]],
		cl[s.rooms[1][2]],
		cl[s.rooms[2][2]],
		cl[s.rooms[3][2]],

		cl[s.rooms[0][3]],
		cl[s.rooms[1][3]],
		cl[s.rooms[2][3]],
		cl[s.rooms[3][3]],
	)
}


var day23bMemory = make(map[state23b]int)


func day23bCost(state state23b) int {
	result, ok := day23bMemory[state]
	if ok {
		return result
	}
	//count++
	//if count == 288020 {
	//	fmt.Println(state)
	//}
	if state.haveWon() {
		return state.cost
	}
	min := math.MaxInt
	for _,spot := range validHallSpots {
		//try moving it back into a room
		if state.hallway[spot] != 0 {
			room := state.hallway[spot]-1
			canMoveBack := true
			spotsAvailable := 0
			for _,occupant := range state.rooms[room] {
				if occupant != 0 && occupant != room+1 {
					canMoveBack = false
				}
				if occupant == 0 {
					spotsAvailable++
				}
			}
			state.hallway[spot] = 0
			hallwayClear := canMoveToHallway(&state.hallway, spot, room)
			state.hallway[spot] = room+1
			if canMoveBack && hallwayClear && spotsAvailable > 0{
				nextState := state
				nextState.hallway[spot] = 0
				roomDepth := 3
				for nextState.rooms[room][roomDepth] != 0 {
					roomDepth--
				}
				nextState.rooms[room][roomDepth] = state.hallway[spot]
				//fmt.Printf("put 1 back!\n")
				distance := intABS(spot-((room*2)+2)) + roomDepth + 1
				nextState.cost += int(math.Pow10(state.hallway[spot]-1)) * distance
				cost := day23bCost(nextState)
				if cost < min {
					min = cost
				}
			}
			continue
		}

		//try moving things into the hallway
		for room := 0; room < len(state.rooms); room++ {
			if !canMoveToHallway(&state.hallway, spot, room) {
				continue
			}
			for i,occupant := range state.rooms[room] {
				if occupant > 0 {
					if !state.canLeave(room, i) {
						break
					}
					nextState := state
					nextState.hallway[spot] = occupant
					nextState.rooms[room][i] = 0
					distance := intABS(spot-((room*2)+2)) + i + 1
					nextState.cost += int(math.Pow10(occupant-1)) * distance
					cost := day23bCost(nextState)
					if cost < min {
						min = cost
					}
					break
				}
			}
		}
	}
	day23bMemory[state] = min
	return min
}




func adventDay23B(path string) {
	state := state23b{
		hallway: [11]int{},
		rooms:   [4][4]int{{4,4,4,2},{4,3,2,1},{3,2,1,1},{2,1,3,3}},
		cost:    0,
	}

	////test
	//state := state23b{
	//	hallway: [11]int{},
	//	rooms:   [4][4]int{{2,4,4,1},{3,3,2,4},{2,2,1,3},{4,1,3,1}},
	//	cost:    0,
	//}

	fmt.Printf("%s\n", state)
	minCost := day23bCost(state)
	fmt.Printf("minimum cost is %d\n", minCost)

}
