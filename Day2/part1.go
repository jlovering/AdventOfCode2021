package AdventOfCode

import (
	"bufio"
	"fmt"
	"os"
)

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer sdout_writer.Flush()

	f, err := os.Open(filename)
	check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value int
		fmt.Sscanf(line, "%d", value)
	}

	return ""
}
