package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingocard struct {
	card          [5][5]string
	calledNumbers map[string]bool
}

func (b *bingocard) callNumber(number string) {
	b.calledNumbers[number] = true
}

func (b *bingocard) isCalled(i int, j int) bool {
	return b.calledNumbers[b.card[i][j]]
}

func (b *bingocard) checkWin() bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.isCalled(i, j) {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if !b.isCalled(i, j) {
				break
			}
			if i == 4 {
				return true
			}
		}
	}

	return false
}

func (b *bingocard) callAndCheck(number string) bool {
	b.callNumber(number)
	return b.checkWin()
}

func (b *bingocard) scoreCard() int {
	total := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.isCalled(i, j) {
				numNum, err := strconv.Atoi(b.card[i][j])
				util.Check_error(err)
				total += int(numNum)
			}
		}
	}

	return total
}

func playGame(calledNumbers []string, bingocards []bingocard) int {
	for _, num := range calledNumbers {
		for _, bc := range bingocards {
			if bc.callAndCheck(num) {
				numNum, err := strconv.Atoi(num)
				util.Check_error(err)
				cardScore := bc.scoreCard()
				return cardScore * numNum
			}
		}
	}
	return 0
}

func Part1(filename string) string {
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

	return fmt.Sprintf("%d", playGame(calledNumbers, bingocards))
}
