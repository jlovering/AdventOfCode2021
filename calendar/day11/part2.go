package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkGrid(grid [][]int) bool {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] != 0 {
				return false
			}
		}
	}
	return true
}

func Part2(filename string) string {
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

	i := 1
	for !checkGrid(grid) {
		util.Dprintf("Step: %d\n", i)
		computeStep(grid)
		printGrid(grid)
		i++
	}

	return fmt.Sprintf("%d", i-1)
}
