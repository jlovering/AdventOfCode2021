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
		util.Dprintf("%v\n", bitcount)
	}

	gammarate := 0
	epsilonrate := 0
	util.Dprintf("%d %d\n", total_values, (len(bitcount)))
	for i, c := range bitcount {
		if c > total_values/2 {
			gammarate |= 0x1 << (len(bitcount) - 1 - i)
		} else {
			epsilonrate |= 0x1 << (len(bitcount) - 1 - i)
		}
	}
	util.Dprintf("%d %d", gammarate, epsilonrate)

	return fmt.Sprintf("%d", gammarate*epsilonrate)
}
