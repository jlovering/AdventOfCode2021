package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func difference(a, b string) string {
	mb := make(map[rune]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = diff + string(x)
		}
	}
	return diff
}

func findNine(candidates []string, four string) (string, int) {
	util.Dprintf("%v %v\n", candidates, four)

	for i, v := range candidates {
		diff := difference(v, four)
		util.Dprintf("%v %v\n", v, diff)
		if len(diff) == 2 {
			return v, i
		}
	}
	panic(1)
}

func findZeroAndSix(candidates []string, one string) (string, string) {
	util.Dprintf("%v %v\n", candidates, one)

	if len(candidates) != 2 {
		panic(1)
	}
	diff := difference(candidates[0], one)
	util.Dprintf("%v\n", diff)
	if len(diff) == 4 {
		return candidates[0], candidates[1]
	} else if len(diff) == 5 {
		return candidates[1], candidates[0]
	}
	panic(1)
}

func findFive(candidates []string, six string) (string, int) {

	for i, v := range candidates {
		diff := difference(six, v)
		util.Dprintf("%v %v\n", v, diff)
		if len(diff) == 1 {
			return v, i
		}
	}
	panic(1)
}

func findTwoAndThree(candidates []string, nine string) (string, string) {
	util.Dprintf("%v %v\n", candidates, nine)

	if len(candidates) != 2 {
		panic(1)
	}
	diff := difference(nine, candidates[0])
	util.Dprintf("%v\n", diff)
	if len(diff) == 2 {
		return candidates[0], candidates[1]
	} else if len(diff) == 1 {
		return candidates[1], candidates[0]
	}
	panic(1)
}

func bruteforcesolveysolve(translationMap map[string]int, sortedInput []string) {
	uniqueMap := make(map[int]int)

	for k, _ := range translationMap {
		for _, s := range sortedInput {
			if len(k) < len(s) {
				uniqueMap[len(difference(s, k))]++
				util.Dprintf("%v,%v = %v\n", s, k, len(difference(s, k)))
			} else if len(s) < len(k) {
				uniqueMap[len(difference(s, k))]++
				util.Dprintf("%v,%v = %v\n", k, s, len(difference(k, s)))
			}
		}
	}
	util.Dprintf("%v\n", uniqueMap)
}

func processInput(input []string) map[string]int {
	var translationMap map[string]int = make(map[string]int, 10)

	sort.Slice(input, func(i, j int) bool {
		return len(input[i]) < len(input[j])
	})

	var sortedInput []string = make([]string, 0, len(input))
	for i := 0; i < len(input); i++ {
		sortedInput = append(sortedInput, SortString(input[i]))
	}

	translationMap[sortedInput[0]] = 1
	translationMap[sortedInput[1]] = 7
	translationMap[sortedInput[2]] = 4
	translationMap[sortedInput[9]] = 8

	candids := sortedInput[6:9]
	nine, pos := findNine(candids, sortedInput[2])
	translationMap[nine] = 9
	if pos == 0 {
		candids = candids[1:]
	} else if pos == 1 {
		candids = append(candids[:1], candids[2])
	} else {
		candids = candids[:pos]
	}
	zero, six := findZeroAndSix(candids, sortedInput[0])
	translationMap[six] = 6
	translationMap[zero] = 0

	candids = sortedInput[3:6]
	five, pos := findFive(candids, six)
	translationMap[five] = 5
	if pos == 0 {
		candids = candids[1:]
	} else if pos == 1 {
		candids = append(candids[:1], candids[2])
	} else {
		candids = candids[:pos]
	}
	two, three := findTwoAndThree(candids, nine)
	translationMap[two] = 2
	translationMap[three] = 3

	//bruteforcesolveysolve(translationMap, sortedInput)
	return translationMap
}

func doTranslation(translationMap map[string]int, segments []string) int {
	outValue := 0
	for _, v := range segments {
		sortSegs := SortString(v)
		outValue *= 10
		outValue += translationMap[sortSegs]
	}
	return outValue
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var input []string
	var output []string
	sum := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		split1 := strings.Split(line, "|")
		input = strings.Fields(split1[0])
		output = strings.Fields(split1[1])
		translationMap := processInput(input)
		sum += doTranslation(translationMap, output)
		util.Dprintf("%v\n", translationMap)
	}

	return fmt.Sprintf("%d", sum)
}
