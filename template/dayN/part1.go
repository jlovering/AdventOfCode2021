package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func parseInput(file_scanner *bufio.Scanner) string {
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value int
		fmt.Sscanf(line, "%d", &value)
	}
	return ""
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	_ = parseInput(file_scanner)

	return ""
}
