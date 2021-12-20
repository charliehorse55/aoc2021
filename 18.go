package main

import (
	"fmt"
	"strconv"
)

type pair struct {
	parent *pair
	left fish
	right fish
}


func (p pair)String() string {
	return fmt.Sprintf("[%s,%s]", p.left, p.right)
}

func (p *pair)AddRight(x int) {
	p.right.AddRight(x)
}

func (p *pair)AddLeft(x int) {
	p.left.AddLeft(x)
}

func (p *pair)Value() int {
	return 3*p.left.Value() + 2*p.right.Value()
}

func (p *pair)SplitIfAble() bool {
	if p.left.SplitIfAble() {
		return true
	}
	if p.right.SplitIfAble() {
		return true
	}
	return false
}

func (p *pair)ExplodeIfAble(depth int) bool {
	if depth >= 4 {
		//find left
		parent := p.parent
		this := p
		for parent != nil {
			if parent.left != this {
				parent.left.AddRight(p.left.Value())
				break
			}
			parent, this = parent.parent, parent
		}
		//find right
		parent = p.parent
		this = p
		for parent != nil {
			if parent.right != this {
				parent.right.AddLeft(p.right.Value())
				break
			}
			parent, this = parent.parent, parent
		}

		newVal := intVal{
			parent: p.parent,
			val:    0,
		}
		if p.parent.left == p {
			p.parent.left = &newVal
		} else {
			p.parent.right = &newVal
		}
		return true
	}
	if p.left.ExplodeIfAble(depth+1) {
		return true
	}
	if p.right.ExplodeIfAble(depth+1) {
		return true
	}
	return false
}

func (p *pair)Reduce() {
	for {
		if p.ExplodeIfAble(0) {
			continue
		}
		if p.SplitIfAble() {
			continue
		}
		break
	}
}

type intVal struct {
	parent *pair
	val int

}

func (i *intVal)ExplodeIfAble(depth int) bool {
	return false
}

func (i *intVal)SplitIfAble() bool {
	if i.val >= 10 {
		p := &pair{
			parent: i.parent,
		}
		p.left = &intVal{
			parent: p,
			val:    i.val/2,
		}
		p.right = &intVal{
			parent: p,
			val:    (i.val+1)/2,
		}
		if i.parent.left == i {
			i.parent.left = p
		} else {
			i.parent.right = p
		}
		return true
	}
	return false
}

func (i *intVal)String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i *intVal)AddLeft(x int) {
	i.val += x
}
func (i *intVal)AddRight(x int) {
	i.val += x
}

func (i *intVal)Value() int {
	return i.val
}

type fish interface {
	ExplodeIfAble(depth int) bool
	SplitIfAble() bool
	Value() int
	AddRight(x int)
	AddLeft(x int)
	String() string
}

func parsePair(line string) *pair {
	var currPair *pair
	for _,r := range line {
		switch r {
		case '[':
			currPair = &pair{
				parent: currPair,
			}
		case ']':
			if currPair.parent != nil {
				if currPair.parent.left == nil {
					currPair.parent.left = currPair
				} else {
					currPair.parent.right = currPair
				}
				currPair = currPair.parent
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, _ := strconv.Atoi(string(r))
			ival := intVal{
				parent: currPair,
				val:    val,
			}
			if currPair.left == nil {
				currPair.left = &ival
			} else {
				currPair.right = &ival
			}
		}
	}
	return currPair
}

func parsePairs(path string) []*pair {
	lines := readLines(path)

	pairs := make([]*pair, len(lines))
	for i,line := range lines {
		//char := 0
		pairs[i] = parsePair(line)
	}
	return pairs
}

func adventDay18A(path string) {
	pairs := parsePairs(path)
	sum := pairs[0]
	for _,p := range pairs[1:] {
		nextSum := &pair{
			parent: nil,
			left:   sum,
			right:  p,
		}
		sum.parent = nextSum
		p.parent = nextSum
		nextSum.Reduce()
		sum = nextSum
	}
	fmt.Printf("reduced:\n%s\nmagnitude: %d\n", sum, sum.Value())
}

func adventDay18B(path string) {
	lines := readLines(path)
	max := 0
	for _, l1 := range lines {
		for _, l2 := range lines {
			if l1 == l2 {
				continue
			}
			p1 := parsePair(l1)
			p2 := parsePair(l2)
			sum := &pair{
				parent: nil,
				left:   p1,
				right:  p2,
			}
			p1.parent = sum
			p2.parent = sum
			sum.Reduce()
			mag := sum.Value()
			if mag > max {
				max = mag
			}
		}
	}
	fmt.Printf("max is %d\n", max)
}
