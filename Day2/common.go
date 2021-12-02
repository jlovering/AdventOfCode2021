package AdventOfCode

import (
	"bufio"
	"fmt"
	"os"
)

var debug_output bool = true
var sdio_reader *bufio.Reader = bufio.NewReader(os.Stdin)
var sdout_writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func dprintf(f string, a ...interface{}) {
	if debug_output {
		fmt.Fprintf(sdout_writer, f, a...)
	}
}

func printf(f string, a ...interface{}) { fmt.Fprintf(sdout_writer, f, a...) }

func scanf(f string, a ...interface{}) { fmt.Fscanf(sdio_reader, f, a...) }

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}
