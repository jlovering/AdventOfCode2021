package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type deltaPoint struct {
	dx int
	dy int
}

func printGrid(grid [][]int) {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == 0 {
				util.Dprintf("%2d ", grid[j][i])
			} else {
				util.Dprintf("%2d ", grid[j][i])
			}
		}
		util.Dprintf("\n")
	}
	util.Dprintf("\n")
}

func checkAndFlashAt(grid [][]int, i, j int, visited [][]bool, indent string) int {
	deltas := []deltaPoint{
		{dx: -1, dy: -1},
		{dx: -1, dy: 0},
		{dx: -1, dy: 1},
		{dx: 0, dy: -1},
		{dx: 0, dy: 1},
		{dx: 1, dy: -1},
		{dx: 1, dy: 0},
		{dx: 1, dy: 1},
	}
	flashes := 0
	if grid[j][i] > 9 {
		flashes++
		util.Dprintf(indent+"%d %d flashes\n", i, j)
		printGrid(grid)
		for _, p := range deltas {
			if j+p.dy >= 0 && j+p.dy < len(grid) && i+p.dx >= 0 && i+p.dx < len(grid[0]) {
				if grid[j+p.dy][i+p.dx] != 0 {
					grid[j+p.dy][i+p.dx]++
				}
				util.Dprintf(indent+"%d %d inc: %d\n", i+p.dx, j+p.dy, grid[j+p.dy][i+p.dx])
			}
		}
		for _, p := range deltas {
			if j+p.dy >= 0 && j+p.dy < len(grid) && i+p.dx >= 0 && i+p.dx < len(grid[0]) && !visited[j+p.dy][i+p.dx] {
				visited[j][i] = true
				util.Dprintf(indent+"%d %d check\n", i+p.dx, j+p.dy)
				flashes += checkAndFlashAt(grid, i+p.dx, j+p.dy, visited, indent+"\t")
			}
		}
		grid[j][i] = 0
	}
	return flashes
}

func computeStep(grid [][]int) int {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			grid[j][i]++
		}
	}
	flashes := 0
	visted := make([][]bool, 10)
	for i := range visted {
		visted[i] = make([]bool, 10)
	}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			util.Dprintf("For %d %d\n", i, j)
			flashes += checkAndFlashAt(grid, i, j, visted, "\t")
		}
	}
	return flashes
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}

	j := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		for i, c := range line {
			grid[j][i], err = strconv.Atoi(string(c))
			util.Check_error(err)
		}
		j++
	}

	sumFlashes := 0
	for i := 1; i <= 100; i++ {
		util.Dprintf("Step: %d\n", i)
		sumFlashes += computeStep(grid)
		printGrid(grid)
	}

	return fmt.Sprintf("%d", sumFlashes)
}
