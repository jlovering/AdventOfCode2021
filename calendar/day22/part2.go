package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func (c cube) String() string {
	return fmt.Sprintf("x=%d..%d,y=%d..%d,z=%d..%d", c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax)
}

func (cf cubeField) String() string {
	var sb strings.Builder
	for c, v := range cf {
		if v {
			for x := c.xmin; x < 15 && x <= c.xmax; x++ {
				for y := c.ymin; y < 15 && y <= c.ymax; y++ {
					for z := c.zmin; z < 15 && z <= c.zmax; z++ {
						sb.WriteString(fmt.Sprintf("%v\n", point{x, y, z}))
					}
				}
			}
		}
	}
	return sb.String()
}

func (c cube) cubeSize() uint {
	x := math.Abs(float64(c.xmax) + 1 - float64(c.xmin))
	y := math.Abs(float64(c.ymax) + 1 - float64(c.ymin))
	z := math.Abs(float64(c.zmax) + 1 - float64(c.zmin))

	return uint(x * y * z)
}

func cubeAxisOverlap(c1min, c1max, c2min, c2max int) (int, int, bool) {
	//For any axis there are 4 ways to overlap:
	// |------|
	//|--------|
	if c2min <= c1min && c2max >= c1max {
		return c1min, c1max, true
	}

	// |------|
	//   |--|
	if c1min <= c2min && c1max >= c2max {
		return c2min, c2max, true
	}

	// |------|
	//    |------|
	if c1min <= c2min && c2min <= c1max && c1max <= c2max {
		return c2min, c1max, true
	}

	//     |------|
	// |------|
	if c2min <= c1min && c1min <= c2max && c2max <= c1max {
		return c1min, c2max, true
	}

	return 0, 0, false
}

func (c1 cube) overlap(c2 cube) (cube, bool) {
	overlap := cube{}

	valid := false
	overlap.xmin, overlap.xmax, valid = cubeAxisOverlap(c1.xmin, c1.xmax, c2.xmin, c2.xmax)
	if !valid {
		return cube{}, false
	}
	overlap.ymin, overlap.ymax, valid = cubeAxisOverlap(c1.ymin, c1.ymax, c2.ymin, c2.ymax)
	if !valid {
		return cube{}, false
	}
	overlap.zmin, overlap.zmax, valid = cubeAxisOverlap(c1.zmin, c1.zmax, c2.zmin, c2.zmax)
	if !valid {
		return cube{}, false
	}

	return overlap, true
}

func (c cube) cubeSplitX(split int) []cube {
	cubes := []cube{}

	util.IncreasedebugIndent()
	util.Dprintf("Split X: %d (%d) %d\n", c.xmin, split, c.xmax)
	if split < c.xmin || split >= c.xmax {
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
	} else if split == c.xmin {
		cubes = append(cubes, cube{c.xmin, c.xmin, c.ymin, c.ymax, c.zmin, c.zmax})
		cubes = append(cubes, cube{split + 1, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
	} else {
		cubes = append(cubes, cube{c.xmin, split - 1, c.ymin, c.ymax, c.zmin, c.zmax})
		cubes = append(cubes, cube{split, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
	}

	util.DecreasedebugIndent()
	return cubes
}

func (c cube) cubeDblSplitX(split1, split2 int) []cube {
	cubes := []cube{}

	util.IncreasedebugIndent()
	util.Dprintf("Split X: %d (%d)(%d) %d\n", c.xmin, split1, split2, c.xmax)
	if split1 <= c.xmin {
		if split2 >= c.xmax {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
		} else {
			cubes = append(cubes, cube{c.xmin, split2, c.ymin, c.ymax, c.zmin, c.zmax})
			cubes = append(cubes, cube{split2 + 1, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
		}
	} else /*split1 > c.xmin*/ {
		if split2 >= c.xmax {
			cubes = append(cubes, cube{c.xmin, split1 - 1, c.ymin, c.ymax, c.zmin, c.zmax})
			cubes = append(cubes, cube{split1, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
		} else {
			cubes = append(cubes, cube{c.xmin, split1 - 1, c.ymin, c.ymax, c.zmin, c.zmax})
			cubes = append(cubes, cube{split1, split2, c.ymin, c.ymax, c.zmin, c.zmax})
			cubes = append(cubes, cube{split2 + 1, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
		}
	}

	util.DecreasedebugIndent()
	return cubes
}

func (c cube) cubeSplitY(split int) []cube {
	cubes := []cube{}

	util.IncreasedebugIndent()
	util.Dprintf("Split Y: %d (%d) %d\n", c.ymin, split, c.ymax)
	if split < c.ymin || split >= c.ymax {
		//No split
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
	} else if split == c.ymin {
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymin, c.zmin, c.zmax})
		cubes = append(cubes, cube{c.xmin, c.xmax, split + 1, c.ymax, c.zmin, c.zmax})
	} else {
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, split - 1, c.zmin, c.zmax})
		cubes = append(cubes, cube{c.xmin, c.xmax, split, c.ymax, c.zmin, c.zmax})
	}

	util.DecreasedebugIndent()
	return cubes
}

func (c cube) cubeDblSplitY(split1, split2 int) []cube {
	cubes := []cube{}

	util.IncreasedebugIndent()
	util.Dprintf("Split Y: %d (%d)(%d) %d\n", c.ymin, split1, split2, c.ymax)
	if split1 <= c.ymin {
		if split2 >= c.ymax {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
		} else {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, split2, c.zmin, c.zmax})
			cubes = append(cubes, cube{c.xmin, c.xmax, split2 + 1, c.ymax, c.zmin, c.zmax})
		}
	} else /*split1 > c.xmin*/ {
		if split2 >= c.ymax {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, split1 - 1, c.zmin, c.zmax})
			cubes = append(cubes, cube{c.xmin, c.xmax, split1, c.ymax, c.zmin, c.zmax})
		} else {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, split1 - 1, c.zmin, c.zmax})
			cubes = append(cubes, cube{c.xmin, c.xmax, split1, split2, c.zmin, c.zmax})
			cubes = append(cubes, cube{c.xmin, c.xmax, split2 + 1, c.ymax, c.zmin, c.zmax})
		}
	}

	util.DecreasedebugIndent()
	return cubes
}

func (c cube) cubeSplitZ(split int) []cube {
	cubes := []cube{}

	util.IncreasedebugIndent()
	util.Dprintf("Split Z: %d (%d) %d\n", c.zmin, split, c.zmax)
	if split < c.zmin || split >= c.zmax {
		//No split
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
	} else if split == c.zmin {
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmin})
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, split + 1, c.zmax})
	} else {
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, split - 1})
		cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, split, c.zmax})
	}

	util.DecreasedebugIndent()
	return cubes
}

func (c cube) cubeDblSplitZ(split1, split2 int) []cube {
	cubes := []cube{}

	util.IncreasedebugIndent()
	util.Dprintf("Split Z: %d (%d)(%d) %d\n", c.zmin, split1, split2, c.zmax)
	if split1 <= c.zmin {
		if split2 >= c.zmax {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, c.zmax})
		} else {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, split2})
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, split2 + 1, c.zmax})
		}
	} else /*split1 > c.xmin*/ {
		if split2 >= c.zmax {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, split1 - 1})
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, split1, c.zmax})
		} else {
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, c.zmin, split1 - 1})
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, split1, split2})
			cubes = append(cubes, cube{c.xmin, c.xmax, c.ymin, c.ymax, split2 + 1, c.zmax})
		}
	}

	util.DecreasedebugIndent()
	return cubes
}

func (c1 cube) removeCube(c2 cube) []cube {
	util.IncreasedebugIndent()
	// Slice the cube on the x axis, then slice all resulting cubes on the y axis, and then all resulting cubes again on the z axis
	// This is not optimal, but should be OK

	//Split will return no cubes if the cut line is out of bound, the original cube if the slice line is on a boundary, or two cubes if a split happens
	//Because of this, the second split has to happen on the last item of the previous (as it's the split's tail)
	// cubesX = append(cubesX, c1.cubeSplitX(c2.xmin)...)
	// cn := cubesX[len(cubesX)-1]
	// cubesX = cubesX[:len(cubesX)-1]
	// cubesX = append(cubesX, cn.cubeSplitX(c2.xmax)...)
	cubesX := c1.cubeDblSplitX(c2.xmin, c2.xmax)

	util.Dprintf("%v\n", cubesX)

	cubesY := []cube{}
	for _, c := range cubesX {
		if _, exist := c.overlap(c2); exist {
			cubesY = append(cubesY, c.cubeDblSplitY(c2.ymin, c2.ymax)...)
			// cn := cubesY[len(cubesY)-1]
			// cubesY = cubesY[:len(cubesY)-1]
			// cubesY = append(cubesY, cn.cubeSplitY(c2.ymax)...)
		} else {
			cubesY = append(cubesY, c)
		}
	}

	util.Dprintf("%v\n", cubesY)

	cubesZ := []cube{}
	for _, c := range cubesY {
		if _, exist := c.overlap(c2); exist {
			cubesZ = append(cubesZ, c.cubeDblSplitZ(c2.zmin, c2.zmax)...)
			// cn := cubesZ[len(cubesZ)-1]
			// cubesZ = cubesZ[:len(cubesZ)-1]
			// cubesZ = append(cubesZ, cn.cubeSplitZ(c2.zmax)...)
		} else {
			cubesZ = append(cubesZ, c)
		}
	}

	util.Dprintf("%v\n", cubesZ)

	outCubes := []cube{}
	for _, c := range cubesZ {
		if c == c2 {
			util.Dprintf("Ignoring overlap: %v\n", c)
			continue
		}
		outCubes = append(outCubes, c)
	}

	util.Dprintf("%v\n", outCubes)

	util.DecreasedebugIndent()
	return outCubes
}

func (cf *cubeField) switchCubes(ci cube, state bool) {
	if len(*cf) == 0 && state {
		util.Dprintf("Inserting first: %v\n", ci)
		(*cf)[ci] = state
		return
	}

	util.Dprintf("Inserting %v\n", ci)
	nCf := cubeField{}
	for c, v := range *cf {
		util.IncreasedebugIndent()
		util.Dprintf("Considering: %v\n", c)
		over, exists := c.overlap(ci)
		if exists {
			util.IncreasedebugIndent()
			util.Dprintf("Overlap: %v (%v)\n", over, c)
			nCs := c.removeCube(over)
			for _, nc := range nCs {
				util.IncreasedebugIndent()
				util.Dprintf("Inserting subCube: %v\n", nc)
				nCf[nc] = v
				util.DecreasedebugIndent()
			}
			util.Dprintf("Final Inserted subCubes:\n")
			for c := range nCf {
				util.IncreasedebugIndent()
				util.Dprintf("%v\n", c)
				util.DecreasedebugIndent()
			}
			util.DecreasedebugIndent()
		} else {
			nCf[c] = v
		}
		util.DecreasedebugIndent()
	}
	util.Dprintf("Inserting new: %v\n", ci)
	nCf[ci] = state

	*cf = nCf
}

func (cf cubeField) countEnabled() uint {
	count := uint(0)
	//util.Dprintf("Counting...\n")
	util.IncreasedebugIndent()
	for c, v := range cf {
		if v {
			//util.Dprintf("%v : %d\n", c, c.cubeSize())
			count += c.cubeSize()
		}
	}
	util.DecreasedebugIndent()
	return count
}

func parseInput2(file_scanner *bufio.Scanner) cubeField {
	cf := cubeField{}
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
		c := cube{xstart, xend, ystart, yend, zstart, zend}
		if op == "on" {
			cf.switchCubes(c, true)
			//util.Dprintf("All Points:\n%s\n", cf)
			util.Dprintf("New Count: %d\n\n", cf.countEnabled())
		} else if op == "off" {
			cf.switchCubes(c, false)
			//util.Dprintf("All Points:\n%s\n", cf)
			util.Dprintf("New Count: %d\n\n", cf.countEnabled())
		}
	}
	return cf
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	cf := parseInput2(file_scanner)

	return fmt.Sprintf("%d", cf.countEnabled())
}
