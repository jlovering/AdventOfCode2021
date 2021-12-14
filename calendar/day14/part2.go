package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func mergeHist(h1, h2 map[rune]uint64) map[rune]uint64 {
	hist := make(map[rune]uint64)
	for k, v := range h1 {
		hist[k] = v
	}
	for k, v := range h2 {
		hist[k] += v
	}
	return hist
}

func maxMinBig(mp map[rune]uint64) (uint64, uint64) {
	var max, min uint64 = 0, math.MaxUint64
	for _, v := range mp {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func tabBuild(depth int) string {
	var sb strings.Builder
	for i := 0; i < depth; i++ {
		sb.WriteRune(' ')
	}
	return sb.String()
}

type pair struct {
	l rune
	r rune
}

func expandRules(orgRules map[string]rune) map[pair][]pair {
	newRules := make(map[pair][]pair, 0)

	for k, v := range orgRules {
		newRules[pair{l: rune(k[0]), r: rune(k[1])}] = []pair{{l: rune(k[0]), r: v}, {l: v, r: rune(k[1])}}
	}
	return newRules
}

func genPairs(input string) map[pair]uint64 {
	pairs := make(map[pair]uint64)
	for i, j := 0, 1; i < j && j < len(input); {
		pairTmp := pair{l: rune(input[i]), r: rune(input[j])}
		pairs[pairTmp]++
		i++
		j++
	}
	return pairs
}

func explodePairs(pairCount map[pair]uint64) map[rune]uint64 {
	hist := make(map[rune]uint64)
	for k, v := range pairCount {
		hist[k.l] += v
	}
	return hist
}
func doIteration(pairCount map[pair]uint64, newRules map[pair][]pair) map[pair]uint64 {
	npairCount := make(map[pair]uint64)
	for k := range pairCount {
		ruleOut := newRules[k]
		for _, p := range ruleOut {
			npairCount[p] += pairCount[k]
		}
	}
	return npairCount
}

func printExpandRules(pxr map[pair][]pair) {
	for k, v := range pxr {
		util.Dprintf("%c%c : ", k.l, k.r)
		for _, v2 := range v {
			util.Dprintf("%c%c ", v2.l, v2.r)
		}
		util.Dprintf("\n")
	}
	util.Dprintf("\n")
}

func printPairMapThing(pm map[pair]uint64) {
	for k, v := range pm {
		util.Dprintf("%c%c:%d ", k.l, k.r, v)
	}
	util.Dprintf("\n")
}

func printHist(pH map[rune]uint64) {
	for k, v := range pH {
		util.Dprintf("%c:%d ", k, v)
	}
	util.Dprintf("\n")
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	polyRules, template := parseInput(file_scanner)

	nRules := expandRules(polyRules)
	printExpandRules(nRules)
	pairCount := genPairs(template)
	printPairMapThing(pairCount)

	for i := 0; i < 40; i++ {
		pairCount = doIteration(pairCount, nRules)
		util.Dprintf("%d:", i)
		printPairMapThing(pairCount)
		util.Dprintf("\n")
	}

	hist := explodePairs(pairCount)
	hist[rune(template[len(template)-1])]++
	printHist(hist)

	max, min := maxMinBig(hist)

	return fmt.Sprintf("%d", max-min)
}
