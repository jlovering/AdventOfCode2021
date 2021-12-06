package adventofcode

import (
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

	var initialAges []string
	for file_scanner.Scan() {
		line := file_scanner.Text()
		initialAges = strings.Split(line, ",")
	}

	var lampFishes lampfishSchool = make(lampfishSchool, 0, len(initialAges))
	for _, age := range initialAges {
		ageNum, err := strconv.Atoi(age)
		util.Check_error(err)
		lfg := lampfishGroup{size: 1, count: ageNum}
		lampFishes = append(lampFishes, lfg)
	}

	lampFishes = mergeAll(lampFishes)
	lampFishes.print()

	count := runSimulation(lampFishes, 256)
	return fmt.Sprintf("%v", count)
}
