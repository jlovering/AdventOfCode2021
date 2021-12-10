package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
)

func bracketSyntaxCheckAndFix(line string) []interface{} {
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
				return nil
			} else {
				parseStack.Remove(front)
			}
		}
		printList(parseStack)
	}

	ret := make([]interface{}, parseStack.Len())
	for v := parseStack.Front(); v != nil; v = v.Next() {
		r := v.Value
		ret = append(ret, r)
	}
	return ret
}

func scoreLine(missing []interface{}) int {
	sum := 0
	for _, r := range missing {
		pt := 0
		switch r {
		case ')':
			pt = 1
		case ']':
			pt = 2
		case '}':
			pt = 3
		case '>':
			pt = 4
		}
		sum = sum*5 + pt
	}
	return sum
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	scores := make([]int, 0, 1000)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		mVs := bracketSyntaxCheckAndFix(line)
		var MissingValues []interface{} = make([]interface{}, 0, len(mVs))
		for _, mV := range mVs {
			if mV != nil {
				MissingValues = append(MissingValues, mV)
			}
		}
		if len(MissingValues) > 0 {
			ls := scoreLine(MissingValues)
			scores = append(scores, ls)
			util.Dprintf("%d\n", ls)
		}
	}
	sort.Ints(scores)
	util.Dprintf("%v\n", scores)

	return fmt.Sprintf("%d", scores[len(scores)/2])
}
