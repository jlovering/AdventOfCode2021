package adventofcode

import (
	arrayutil "adventofcode/util/array"
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func parseInput(file_scanner *bufio.Scanner) [][]rune {
	grid := [][]rune{}
	j := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		grid = append(grid, make([]rune, len(line)))
		for i, r := range line {
			grid[j][i] = r
		}
		j++
	}
	return grid
}

func simulateStep(grid [][]rune) ([][]rune, bool) {
	grid1, east := simulateStepEast(grid)
	grid2, south := simulateStepSouth(grid1)

	return grid2, east || south
}

func simulateStepEast(grid [][]rune) ([][]rune, bool) {
	moved := false
	nGrid := make([][]rune, len(grid))
	for j := 0; j < len(nGrid); j++ {
		nGrid[j] = make([]rune, len(grid[j]))
	}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			//util.Dprintf("%d,%d %c\n", i, j, grid[j][i])
			//util.Dprintf("%s\n", arrayutil.SPrintArrayYX(nGrid, "%c"))
			if grid[j][i] == '>' {
				if (i + 1) >= len(grid[j]) {
					if grid[j][0] == '.' {
						nGrid[j][i] = '.'
						nGrid[j][0] = '>'
						moved = true
					} else {
						nGrid[j][i] = grid[j][i]
					}
				} else {
					if grid[j][i+1] == '.' {
						nGrid[j][i] = '.'
						nGrid[j][i+1] = '>'
						moved = true
					} else {
						nGrid[j][i] = grid[j][i]
					}
				}
			} else {
				if nGrid[j][i] != '>' {
					nGrid[j][i] = grid[j][i]
				}
			}
		}
	}
	return nGrid, moved
}

func simulateStepSouth(grid [][]rune) ([][]rune, bool) {
	moved := false
	nGrid := make([][]rune, len(grid))
	for j := 0; j < len(nGrid); j++ {
		nGrid[j] = make([]rune, len(grid[j]))
	}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == 'v' {
				if (j + 1) >= len(grid) {
					if grid[0][i] == '.' {
						nGrid[j][i] = '.'
						nGrid[0][i] = 'v'
						moved = true
					} else {
						nGrid[j][i] = grid[j][i]
					}
				} else {
					if grid[j+1][i] == '.' {
						nGrid[j][i] = '.'
						nGrid[j+1][i] = 'v'
						moved = true
					} else {
						nGrid[j][i] = grid[j][i]
					}
				}
			} else {
				if nGrid[j][i] != 'v' {
					nGrid[j][i] = grid[j][i]
				}
			}
		}
	}
	return nGrid, moved
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	grid := parseInput(file_scanner)

	util.Dprintf("%d x %d\n", len(grid), len(grid[0]))
	moved := true
	i := 0
	for moved {
		util.Dprintf("Step %d\n%s\n", i, arrayutil.SPrintArrayYX(grid, "%c"))
		grid, moved = simulateStep(grid)
		i++
	}

	return fmt.Sprintf("%d", i)
}
