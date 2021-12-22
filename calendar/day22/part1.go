package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
	z int
}

/*
   p7------p8
  /|      /|
p3------p4 |
|  |    |  |
|  p5---|--p6
| /     | /
p1------p2
y z
|/
x---
*/
type cube struct {
	xmin, xmax int
	ymin, ymax int
	zmin, zmax int
}

type cubeField map[cube]bool

type pointField map[point]bool

func (p point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func (pf *pointField) switchCubes(ci cube, state bool) {
	for x := ci.xmin; x <= ci.xmax; x++ {
		for y := ci.ymin; y <= ci.ymax; y++ {
			for z := ci.zmin; z <= ci.zmax; z++ {
				(*pf)[point{x, y, z}] = state
			}
		}
	}
}

func (pf pointField) countEnabled() uint {
	count := uint(0)
	for _, v := range pf {
		if v {
			count++
		}
	}
	return count
}

func parseInput(file_scanner *bufio.Scanner) pointField {
	pf := pointField{}
	for file_scanner.Scan() {
		line := file_scanner.Text()
		s := strings.Split(line, " ")
		op := s[0]
		xstart, xend := 0, 0
		ystart, yend := 0, 0
		zstart, zend := 0, 0
		var err error

		axis := strings.Split(s[1], ",")
		//util.Dprintf("%v\n", axis)

		vals := strings.SplitN(axis[0][2:], "..", 2)
		xstart, err = strconv.Atoi(vals[0])
		util.Check_error(err)
		xend, err = strconv.Atoi(vals[1])
		util.Check_error(err)

		vals = strings.SplitN(axis[1][2:], "..", 2)
		ystart, err = strconv.Atoi(vals[0])
		util.Check_error(err)
		yend, err = strconv.Atoi(vals[1])
		util.Check_error(err)

		vals = strings.SplitN(axis[2][2:], "..", 2)
		zstart, err = strconv.Atoi(vals[0])
		util.Check_error(err)
		zend, err = strconv.Atoi(vals[1])
		util.Check_error(err)

		//util.Dprintf("x=%d..%d y=%d..%d z=%d..%d", xstart, xend, ystart, yend, zstart, zend)
		if xstart <= -50 {
			xstart = -50
		}
		if xend >= 50 {
			xend = 50
		}
		if ystart <= -50 {
			ystart = -50
		}
		if yend >= 50 {
			yend = 50
		}
		if zstart <= -50 {
			zstart = -50
		}
		if zend >= 50 {
			zend = 50
		}
		c := cube{xstart, xend, ystart, yend, zstart, zend}
		if op == "on" {
			pf.switchCubes(c, true)
		} else if op == "off" {
			pf.switchCubes(c, false)
		}
	}
	return pf
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	cf := parseInput(file_scanner)

	return fmt.Sprintf("%d", cf.countEnabled())
}
