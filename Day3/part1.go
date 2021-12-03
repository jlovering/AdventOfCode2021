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

	bitcount := make(map[int]int)
	total_values := 0

	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value string
		fmt.Sscanf(line, "%s", &value)
		for i, c := range value {
			if c == '1' {
				bitcount[i]++
			}
		}
		total_values++
		dprintf("%v\n", bitcount)
	}

	gammarate := 0
	epsilonrate := 0
	dprintf("%d %d\n", total_values, (len(bitcount)))
	for i, c := range bitcount {
		if c > total_values/2 {
			gammarate |= 0x1 << (len(bitcount) - 1 - i)
		} else {
			epsilonrate |= 0x1 << (len(bitcount) - 1 - i)
		}
	}
	dprintf("%d %d", gammarate, epsilonrate)

	return fmt.Sprintf("%d", gammarate*epsilonrate)
}
