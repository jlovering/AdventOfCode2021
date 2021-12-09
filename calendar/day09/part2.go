package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x int
	y int
}

func printGridBool(grid [][]bool) {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			util.Dprintf("%-5v ", grid[j][i])
		}
		util.Dprintf("\n")
	}
}

func findLowPointsCoords(grid [][]int) []point {
	printGrid(grid)
	var ret []point = make([]point, 0, len(grid)*len(grid[0]))
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if checkLowPoint(grid, j, i) {
				//util.Dprintf("%d %d %d is lowest\n", i, j, grid[j][i])
				ret = append(ret, point{x: i, y: j})
			}
		}
	}
	return ret
}

func floodBasin(grid [][]int, start point, visited [][]bool) int {
	deltas := []deltaPoint{
		{dx: -1, dy: 0},
		{dx: 1, dy: 0},
		{dx: 0, dy: -1},
		{dx: 0, dy: 1},
	}

	i := start.x
	j := start.y

	visited[j][i] = true
	util.Dprintf("%d %d\n", i, j)
	printGridBool(visited)
	util.Dprintf("\n")
	sumAll := 0
	for _, d := range deltas {
		if (i+d.dx) >= 0 && (i+d.dx) < len(grid[0]) && (j+d.dy) >= 0 && (j+d.dy) < len(grid) && !visited[j+d.dy][i+d.dx] {
			if grid[j+d.dy][i+d.dx] == 9 {
				continue
			} else {
				sumAll += floodBasin(grid, point{x: i + d.dx, y: j + d.dy}, visited)
			}
		}
	}
	return sumAll + 1
}

func Part2(filename string) string {
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
	lows := findLowPointsCoords(floorMap)

	var basinsSizes []int = make([]int, 0, len(floorMap))
	var visited [][]bool = make([][]bool, len(floorMap))
	for j := range visited {
		visited[j] = make([]bool, len(floorMap[0]))
	}
	util.Dprintf("%v\n", lows)
	for _, v := range lows {
		basinsSizes = append(basinsSizes, floodBasin(floorMap, v, visited))
		util.Dprintf("%v\n\n", basinsSizes)
	}
	sort.Ints(basinsSizes)
	util.Dprintf("%v\n", basinsSizes)

	product := 1
	for _, v := range basinsSizes[len(basinsSizes)-3:] {
		util.Dprintf("%v ", v)
		product *= v
	}
	util.Dprintf("\n")
	return fmt.Sprintf("%d", product)
}
