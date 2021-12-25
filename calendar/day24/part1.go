package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inputBank struct {
	div    int
	check  int
	offset int

	count int
}

func (iB *inputBank) processInst(vals string) bool {
	iB.count++

	//util.Dprintf("Adding: %s %d\n", vals, iB.count)
	switch iB.count {
	case 5:
		val, err := strconv.Atoi(vals)
		util.Check_error(err)
		iB.div = val
	case 6:
		val, err := strconv.Atoi(vals)
		util.Check_error(err)
		iB.check = val
	case 16:
		val, err := strconv.Atoi(vals)
		util.Check_error(err)
		iB.offset = val
	case 18:
		return true
	}
	return false
}

func parseInput(file_scanner *bufio.Scanner) []inputBank {

	iBs := []inputBank{}
	curIB := inputBank{}
	for file_scanner.Scan() {
		line := file_scanner.Text()
		s := strings.SplitN(line, " ", 3)
		if len(s) == 2 {
			curIB.processInst("")
			continue
		}
		_, _, vals := s[0], s[1], s[2]
		if curIB.processInst(vals) {
			iBs = append(iBs, curIB)
			curIB = inputBank{}
		}
	}

	return iBs
}

type stackVal struct {
	str string
	src int
	val int
}
type stack []stackVal

func (s stack) push(v stackVal) stack {
	ns := append(s, v)
	return ns
}

func (s stack) pop() (stackVal, stack) {
	o := s[len(s)-1]
	ns := s[:len(s)-1]
	return o, ns
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	iBs := parseInput(file_scanner)

	zStack := stack{}
	outNum := [14]int{}
	for i, v := range iBs {
		if v.check > 9 {
			//util.Dprintf("PUSH input[%d] + %d\n", i, v.offset)
			str := fmt.Sprintf("input[%d]", i)
			zStack = zStack.push(stackVal{str, i, v.offset})
		} else if v.check <= 0 {
			//util.Dprintf("POP input[%d] == popped + %d\n", i, v.check)
			var sV stackVal
			sV, zStack = zStack.pop()
			comp := sV.val + v.check
			if comp >= 0 {
				outNum[sV.src] = 9 - comp
				outNum[i] = 9
			} else {
				outNum[i] = 9 + comp
				outNum[sV.src] = 9
			}
			fmt.Printf("input[%d] = %s + %d \n", i, sV.str, comp)
		} else {
			util.Dprintf("%v\n", v)
			panic("")
		}
	}
	if len(zStack) != 0 {
		panic("Unbalanced")
	}

	for _, d := range outNum {
		fmt.Printf("%d", d)
	}
	fmt.Printf("\n")

	return " "
}
