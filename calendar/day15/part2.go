package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"

	astar "github.com/beefsack/go-astar"
)

func explodeGrid(origGrid map[point]*Tile, mxy point, factor int) (map[point]*Tile, point) {
	var nGrid map[point]*Tile = make(map[point]*Tile, len(origGrid)*factor*factor)
	for j := 0; j < factor; j++ {
		for j_j := 0; j_j <= mxy.y; j_j++ {
			for i := 0; i < factor; i++ {
				for i_i := 0; i_i <= mxy.x; i_i++ {
					ogpt := point{i_i, j_j}
					oT := origGrid[ogpt]
					nP := point{i_i + i*(mxy.x+1), j_j + j*(mxy.y+1)}
					nCost := oT.cost + i + j
					if nCost > 9 {
						nCost = nCost - 9
					}
					nGrid[nP] = &Tile{cost: nCost, refGrid: &nGrid, loc: nP}
					//util.Dprintf("{%d %d}(%d) ", nP.x, nP.y, nCost)
				}
			}
			//util.Dprintf("\n")
		}
	}
	return nGrid, point{(mxy.x+1)*factor - 1, (mxy.y+1)*factor - 1}
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	grid, mxy := parseInput(file_scanner)

	nGrid, nmxy := explodeGrid(grid, mxy, 5)
	printGrid(nGrid, nmxy)

	path, distance, found := astar.Path(nGrid[point{0, 0}], nGrid[nmxy])

	if !found {
		panic(1)
	}

	printPatherSlice(path)
	util.Dprintf("%f\n", distance)

	return fmt.Sprintf("%d", int(distance))
}
