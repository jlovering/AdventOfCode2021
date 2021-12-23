package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"

	"github.com/beefsack/go-astar"
)

func parseInput2(file_scanner *bufio.Scanner) burrow {

	barrow := burrow{}

	file_scanner.Scan()
	_ = file_scanner.Text()
	file_scanner.Scan()
	_ = file_scanner.Text()
	file_scanner.Scan()
	s1 := file_scanner.Text()
	file_scanner.Scan()
	s2 := file_scanner.Text()
	file_scanner.Scan()
	s3 := file_scanner.Text()
	file_scanner.Scan()
	s4 := file_scanner.Text()

	barrow.roomA.size = 4
	barrow.roomB.size = 4
	barrow.roomC.size = 4
	barrow.roomD.size = 4

	barrow.goToYourRoom("roomA", rune(s4[3]))
	barrow.goToYourRoom("roomB", rune(s4[5]))
	barrow.goToYourRoom("roomC", rune(s4[7]))
	barrow.goToYourRoom("roomD", rune(s4[9]))

	barrow.goToYourRoom("roomA", rune(s3[3]))
	barrow.goToYourRoom("roomB", rune(s3[5]))
	barrow.goToYourRoom("roomC", rune(s3[7]))
	barrow.goToYourRoom("roomD", rune(s3[9]))

	barrow.goToYourRoom("roomA", rune(s2[3]))
	barrow.goToYourRoom("roomB", rune(s2[5]))
	barrow.goToYourRoom("roomC", rune(s2[7]))
	barrow.goToYourRoom("roomD", rune(s2[9]))

	barrow.goToYourRoom("roomA", rune(s1[3]))
	barrow.goToYourRoom("roomB", rune(s1[5]))
	barrow.goToYourRoom("roomC", rune(s1[7]))
	barrow.goToYourRoom("roomD", rune(s1[9]))

	barrow.roomA.lock = true
	barrow.roomB.lock = true
	barrow.roomC.lock = true
	barrow.roomD.lock = true
	return barrow
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	b := parseInput2(file_scanner)
	util.Dprintf("%v\n", b)

	start := graphNode{
		state: b,
	}

	end := graphNode{
		state: burrow{
			hallway: [11]rune{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			roomA: burrowRoom{
				room:  [4]rune{'A', 'A', 'A', 'A'},
				lock:  false,
				count: 4,
				size:  4,
			},
			roomB: burrowRoom{
				room:  [4]rune{'B', 'B', 'B', 'B'},
				lock:  false,
				count: 4,
				size:  4,
			},
			roomC: burrowRoom{
				room:  [4]rune{'C', 'C', 'C', 'C'},
				lock:  false,
				count: 4,
				size:  4,
			},
			roomD: burrowRoom{
				room:  [4]rune{'D', 'D', 'D', 'D'},
				lock:  false,
				count: 4,
				size:  4,
			},
		},
	}

	util.Dprintf("%v\n", end.state)
	nastyNodeMap[end.state] = &end

	path, distance, found := astar.Path(&start, &end)

	if !found {
		panic("No path")
	}

	for i := range path {
		gn := path[len(path)-1-i].(*graphNode)
		fmt.Printf("%v\n", gn.state)
	}

	return fmt.Sprintf("%d", uint(distance))
}
