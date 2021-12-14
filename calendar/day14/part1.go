package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func polyStep(substr string, rules map[string]rune) string {
	util.Dprintf("polyStep: %s\n", substr)
	if len(substr) != 2 {
		panic(1)
	}
	var sb strings.Builder
	sb.WriteRune(rules[substr])
	sb.WriteByte(substr[1])
	return sb.String()
}

func polyIter(curPoly string, rules map[string]rune) string {
	var sb strings.Builder

	sb.WriteByte(curPoly[0])
	for i, j := 0, 1; i < j && j < len(curPoly); {
		util.Dprintf("polyIter: %d %d\n", i, j)
		sb.WriteString(polyStep(curPoly[i:j+1], rules))
		i++
		j++
	}
	return sb.String()
}

func runeHist(str string) map[rune]int {
	hist := make(map[rune]int)
	for _, v := range str {
		hist[rune(v)]++
	}
	return hist
}

func sortMapValues(mp map[rune]int) []int {
	out := make([]int, 0)
	for _, v := range mp {
		out = append(out, v)
	}
	sort.Ints(out)
	return out
}

func parseInput(file_scanner *bufio.Scanner) (map[string]rune, string) {
	polyRules := make(map[string]rune)
	var template string
	for file_scanner.Scan() {
		line := file_scanner.Text()
		if strings.Contains(line, "->") {
			sep := strings.Split(line, " -> ")
			polyRules[sep[0]] = rune(sep[1][0])
		} else if line != "" {
			template = line
		}
	}

	return polyRules, template
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	polyRules, template := parseInput(file_scanner)

	util.Dprintf("%v\n%v\n", template, polyRules)

	for i := 0; i < 10; i++ {
		template = polyIter(template, polyRules)
		util.Dprintf("%d\n", len(template))
	}
	counts := runeHist(template)
	util.Dprintf("%v\n", counts)

	sortCounts := sortMapValues(counts)

	return fmt.Sprintf("%d", sortCounts[len(sortCounts)-1] - sortCounts[0])
}
