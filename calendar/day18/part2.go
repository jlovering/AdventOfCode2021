package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	nss := make([]string, 0)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		nss = append(nss, line)
	}

	mags := make([]int, 0)
	for sfn1 := range nss {
		for sfn2 := range nss {
			if sfn1 == sfn2 {
				continue
			}

			sn1 := snailParse(nss[sfn1])
			sn2 := snailParse(nss[sfn2])
			util.Setdebug(true)
			util.Dprintf("%s + %s = ", printSnail(sn1), printSnail(sn2))
			util.Setdebug(false)
			cand := addAll(snailSet{sn1, sn2})
			mag := snailMagnitude(cand)
			util.Setdebug(true)
			util.Dprintf("%s (%s)\n", printSnail(cand), mag)
			util.Setdebug(false)
			mags = append(mags, mag)

			sn1 = snailParse(nss[sfn1])
			sn2 = snailParse(nss[sfn2])
			util.Setdebug(true)
			util.Dprintf("%s + %s = ", printSnail(sn1), printSnail(sn2))
			util.Setdebug(false)
			cand = addAll(snailSet{sn2, sn1})
			mag = snailMagnitude(cand)
			mags = append(mags, mag)
			util.Setdebug(true)
			util.Dprintf("%s (%s)\n", printSnail(cand), mag)
			util.Setdebug(false)
		}
	}
	sort.Ints(mags)

	return fmt.Sprintf("%d", mags[len(mags)-1])
}
