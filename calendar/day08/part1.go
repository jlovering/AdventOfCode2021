package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processOutput(output string) int {
	values := strings.Fields(output)

	count := 0
	for _, v := range values {
		util.Dprintf("%v ", v)
		switch len(v) {
		case 2: //1
			fallthrough
		case 3: //7
			fallthrough
		case 4: //4
			fallthrough
		case 7: //8
			count += 1
			util.Dprintf("(%d) %d\n", len(v), count)
		default:
			util.Dprintf("(%d) %d\n", len(v), count)
		}
	}
	util.Dprintf("%d\n", count)
	return count
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var output string
	sum := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		split1 := strings.Split(line, "|")
		output = split1[1]
		sum += processOutput(output)
	}

	return fmt.Sprintf("%d", sum)
}
