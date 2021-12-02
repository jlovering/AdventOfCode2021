package AdventOfCode

import (
	"bufio"
	"fmt"
	"os"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer sdout_writer.Flush()

	f, err := os.Open(filename)
	check_error(err)
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
		check_error(err)

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
			dprintf("%d ", v)
		}
		dprintf("\n")

		if prev_value == -1 {
			dprintf("%d (N/A)\n", value)
		} else if value > prev_value {
			dprintf("%d (increased)\n", value)
			increasing++
		} else if value < prev_value {
			decreasing++
			dprintf("%d (decreased)\n", value)
		} else if value == prev_value {
			dprintf("%d (no change)\n", value)
		}
		prev_value = value
	}

	return fmt.Sprintf("%d", increasing)
}
