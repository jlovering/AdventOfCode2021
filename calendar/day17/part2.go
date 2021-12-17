package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	file_scanner.Scan()
	line := file_scanner.Text()
	trg := target{}
	_, err = fmt.Sscanf(line, "target area: x=%d..%d, y=%d..%d", &trg.tl.x, &trg.br.x, &trg.br.y, &trg.tl.y)
	util.Check_error(err)

	util.Dprintf("%v\n", trg)
	//loc, cur := point{0, 0}, point{0, 0}

	hitters := make([]point, 0)
	for x := 0; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			_, hit := simulateFire(point{x, y}, trg)
			util.Dprintf("\n")
			if hit {
				hitters = append(hitters, point{x, y})
				util.Dprintf("Hit Pair: %d %d\n", x, y)
			}
		}
	}

	util.Dprintf("%v\n", hitters)
	return fmt.Sprintf("%d", len(hitters))
}
