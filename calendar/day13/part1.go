package adventofcode

import (
	arrayutil "adventofcode/util/array"
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fold struct {
	axis     string
	location int
}

func foldY(grid [][]bool, foldLine int) [][]bool {
	foldedGrid := arrayutil.SliceBuilder2DBool(len(grid), foldLine)
	ny := foldLine - 1
	for y := foldLine + 1; y < len(grid[0]); y++ {
		if ny < 0 {
			panic(1)
		}
		for x := 0; x < len(grid); x++ {
			foldedGrid[x][ny] = grid[x][ny] || grid[x][y]
		}
		ny--
	}
	return foldedGrid
}

func foldX(grid [][]bool, foldLine int) [][]bool {
	foldedGrid := arrayutil.SliceBuilder2DBool(foldLine, len(grid[0]))
	nx := foldLine - 1
	for x := foldLine + 1; x < len(grid); x++ {
		if nx < 0 {
			panic(1)
		}
		for y := 0; y < len(grid[0]); y++ {
			foldedGrid[nx][y] = grid[nx][y] || grid[x][y]
		}
		nx--
	}
	return foldedGrid
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	tooBigGrid := arrayutil.SliceBuilder2DBool(2000, 2000)
	foldInstructions := []fold{}

	maxX := 0
	maxY := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		if strings.Contains(line, "fold along") {
			sep := strings.Split(line, " ")
			sepsep := strings.Split(sep[2], "=")
			axis := sepsep[0]
			value, err := strconv.Atoi(sepsep[1])
			util.Check_error(err)
			foldInstructions = append(foldInstructions, fold{axis: axis, location: value})
		} else if line != "" {
			sep := strings.Split(line, ",")
			x, err := strconv.Atoi(sep[0])
			util.Check_error(err)
			y, err := strconv.Atoi(sep[1])
			util.Check_error(err)
			tooBigGrid[x][y] = true
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	fmt.Printf("Input\n")
	fullGrid := arrayutil.SliceBuilder2DBool(maxX+1, maxY+1)

	for i := range fullGrid {
		for j := range fullGrid[i] {
			fullGrid[i][j] = tooBigGrid[i][j]
		}
	}
	fmt.Printf("Copied\n")

	//util.Dprintf("%s\n", arrayutil.SPrintArrayXY(fullGrid, "%6v"))
	util.Dprintf("%v\n", foldInstructions)
	fmt.Printf("Fold\n")
	var foldGrid [][]bool
	switch foldInstructions[0].axis {
	case "y":
		fmt.Printf("Folding Y at %d\n", foldInstructions[0].location)
		foldGrid = foldY(fullGrid, foldInstructions[0].location)
	case "x":
		fmt.Printf("Folding X at %d\n", foldInstructions[0].location)
		foldGrid = foldX(fullGrid, foldInstructions[0].location)
	}
	//util.Dprintf("%s\n", arrayutil.SPrintArrayXY(foldGrid, "%6v"))
	//foldGrid = foldX(foldGrid, 5)
	//util.Dprintf("%s\n", arrayutil.SPrintArrayXY(foldGrid, "%6v"))

	count := 0
	for i := range foldGrid {
		for j := range foldGrid[i] {
			if foldGrid[i][j] {
				count++
			}
		}
	}
	return fmt.Sprintf("%d", count)
}
