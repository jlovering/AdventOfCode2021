package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

type target struct {
	tl point
	br point
}

func simulateFire(init point, trg target) (int, bool) {
	vel := init
	loc := point{0, 0}
	hit := false
	peakY := 0
	util.Dprintf("%v\n", vel)
	for loc.y >= trg.br.y {
		loc = point{loc.x + vel.x, loc.y + vel.y}
		util.Dprintf("\t%v\n", loc)
		if vel.x > 0 {
			vel.x = vel.x - 1
		} else if vel.x < 0 {
			vel.x = vel.x + 1
		}
		vel.y -= 1
		//if vel.y == 0 && loc.y > -1*trg.br.y {
		//	util.Dprintf("\tToo Fast")
		//	break
		//}
		if loc.y > peakY {
			peakY = loc.y
		}
		if loc.y <= trg.tl.y && loc.y >= trg.br.y && loc.x >= trg.tl.x && loc.x <= trg.br.x {
			util.Dprintf("\t\tHit\n")
			hit = true
			break
		}
	}
	return peakY, hit
}

func findTerminalXSet(trg target) []int {
	x := 0
	xes := make([]int, 0)
	for {
		xf := (x * (x + 1)) / 2
		if xf >= trg.tl.x && xf < trg.br.x {
			util.Dprintf("x: %d, %d\n", x, xf)
			xes = append(xes, x)
		}
		if xf > trg.br.x {
			break
		}
		x++
	}
	return xes
}

func Part1(filename string) string {
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

	xes := findTerminalXSet(trg)
	util.Dprintf("xes: %v\n", xes)

	x := 0
	maxY := 0
	for _, x = range xes {
		for y := 0; y < 100; y++ {
			peakY, hit := simulateFire(point{x, y}, trg)
			util.Dprintf("\n")
			if hit && peakY > maxY {
				maxY = peakY
				util.Dprintf("Hit Pair: %d %d\n", x, y)
			}
		}
	}

	fmt.Printf("%d %d\n", x, maxY)
	return fmt.Sprintf("%d", maxY)
}
