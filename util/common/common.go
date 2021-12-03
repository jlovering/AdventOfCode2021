package util

import (
	"bufio"
	"fmt"
	"os"
)

var debug_output bool = true
var sdio_reader *bufio.Reader = bufio.NewReader(os.Stdin)
var sdout_writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Dprintf(f string, a ...interface{}) {
	if debug_output {
		fmt.Fprintf(sdout_writer, f, a...)
	}
}

func Printf(f string, a ...interface{}) { fmt.Fprintf(sdout_writer, f, a...) }

func Scanf(f string, a ...interface{}) { fmt.Fscanf(sdio_reader, f, a...) }

func Check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func Setdebug(db bool) {
	debug_output = db
}

func SdoutFlush() {
	sdout_writer.Flush()
}
