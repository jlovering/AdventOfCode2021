package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
)

func (v *ventmap) markDiagLine(x1 int, y1 int, x2 int, y2 int) {
	i, j := x1, y1
	xstep := 1
	ystep := 1
	if x1 > x2 {
		xstep = -1
	}
	if y1 > y2 {
		ystep = -1
	}
	for i != x2 && j != y2 {
		util.Dprintf("%d,%d\n", j, i)
		v.ventgrid[i][j]++
		i += xstep
		j += ystep
	}
	v.ventgrid[x2][y2]++
}

func Part2(filename string) string {
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
		if x1 == x2 {
			vents.markVertLine(x1, y1, y2)
		} else if y1 == y2 {
			vents.markHorizLine(y1, x1, x2)
		} else if math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
			vents.markDiagLine(x1, y1, x2, y2)
		} else {
			panic(line)
		}
		vents.printVentMap()
	}

	return fmt.Sprintf("%d", vents.countGreaterThanOrEqual(2))
}
