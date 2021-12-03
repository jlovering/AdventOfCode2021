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

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	prev_value := -1
	increasing := 0
	decreasing := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value int
		_, err := fmt.Sscanf(line, "%d", &value)
		util.Check_error(err)

		if prev_value == -1 {
			util.Dprintf("%s (N/A)\n", line)
		} else if value > prev_value {
			util.Dprintf("%s (increased)\n", line)
			increasing++
		} else if value < prev_value {
			decreasing++
			util.Dprintf("%s (decreased)\n", line)
		}
		prev_value = value
	}

	return fmt.Sprintf("%d", increasing)
}
