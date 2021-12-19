package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var master_debug bool = true
var debug_output bool = true
var debug_indent int = 0
var sdio_reader *bufio.Reader = bufio.NewReader(os.Stdin)
var sdout_writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Dprintf(f string, a ...interface{}) {
	if master_debug && debug_output {
		fmt.Fprintf(sdout_writer, strings.Repeat(" ", debug_indent)+f, a...)
	}
}

func Printf(f string, a ...interface{}) { fmt.Fprintf(sdout_writer, f, a...) }

func Scanf(f string, a ...interface{}) { fmt.Fscanf(sdio_reader, f, a...) }

func Check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func Setmasterdebug(db bool) {
	master_debug = db
}

func Setdebug(db bool) {
	debug_output = db
}

func SetdebugIndent(indent int) {
	debug_indent = indent
}

func IncreasedebugIndent() {
	debug_indent++
}

func DecreasedebugIndent() {
	debug_indent--
}

func SdoutFlush() {
	sdout_writer.Flush()
}
