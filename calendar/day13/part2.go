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

func Part2(filename string) string {
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
	var foldGrid [][]bool = fullGrid
	for _, fi := range foldInstructions {
		switch fi.axis {
		case "y":
			fmt.Printf("Folding Y at %d\n", fi.location)
			foldGrid = foldY(foldGrid, fi.location)
		case "x":
			fmt.Printf("Folding X at %d\n", fi.location)
			foldGrid = foldX(foldGrid, fi.location)
		}
		//util.Dprintf("%s\n", arrayutil.SPrintArrayXY(foldGrid, "%6v"))
		//foldGrid = foldX(foldGrid, 5)
		//util.Dprintf("%s\n", arrayutil.SPrintArrayXY(foldGrid, "%6v"))
	}

	for y := range foldGrid[0] {
		for x := range foldGrid {
			if foldGrid[x][y] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	return ""
}
