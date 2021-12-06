package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type lampfishGroup struct {
	count int
	size  int64
}

type lampfishSchool []lampfishGroup

func (l *lampfishGroup) passDay() bool {
	if l.count == 0 {
		l.count = 6
		return true
	} else {
		l.count--
		return false
	}
}

func (l *lampfishGroup) spawn() lampfishGroup {
	newLampFish := lampfishGroup{count: 8, size: l.size}
	return newLampFish
}

func (l *lampfishGroup) print() {
	util.Dprintf("(%d:%d)", l.count, l.size)
}

func (l lampfishSchool) Len() int {
	return len(l)
}

func (l lampfishSchool) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l lampfishSchool) Less(i, j int) bool {
	return l[i].count < l[j].count
}

func (ls *lampfishSchool) passDayAllAndSpawn() lampfishSchool {
	var nls lampfishSchool = make(lampfishSchool, 0, len(*ls))
	for i := range *ls {
		if (*ls)[i].passDay() {
			nls = append(nls, (*ls)[i].spawn())
		}
	}
	return append(*ls, nls...)
}

func merge(ls lampfishSchool, start, end int) lampfishGroup {
	i := 0
	for i = start + 1; i < end; i++ {
		ls[start].size += ls[i].size
	}
	return ls[start]
}

func mergeAll(ls lampfishSchool) lampfishSchool {
	sort.Sort(ls)
	i := 0
	j := i + 1
	var nls lampfishSchool = make(lampfishSchool, 0, len(ls))
	for i = 0; i < len(ls); {
		for j = i; j < len(ls) && ls[i].count == ls[j].count; {
			j++
		}
		//util.Dprintf("%d %d\n", i, j)
		if j > i+1 {
			//util.Dprintf("%d %d\n", i, j)
			nls = append(nls, merge(ls, i, j))
			i = j
		} else {
			nls = append(nls, ls[i])
			i++
		}
	}
	return nls
}

func (ls *lampfishSchool) print() {
	for i := 0; i < len(*ls); i++ {
		(*ls)[i].print()
	}
	util.Dprintf("\n")
}

func runSimulation(ls lampfishSchool, ittr int) int64 {
	for i := 0; i < ittr; i++ {
		ls = ls.passDayAllAndSpawn()
		ls = mergeAll(ls)
		ls.print()
	}

	sum := int64(0)
	for _, f := range ls {
		sum += f.size
	}
	return sum
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var initialAges []string
	for file_scanner.Scan() {
		line := file_scanner.Text()
		initialAges = strings.Split(line, ",")
	}

	var lampFishes lampfishSchool = make(lampfishSchool, 0, len(initialAges))
	for _, age := range initialAges {
		ageNum, err := strconv.Atoi(age)
		util.Check_error(err)
		lfg := lampfishGroup{size: 1, count: ageNum}
		lampFishes = append(lampFishes, lfg)
	}

	lampFishes = mergeAll(lampFishes)
	lampFishes.print()

	count := runSimulation(lampFishes, 80)
	return fmt.Sprintf("%v", count)
}
