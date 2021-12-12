package main

import (
	"fmt"
	"strings"
)

type cave struct {
	name string
	visited int
	passages []string
}

func getCave(index map[string]*cave, name string) *cave {
	c, ok := index[name]
	if ok {
		return c
	} else {
		index[name] = &cave{name:name}
		return index[name]
	}
}

func traverse(index map[string]*cave, at *cave) int {
	score := 0
	for _,path := range at.passages {
		if path == "end" {
			score++
			continue
		}
		nextCave := index[path]
		if !strings.ContainsAny(path, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && nextCave.visited > 0 {
			continue
		}
		nextCave.visited++
		score += traverse(index, nextCave)
		nextCave.visited--
	}
	return score
}

func adventDay12A(path string) {
	paths := readLines(path)

    caves := make(map[string]*cave, len(paths))

	for _,line := range paths {
		parts := strings.Split(line, "-")
		a := getCave(caves, parts[0])
		b := getCave(caves, parts[1])
		a.passages = append(a.passages, parts[1])
		b.passages = append(b.passages, parts[0])
	}
	//for key,val := range caves {
	//	for _,passage := range val.passages {
	//		fmt.Printf("%s -> %s\n", key, passage)
	//	}
	//}
	caves["start"].visited = 1
	fmt.Printf("found %d paths\n", traverse(caves, caves["start"]))

}

func traverseDoubled(index map[string]*cave, at *cave, doubled bool) int {
	score := 0
	for _,path := range at.passages {
		if path == "end" {
			score++
			continue
		}
		nextCave := index[path]
		if !strings.ContainsAny(path, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && nextCave.visited > 0 {
			if !doubled && nextCave.name != "start" {
				nextCave.visited++
				score += traverseDoubled(index, nextCave, true)
				nextCave.visited--
			}
			continue
		}
		nextCave.visited++
		score += traverseDoubled(index, nextCave, doubled)
		nextCave.visited--
	}
	return score
}


func adventDay12B(path string) {
	paths := readLines(path)

	caves := make(map[string]*cave, len(paths))

	for _,line := range paths {
		parts := strings.Split(line, "-")
		a := getCave(caves, parts[0])
		b := getCave(caves, parts[1])
		a.passages = append(a.passages, parts[1])
		b.passages = append(b.passages, parts[0])
	}
	caves["start"].visited = 1
	fmt.Printf("found %d paths\n", traverseDoubled(caves, caves["start"], false))

}