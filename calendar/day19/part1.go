package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
	z int
}

type vector struct {
	dx int
	dy int
	dz int
}

type signature struct {
	mag1 uint
	//mag2 uint
	//mag3 uint
}

type signatureGeneration struct {
	p  point
	v1 vector
	//v2 vector
	//v3 vector
}

type rotationMap struct {
	xx int
	yx int
	zx int
	xy int
	yy int
	zy int
	xz int
	yz int
	zz int
}

type translationMap struct {
	xoffset int
	yoffset int
	zoffset int
}

type transformation struct {
	rMap rotationMap
	tMap translationMap
}

type scannerField map[point]interface{}

type scannerSignatures map[signature]signatureGeneration

type vectorMagPair struct {
	v   vector
	mag uint
}

type vectorMagPairSet []vectorMagPair

func (vps vectorMagPairSet) Len() int {
	return len(vps)
}

func (vps vectorMagPairSet) Less(i, j int) bool {
	return vps[i].mag < vps[j].mag
}

func (vps vectorMagPairSet) Swap(i, j int) {
	vps[i], vps[j] = vps[j], vps[i]
}

func (p point) vectorMagTo(pto point) vectorMagPair {
	dx := pto.x - p.x
	dy := pto.y - p.y
	dz := pto.z - p.z

	return vectorMagPair{vector{dx, dy, dz}, uint(dx*dx + dy*dy + dz*dz)}
}

func (sf scannerField) String() string {
	var sb strings.Builder
	for p := range sf {
		sb.WriteString(fmt.Sprintf("%d,%d,%d\n", p.x, p.y, p.z))
	}
	return sb.String()
}

func (sf scannerField) computeSignatures() scannerSignatures {
	scanSigs := scannerSignatures{}

	//vecMags := vectorMagPairSet{}
	for p1 := range sf {
		for p2 := range sf {
			if p1 == p2 {
				continue
			}
			vm := p1.vectorMagTo(p2)
			scanSigs[signature{vm.mag}] = signatureGeneration{p: p1, v1: vm.v}
			//if len(vecMags) < 3 || vm.mag < vecMags[2].mag {
			//	vecMags = append(vecMags, vm)
			//	sort.Sort(vecMags)
			//}
		}
		//scanSigs[signature{vecMags[0].mag, vecMags[1].mag, vecMags[2].mag}] = signatureGeneration{p: p1, v1: vecMags[0].v, v2: vecMags[1].v, v3: vecMags[2].v}
	}
	return scanSigs
}

func (p point) translate(tm translationMap) point {
	return point{x: p.x + tm.xoffset, y: p.y + tm.yoffset, z: p.z + tm.zoffset}
}

func (p point) rotate(rm rotationMap) point {
	p2x := p.x*rm.xx + p.y*rm.yx + p.z*rm.zx
	p2y := p.x*rm.xy + p.y*rm.yy + p.z*rm.zy
	p2z := p.x*rm.xz + p.y*rm.yz + p.z*rm.zz
	p2n := point{x: p2x, y: p2y, z: p2z}
	return p2n
}

func (p point) transform(trn transformation) point {
	p2 := p.rotate(trn.rMap)
	p3 := p2.translate(trn.tMap)
	return p3
}

func (sig1 signatureGeneration) computeTranslation(sig2 signatureGeneration, rm rotationMap) translationMap {
	p2n := sig2.p.rotate(rm)
	util.Dprintf("\t\t%v->%v (%v)\n", sig2.p, p2n, rm)
	dx := sig1.p.x - p2n.x
	dy := sig1.p.y - p2n.y
	dz := sig1.p.z - p2n.z
	return translationMap{xoffset: dx, yoffset: dy, zoffset: dz}
}

func (rm rotationMap) String() string {
	out := ""
	if rm.xx == 1 {
		out += "x->x"
	} else if rm.xx == -1 {
		out += "-x->x"
	} else if rm.yx == 1 {
		out += "y->x"
	} else if rm.yx == -1 {
		out += "-y->x"
	} else if rm.zx == 1 {
		out += "z->x"
	} else if rm.zx == -1 {
		out += "-z->x"
	}
	out += " "
	if rm.xy == 1 {
		out += "x->y"
	} else if rm.xy == -1 {
		out += "-x->y"
	} else if rm.yy == 1 {
		out += "y->y"
	} else if rm.yy == -1 {
		out += "-y->y"
	} else if rm.zy == 1 {
		out += "z->y"
	} else if rm.zy == -1 {
		out += "-z->y"
	}
	out += " "
	if rm.xz == 1 {
		out += "x->z"
	} else if rm.xz == -1 {
		out += "-x->z"
	} else if rm.yz == 1 {
		out += "y->z"
	} else if rm.yz == -1 {
		out += "-y->z"
	} else if rm.zz == 1 {
		out += "z->z"
	} else if rm.zz == -1 {
		out += "-z->z"
	}
	out += " "
	return out
}

func (sig1 signatureGeneration) computeRotation(sig2 signatureGeneration) (rotationMap, bool) {
	outRotation := rotationMap{}
	util.Dprintf("\tsig1.dx: %d %d %d Sig2: %d %d %d\n", sig1.v1.dx, sig1.v1.dy, sig1.v1.dz, sig2.v1.dx, sig2.v1.dy, sig2.v1.dz)
	switch sig1.v1.dx {
	case sig2.v1.dx:
		outRotation.xx = 1
	case -sig2.v1.dx:
		outRotation.xx = -1
	case sig2.v1.dy:
		outRotation.yx = 1
	case -sig2.v1.dy:
		outRotation.yx = -1
	case sig2.v1.dz:
		outRotation.zx = 1
	case -sig2.v1.dz:
		outRotation.zx = -1
	default:
		util.Dprintf("Couldn't match dx\n")
		return outRotation, false
	}
	switch sig1.v1.dy {
	case sig2.v1.dx:
		outRotation.xy = 1
	case -sig2.v1.dx:
		outRotation.xy = -1
	case sig2.v1.dy:
		outRotation.yy = 1
	case -sig2.v1.dy:
		outRotation.yy = -1
	case sig2.v1.dz:
		outRotation.zy = 1
	case -sig2.v1.dz:
		outRotation.zy = -1
	default:
		util.Dprintf("Couldn't match dy\n")
		return outRotation, false
	}
	switch sig1.v1.dz {
	case sig2.v1.dx:
		outRotation.xz = 1
	case -sig2.v1.dx:
		outRotation.xz = -1
	case sig2.v1.dy:
		outRotation.yz = 1
	case -sig2.v1.dy:
		outRotation.yz = -1
	case sig2.v1.dz:
		outRotation.zz = 1
	case -sig2.v1.dz:
		outRotation.zz = -1
	default:
		util.Dprintf("Couldn't match dz\n")
		return outRotation, false
	}

	util.Dprintf("\t%v\n", outRotation)
	return outRotation, true
}

func (sf *scannerField) mergeField(sf2 scannerField) {
	for k := range sf2 {
		(*sf)[k] = true
	}
}

func (sf scannerField) applyTransform(trn transformation) scannerField {
	nSF := scannerField{}
	for p := range sf {
		np := p.transform(trn)
		nSF[np] = true
	}
	return nSF
}

func (tm translationMap) String() string {
	return fmt.Sprintf("%d,%d,%d", tm.xoffset, tm.yoffset, tm.zoffset)
}

func (trn transformation) String() string {
	return fmt.Sprintf("%s, %s", trn.rMap, trn.tMap)
}

func (sig1 signatureGeneration) computeTransform(sig2 signatureGeneration) (transformation, bool) {
	rm, valid := sig1.computeRotation(sig2)
	if !valid {
		return transformation{}, false
	}
	tm := sig1.computeTranslation(sig2, rm)
	return transformation{tMap: tm, rMap: rm}, true
}

func (sigs1 scannerSignatures) intersection(sigs2 scannerSignatures) []signature {
	//util.Dprintf("%v\n%v\n", sigs1, sigs2)
	intersection := []signature{}
	for k1 := range sigs1 {
		if _, ok := sigs2[k1]; ok {
			intersection = append(intersection, k1)
		}
	}
	return intersection
}

func parseInput(file_scanner *bufio.Scanner) map[string]scannerField {
	scanners := map[string]scannerField{}

	curScanner := ""
	for file_scanner.Scan() {
		line := file_scanner.Text()
		//util.Dprintf("%s\n", line)
		if strings.Contains(line, "--- ") {
			s := strings.Split(line, " ")
			curScanner = s[2]
			scanners[curScanner] = scannerField{}
		} else if line != "" {
			s := strings.Split(line, ",")
			x, err := strconv.Atoi(s[0])
			util.Check_error(err)
			y, err := strconv.Atoi(s[1])
			util.Check_error(err)
			z, err := strconv.Atoi(s[2])
			util.Check_error(err)
			p := point{x, y, z}
			scanners[curScanner][p] = true
		} else {

		}
	}
	return scanners
}

type stringSlice []string

func (ss stringSlice) contains(f string) bool {
	for _, s := range ss {
		if s == f {
			return true
		}
	}
	return false
}

func Part1(filename string) string {
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

	return fmt.Sprintf("%d", len(masterPointMap))
}
