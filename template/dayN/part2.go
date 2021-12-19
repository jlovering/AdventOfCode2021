package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"os"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	_ = parseInput(file_scanner)

	return ""
}
