package main

import "fmt"

func adventDay25A(path string) {
	lines := readLines(path)

	width := len(lines[0])
	height := len(lines)
	grid := make([]rune, width*height)

	for i,line := range lines {
		for j,r := range line {
			grid[i*width + j] = r
		}
	}

	nextGrid := make([]rune, width*height)


	steps := 0
	nmoved := 1
	for nmoved > 0 {
		//for i := 0; i < height; i++ {
		//	for j := 0; j < width; j++ {
		//		fmt.Printf("%c", grid[i*width + j])
		//	}
		//	fmt.Printf("\n")
		//}
		//fmt.Printf("\n")


		nmoved = 0
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				nextGrid[i*width+j] = '.'
			}
		}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if grid[i*width + j] == '>' {
					jIndex := (j + 1) % width
					if grid[i*width + jIndex] == '.' {
						nmoved++
						nextGrid[i*width + jIndex] = '>'
					} else {
						nextGrid[i*width + j] = '>'
					}
				}
				if grid[i*width + j] == 'v' {
					nextGrid[i*width+j] = 'v'
				}
			}
		}
		nextGrid, grid = grid,nextGrid
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				nextGrid[i*width+j] = '.'
			}
		}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if grid[i*width + j] == 'v' {
					iIndex := (i + 1) % height
					if grid[iIndex*width + j] == '.' {
						nmoved++
						nextGrid[iIndex*width + j] = 'v'
					} else {
						nextGrid[i*width + j] = 'v'
					}
				}
				if grid[i*width + j] == '>' {
					nextGrid[i*width+j] = '>'
				}

			}
		}
		nextGrid, grid = grid,nextGrid

		steps++
	}
	//for i := 0; i < height; i++ {
	//	for j := 0; j < width; j++ {
	//		fmt.Printf("%c", grid[i*width + j])
	//	}
	//	fmt.Printf("\n")
	//}
	//fmt.Printf("\n")

			fmt.Printf("")
	fmt.Printf("stopped after %d steps\n", steps)
}

func adventDay25B(path string) {

}
