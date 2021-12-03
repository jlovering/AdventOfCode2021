package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	util.Setdebug(false)
	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	horz_pos := 0
	depth := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var direction string
		var value int
		_, err = fmt.Sscanf(line, "%s %d", &direction, &value)
		util.Check_error(err)
		switch direction {
		case "forward":
			horz_pos += value
			break
		case "down":
			depth += value
			break
		case "up":
			depth -= value
			break
		}
		util.Dprintf("%s %d %d\n", direction, horz_pos, depth)
	}

	return fmt.Sprintf("%d", horz_pos*depth)
}
