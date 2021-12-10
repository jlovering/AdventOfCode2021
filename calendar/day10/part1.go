package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func printList(l *list.List) {
	if l == nil {
		return
	}
	for v := l.Front(); v != nil; v = v.Next() {
		util.Dprintf("%c ", v.Value)
	}
	util.Dprintf("\n")
}

func bracketSyntaxCheck(line string) rune {
	parseStack := list.New()
	runeLine := []rune(line)
	util.Dprintf(line)
	util.Dprintf("\n")
	for _, r := range runeLine {
		switch r {
		case '(':
			parseStack.PushFront(')')
		case '[':
			parseStack.PushFront(']')
		case '{':
			parseStack.PushFront('}')
		case '<':
			parseStack.PushFront('>')
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			front := parseStack.Front()
			if r != front.Value {
				util.Dprintf("Expected %c, but found %c instead.", r, front.Value)
				return r
			} else {
				parseStack.Remove(front)
			}
		}
		printList(parseStack)
	}
	return '0'
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var badValues []rune = make([]rune, 0, 1000)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		bV := bracketSyntaxCheck(line)
		if bV != '0' {
			badValues = append(badValues, bV)
		}
	}

	sum := 0
	for _, r := range badValues {
		switch r {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	return fmt.Sprintf("%d", sum)
}
