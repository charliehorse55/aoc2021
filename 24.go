package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

const (
	 xReg = iota
	 yReg
	 zReg
	 wReg
)

func Inp(a, b int) int {
	return b
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a*b
}

func Div(a,b int) int {
	return a/b
}

func Mod(a,b int) int {
	return a % b
}

func Eql(a,b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}
}

var operations = map[string]operation{
	"inp": Inp,
	"add": Add,
	"mul": Mul,
	"div": Div,
	"mod": Mod,
	"eql": Eql,
}

var regNames = map[string]int{
	"x": xReg,
	"y": yReg,
	"z": zReg,
	"w": wReg,
}

type operation func(int, int) int

type instruction struct {
	op operation
	dest int
	b int
	bLiteral bool
	isInput bool
}

type aluState struct {
	regs [4]int
}

func (l *aluState)Run(program []instruction, input []int) {
	pos := 0
	for i,ins := range program {
		aVal := l.regs[ins.dest]
		var bVal int
		if ins.bLiteral {
			bVal = ins.b
		} else if ins.isInput {
			bVal = input[pos]
			pos++
		} else {
			bVal = l.regs[ins.b]
		}
		l.regs[ins.dest] = ins.op(aVal, bVal)
		if l.regs[ins.dest] < 0 && (ins.dest == zReg || ins.dest == yReg) {
			fmt.Printf("negative value in reg %d (ins %d)\n", ins.dest, i+1)
			panic("")
		}
	}
}

func makeInput(in string) [14]int {
	if len(in) != 14 {
		panic("wrong length input string!")
	}

	var out [14]int
	for i,r := range in {
		out[i], _ = strconv.Atoi(string(r))
	}
	return out
}

func inputToString(in [14]int) string {
	sb := strings.Builder{}
	for _,v := range in {
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	return sb.String()
}

func inputToNumber(in [14]int) int {
	x := 0
	for _,val := range in {
		x *= 10
		x += val
	}
	return x
}

func makeRandomInput() [14]int {
	var out [14]int
	for i := range out {
		out[i] = rand.Intn(9) + 1
	}
	return out
}

func adventDay24A(path string) {
	lines := readLines(path)

	var program []instruction
	var xOffsets []int
	divsLeft := 7
	divsLeftPerRound := make([]int, 14)

	for i,line := range lines {
		f := strings.Fields(line)
		op := f[0]
		dest := regNames[f[1]]
		ins := instruction{
			op:       operations[op],
			dest:     dest,
			b:        0,
			bLiteral: false,
		}
		if op == "inp" {
			ins.isInput = true
		} else {
			reg, ok := regNames[f[2]]
			if !ok {
				reg, _ = strconv.Atoi(f[2])
				ins.bLiteral = true
			}
			ins.b = reg
		}
		if i % 18 == 5 {
			xOffsets = append(xOffsets, ins.b)
		}
		if i % 18 == 4 {
			if lines[i][5:] != lines[i%18][5:] {
				divsLeft--
			}
			divsLeftPerRound[i/18] = divsLeft
		}

		program = append(program, ins)
	}
	fmt.Printf("divs left: %v\n", divsLeftPerRound)



	//input := [14]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1}
//makeInput("13579246899999")
//	rand.Seed(time.Now().UnixNano())
//	valid := 0
//	const iters = 1000
//	const nthread = 100
//
//	type result struct {
//		min int
//		minInput [14]int
//	}
//
//	results := make(chan result, nthread)
//
//	for thread := 0; thread < nthread; thread++ {
//		go func() {
//			min := math.MaxInt
//			var minInput [14]int
//			for i := 0; i < iters; i++ {
//				input := makeRandomInput()
//				//input[0] = 6
//				//input[1] = 1
//				//input[2] = 1
//				var alu aluState
//				alu.Run(program, input)
//				if alu.regs[zReg] == 0 {
//					valid++
//					fmt.Printf("%v\n", input)
//				}
//				if alu.regs[zReg] < min {
//					min = alu.regs[zReg]
//					minInput = input
//				}
//			}
//			results <- result{
//				min:      min,
//				minInput: minInput,
//			}
//		}()
//	}
//	min := math.MaxInt
//	var minInput [14]int
//	for i := 0; i < nthread; i++ {
//		r := <- results
//		if r.min < min {
//			min = r.min
//			minInput = r.minInput
//		}
//	}
//
//	fmt.Printf("found %d/%d (%.0f%%)\n", valid, iters, float64(valid)/iters)
//	fmt.Printf("min input was %s for %d\n", inputToString(minInput), min)

	type progress struct {
		aluState
		input [14]int
	}

	states := make(map[progress]bool)

	states[progress{
		aluState: aluState{},
		input:    [14]int{},
	}] = true

	//fmt.Printf("%d states\n", len(states))
	for round := 0; round < 14; round++ {
		nextRound := make(map[progress]bool)
		//var minState aluState
		//minState.regs[zReg] = math.MaxInt
		for progOrig := range states {
			for i := 1; i <= 9; i++ {
				prog := progOrig
				prog.input[round] = i
				prog.Run(program[round*18:(round+1)*18], prog.input[round:])
				//save := true
				if prog.regs[zReg] < int(math.Pow(26, float64(divsLeftPerRound[round]))) {
					//nextX := (prog.regs[zReg] % 26) + xOffsets[round+1]
					//if nextX < 1 || nextX > 9 {
					//	save = false
					//}
					nextRound[prog] = true
				}
				//if save {
				//	mod := prog.regs[zReg] % 26
				//	nextRound[mod] = prog
				//}
			}
		}

		states = nextRound
		fmt.Printf("after %d rounds, %d states\n", round+1, len(states))
	}
	fmt.Printf("found %d results\n", len(states))
	//for val := range states {
	//	fmt.Printf("%s\n", inputToString(val.input))
	//}
	//fmt.Printf("pruning non-zeros\n")
	//for key,prog := range states {
	//	if prog.regs[zReg] != 0 {
	//		delete(states, key)
	//	}
	//}

	fmt.Printf("finding max!\n")
	max := 0
	min := math.MaxInt
	for val := range states {
		x := inputToNumber(val.input)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	fmt.Printf("max = %d\n", max)
	fmt.Printf("min = %d\n", min)
	//fmt.Printf("states: %v\n", states)

	//minInput := makeInput("11778183565248")
	//
	//for digit := 0; digit < 14; digit++ {
	//	min := -(minInput[digit]-1)
	//	max := 9-(minInput[digit])
	//	for i := min; i <= max; i++ {
	//		input := minInput
	//		input[digit] += i
			input := makeInput("29578186816119")
			var alu aluState
			alu.Run(program, input[:])
			fmt.Printf("zReg = %d\n", alu.regs[zReg])
			//fmt.Printf("%2d-%d: %d\n", digit, minInput[digit]+i, alu.regs[zReg])
	//	}
	//	fmt.Printf("\n")
	//}

}


func adventDay24B(path string) {
	//lines := readLines(path)
	//
	//divsLeft := 7
	//divsLeftPerRound := make([]int, 14)
	//
	//for i := 18; i < len(lines); i++ {
	//	if i % 18 == 4 {
	//		if lines[i][5:] != lines[i%18][5:] {
	//			divsLeft--
	//		}
	//		divsLeftPerRound[i/18] = divsLeft
	//	}
	//	if lines[i][:3] != lines[i%18][:3] {
	//		fmt.Printf("different op on line %d\n", i+1)
	//	}
	//	if lines[i][4] != lines[i%18][4] {
	//		fmt.Printf("different argument on line %d\n", (i+1)%18)
	//	}
	//	if lines[i][5:] != lines[i%18][5:] {
	//		fmt.Printf("different 2nd argument on line %3d - (%9s vs %3s)\n", (i+1)%18, lines[i], lines[i%18][5:])
	//	}
	//}
}
