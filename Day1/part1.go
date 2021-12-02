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

	prev_value := -1
	increasing := 0
	decreasing := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value int
		_, err := fmt.Sscanf(line, "%d", &value)
		check_error(err)

		if prev_value == -1 {
			printf("%s (N/A)\n", line)
		} else if value > prev_value {
			printf("%s (increased)\n", line)
			increasing++
		} else if value < prev_value {
			decreasing++
			printf("%s (decreased)\n", line)
		}
		prev_value = value
	}

	return fmt.Sprintf("%d", increasing)
}
