package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printGrid(grid [][]int) {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			util.Dprintf("%d ", grid[j][i])
		}
		util.Dprintf("\n")
	}
}

type deltaPoint struct {
	dx int
	dy int
}

func checkLowPoint(grid [][]int, j, i int) bool {
	deltas := []deltaPoint{
		{dx: -1, dy: 0},
		{dx: 1, dy: 0},
		{dx: 0, dy: -1},
		{dx: 0, dy: 1},
	}

	//util.Dprintf("%v\n", deltas)
	for _, d := range deltas {
		if (i+d.dx) >= 0 && (i+d.dx) < len(grid[0]) && (j+d.dy) >= 0 && (j+d.dy) < len(grid) {
			//util.Dprintf("\t%d %d %d %d check\n", i+d.dx, j+d.dy, grid[j][i], grid[j+d.dy][i+d.dx])
			if grid[j][i] >= grid[j+d.dy][i+d.dx] {
				return false
			}
		}
	}
	return true
}

func findLowPoints(grid [][]int) []int {
	printGrid(grid)
	var ret []int = make([]int, 0, len(grid)*len(grid[0]))
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if checkLowPoint(grid, j, i) {
				util.Dprintf("%d %d %d is lowest\n", i, j, grid[j][i])
				ret = append(ret, grid[j][i])
			}
		}
	}
	return ret
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var floorMap [][]int = make([][]int, 0, 1000)

	j := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		floorMap = append(floorMap, make([]int, len(line)))
		for i, v := range line {
			floorMap[j][i], err = strconv.Atoi(string(v))
			util.Check_error(err)
		}
		j++
	}

	printGrid(floorMap)
	lows := findLowPoints(floorMap)

	sum := 0
	for _, v := range lows {
		sum += v + 1
	}
	return fmt.Sprintf("%d", sum)
}
