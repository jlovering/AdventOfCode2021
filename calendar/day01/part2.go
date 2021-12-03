package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	prev_value := -1
	increasing := 0
	decreasing := 0

	var window [3]int
	var populated bool = false
	index := 0

	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value int
		_, err := fmt.Sscanf(line, "%d", &value)
		util.Check_error(err)

		window[index] = value
		index++
		if index >= len(window) {
			populated = true
			index = 0
		}
		if !populated {
			continue
		}

		value = 0
		for _, v := range window {
			value += v
			util.Dprintf("%d ", v)
		}
		util.Dprintf("\n")

		if prev_value == -1 {
			util.Dprintf("%d (N/A)\n", value)
		} else if value > prev_value {
			util.Dprintf("%d (increased)\n", value)
			increasing++
		} else if value < prev_value {
			decreasing++
			util.Dprintf("%d (decreased)\n", value)
		} else if value == prev_value {
			util.Dprintf("%d (no change)\n", value)
		}
		prev_value = value
	}

	return fmt.Sprintf("%d", increasing)
}
