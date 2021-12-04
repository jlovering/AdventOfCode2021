package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func playGameToLast(calledNumbers []string, bingocards []bingocard) int {
	for _, num := range calledNumbers {
		new_bingocards := make([]bingocard, 0, len(bingocards))
		for _, bc := range bingocards {
			if !bc.callAndCheck(num) {
				new_bingocards = append(new_bingocards, bc)
			} else {
				if len(bingocards) == 1 {
					numNum, err := strconv.Atoi(num)
					util.Check_error(err)
					cardScore := bingocards[0].scoreCard()
					return cardScore * numNum
				}
			}
		}
		util.Dprintf("%s %v\n", num, new_bingocards)
		bingocards = new_bingocards
	}
	return 0
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var calledNumbers []string

	file_scanner.Scan()
	line := file_scanner.Text()
	calledNumbers = strings.Split(line, ",")
	util.Dprintf("%v\n", calledNumbers)

	bingocards := make([]bingocard, 0, 100)
	cardcount := 0
	for file_scanner.Scan() {
		_ = file_scanner.Text()
		thisCard := bingocard{calledNumbers: make(map[string]bool)}
		for j := 0; j < 5; j++ {
			file_scanner.Scan()
			line = file_scanner.Text()
			cardnumline := strings.Fields(line)
			for i, n := range cardnumline {
				thisCard.card[j][i] = n
			}
		}
		bingocards = append(bingocards, thisCard)
		util.Dprintf("%v\n\n", bingocards[cardcount].card)
		cardcount++
	}

	return fmt.Sprintf("%d", playGameToLast(calledNumbers, bingocards))
}
