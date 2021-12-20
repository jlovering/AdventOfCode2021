package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
)

func (trn transformation) calculateManhatten(trn2 transformation) uint {
	x := uint(math.Abs(float64(trn.tMap.xoffset - trn2.tMap.xoffset)))
	y := uint(math.Abs(float64(trn.tMap.yoffset - trn2.tMap.yoffset)))
	z := uint(math.Abs(float64(trn.tMap.zoffset - trn2.tMap.zoffset)))

	return x + y + z
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	scannerData := parseInput(file_scanner)

	allScanSigs := map[string]scannerSignatures{}

	for n, sf := range scannerData {
		scanSigs := sf.computeSignatures()
		allScanSigs[n] = scanSigs
	}

	masterPointMap := scannerData["0"]
	masterSigs := allScanSigs["0"]

	matched := stringSlice{"0"}
	scanner_trns := []transformation{}

	for len(allScanSigs) > len(matched) {
		action := false
		for n, sigs := range allScanSigs {
			if matched.contains(n) {
				continue
			}
			sharedSigs := masterSigs.intersection(sigs)
			util.Dprintf("Checking %s: %d Matches\n", n, len(sharedSigs))
			if len(sharedSigs) >= 24 {
				util.IncreasedebugIndent()
				util.Dprintf("Scanner %s hits\n", n)
				//There is a chance that we have key overlaps, so we chose the best transform
				bestMatch := map[transformation]int{}
				for _, s := range sharedSigs {
					if trn, valid := masterSigs[s].computeTransform(sigs[s]); valid {
						bestMatch[trn]++
						util.Dprintf("%v\n", trn)
					} else {
						util.Dprintf("NT: %v\n", trn)
					}
				}
				//util.Dprintf("%d\n", len(bestMatch))
				var maxT transformation
				maxC := 0
				for k, v := range bestMatch {
					if v > maxC {
						maxT = k
						maxC = v
					}
				}
				if maxC >= 12 {
					util.Dprintf("%v %d\n", maxT, maxC)
					nSF := scannerData[n].applyTransform(maxT)
					util.Dprintf("%d %d\n", len(nSF), len(masterPointMap))
					masterPointMap.mergeField(nSF)
					util.Dprintf("%d\n", len(masterPointMap))
					masterSigs = masterPointMap.computeSignatures()
					util.Dprintf("%v\n", masterPointMap)
					matched = append(matched, n)
					scanner_trns = append(scanner_trns, maxT)
					action = true
				} else {
					util.Dprintf("NG best: %v %d\n", maxT, maxC)
				}
				util.DecreasedebugIndent()
			}
		}
		if !action {
			panic("Searched all sigs and no matches?")
		}
	}

	max_man := uint(0)
	for _, t1 := range scanner_trns {
		for _, t2 := range scanner_trns {
			man := t1.calculateManhatten(t2)
			if man > uint(max_man) {
				max_man = man
			}
		}
	}

	return fmt.Sprintf("%d", max_man)
}
