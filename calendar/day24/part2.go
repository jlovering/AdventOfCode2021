package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func Part2(filename string) string {
	/// STDOUT MUST BE FLUSHED MANUALLY!!!
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
				outNum[sV.src] = 1
				outNum[i] = 1 + comp
			} else {
				outNum[i] = 1
				outNum[sV.src] = 1 - comp
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
