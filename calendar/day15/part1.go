package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	astar "github.com/beefsack/go-astar"
)

type point struct {
	x int
	y int
}

type Tile struct {
	cost    int
	refGrid *map[point]*Tile
	loc     point
}

func (t *Tile) PathNeighbors() []astar.Pather {
	var neigh []astar.Pather = make([]astar.Pather, 0)

	util.Dprintf("{%d %d}: ", t.loc.x, t.loc.y)
	for _, d := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		nP := point{t.loc.x + d.x, t.loc.y + d.y}
		grid := *(t.refGrid)
		if v, e := grid[nP]; e {
			neigh = append(neigh, v)
		}
	}

	for _, n := range neigh {
		t := n.(*Tile)
		util.Dprintf("{%d %d} ", t.loc.x, t.loc.y)
	}
	util.Dprintf("\n")

	return neigh
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return float64(toT.cost)
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return math.Abs(float64(toT.loc.x-t.loc.x)) + math.Abs(float64(toT.loc.y-t.loc.y))
}

func parseInput(file_scanner *bufio.Scanner) (map[point]*Tile, point) {
	var grid = make(map[point]*Tile)
	i := 0
	lineLen := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		lineLen = len(line)
		for j, c := range line {
			val, err := strconv.Atoi(string(c))
			util.Check_error(err)
			grid[point{j, i}] = &Tile{cost: val, refGrid: &grid, loc: point{j, i}}
		}
		i++
	}
	return grid, point{i - 1, lineLen - 1}
}

func printGrid(grid map[point]*Tile, mxy point) {
	for j := 0; j <= mxy.y; j++ {
		for i := 0; i <= mxy.x; i++ {
			p := point{i, j}
			if t, e := grid[p]; e {
				util.Dprintf("%d ", t.cost)
			} else {
				util.Dprintf("   ")
			}
		}
		util.Dprintf("\n")
	}
}

func printPatherSlice(ps []astar.Pather) {
	for _, n := range ps {
		t := n.(*Tile)
		util.Dprintf("{%d %d} ", t.loc.x, t.loc.y)
	}
	util.Dprintf("\n")
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	grid, mxy := parseInput(file_scanner)

	printGrid(grid, mxy)

	path, distance, found := astar.Path(grid[point{0, 0}], grid[mxy])

	if !found {
		panic(1)
	}

	printPatherSlice(path)
	util.Dprintf("%f\n", distance)

	return fmt.Sprintf("%d", int(distance))
}
