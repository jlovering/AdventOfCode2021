package AdventOfCode

import (
	"bufio"
	"fmt"
	"os"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer sdout_writer.Flush()
	debug_output = true

	f, err := os.Open(filename)
	check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	horz_pos := 0
	depth := 0
	aim := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var direction string
		var value int
		_, err = fmt.Sscanf(line, "%s %d", &direction, &value)
		check_error(err)
		switch direction {
		case "forward":
			horz_pos += value
			depth += value * aim
			break
		case "down":
			aim += value
			break
		case "up":
			aim -= value
			break
		}
		dprintf("%s %d %d\n", direction, horz_pos, depth)
	}

	return fmt.Sprintf("%d", horz_pos*depth)
}
