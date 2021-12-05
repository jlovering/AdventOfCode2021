package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

type ventmap struct {
	ventgrid [1000][1000]int
}

func (v *ventmap) markHorizLine(j int, start int, end int) {
	if end < start {
		tmp := end
		end = start
		start = tmp
	}
	for i := start; i <= end; i++ {
		v.ventgrid[i][j]++
	}
}

func (v *ventmap) markVertLine(i int, start int, end int) {
	if end < start {
		tmp := end
		end = start
		start = tmp
	}
	for j := start; j <= end; j++ {
		v.ventgrid[i][j]++
	}
}

func (v *ventmap) printVentMap() {
	ilim := len(v.ventgrid[0])
	if ilim > 10 {
		ilim = 10
	}
	jlim := len(v.ventgrid[0])
	if jlim > 10 {
		jlim = 10
	}
	for j := 0; j < ilim; j++ {
		for i := 0; i < ilim; i++ {
			util.Dprintf("%d ", v.ventgrid[i][j])
		}
		util.Dprintf("\n")
	}
	util.Dprintf("\n")
}

func (v *ventmap) countGreaterThanOrEqual(thrs int) int {
	count := 0
	for i := 0; i < len(v.ventgrid[0]); i++ {
		for j := 0; j < len(v.ventgrid[i]); j++ {
			if v.ventgrid[i][j] >= thrs {
				count++
			}
		}
	}
	return count
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var vents ventmap
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var x1, y1, x2, y2 int
		_, err = fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		util.Check_error(err)
		util.Dprintf("%d,%d -> %d,%d\n", x1, y1, x2, y2)
		if x1 != x2 && y1 != y2 {
			continue
		}
		if x1 == x2 {
			vents.markVertLine(x1, y1, y2)
		} else if y1 == y2 {
			vents.markHorizLine(y1, x1, x2)
		}
		vents.printVentMap()
	}

	return fmt.Sprintf("%d", vents.countGreaterThanOrEqual(2))
}
