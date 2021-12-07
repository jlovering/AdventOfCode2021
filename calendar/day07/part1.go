package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type crabSub struct {
	horzPos int
	number  int
	fuelMap map[int]int64
}

func (cs *crabSub) generateFuelMap(maxPos int) {
	cs.fuelMap = make(map[int]int64)
	for i := 0; i <= maxPos; i++ {
		cs.fuelMap[i] = int64(math.Abs(float64(cs.horzPos - i)))
	}
}

func (cs *crabSub) String() string {
	return fmt.Sprintf("%d (%d): %v", cs.horzPos, cs.number, cs.fuelMap)
}

func sumAllAt(css map[int]*crabSub, at int) int64 {
	var sum int64 = 0
	for _, cs := range css {
		sum += cs.fuelMap[at] * int64(cs.number)
	}
	return sum
}

func populateCrabMap(horzPosNum []int) (map[int]*crabSub, int) {
	var crabSubs map[int]*crabSub = make(map[int]*crabSub, len(horzPosNum))
	maxPos := horzPosNum[len(horzPosNum)-1]
	for i := 0; i < len(horzPosNum); i++ {
		if _, exists := crabSubs[horzPosNum[i]]; !exists {
			ncbs := crabSub{horzPos: horzPosNum[i], number: 1, fuelMap: nil}
			crabSubs[horzPosNum[i]] = &ncbs
			crabSubs[horzPosNum[i]].generateFuelMap(maxPos)
		} else {
			crabSubs[horzPosNum[i]].number++
		}

		util.Dprintf("%v\n", crabSubs[horzPosNum[i]])
	}

	return crabSubs, maxPos
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var horzPos []string
	var horzPosNum []int
	for file_scanner.Scan() {
		line := file_scanner.Text()
		horzPos = strings.Split(line, ",")
		horzPosNum = make([]int, 0, len(horzPos))
		for i := 0; i < len(horzPos); i++ {
			num, err := strconv.Atoi(horzPos[i])
			util.Check_error(err)
			horzPosNum = append(horzPosNum, num)
		}
	}

	sort.Ints(horzPosNum)
	util.Dprintf("%v\n", horzPosNum)

	crabSubs, maxPos := populateCrabMap(horzPosNum)

	var min int64 = math.MaxInt64
	for i := 0; i < maxPos; i++ {
		val := sumAllAt(crabSubs, i)
		util.Dprintf("%d %d\n", i, val)
		if val < min {
			min = val
		}
	}

	return fmt.Sprintf("%d", min)
}
