package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

type snail struct {
	number int
	left   *snail
	right  *snail
	parent *snail
}

type snailSet []*snail

func integrityCheck(sn *snail) {
	util.Setdebug(false)
	curSnail := sn

	stack := make([]*snail, 0)
	util.Dprintf("Checking %12p, %v\n", curSnail, *curSnail)

	stack = append(stack, curSnail.parent)
	panicLater := false
	util.Dprintf("\tDown Left\n")
	for curSnail.left != nil {
		util.Dprintf("\t\t%12p %12p %12p %12p %v\n", curSnail, curSnail.left, curSnail.right, curSnail.parent, stack)
		stack = append(stack, curSnail)
		curSnail = curSnail.left
	}
	util.Dprintf("\t\t%12p 0x0000000000 0x0000000000 %12p %d %v\n", curSnail, curSnail.parent, curSnail.number, stack)
	util.Dprintf("\tUp Left\n")
	for curSnail != nil {
		util.Dprintf("\t\t%12p %12p %12p %12p %v\n", curSnail, curSnail.left, curSnail.right, curSnail.parent, stack)
		curSnail = curSnail.parent
		if curSnail != stack[len(stack)-1] {
			util.Dprintf("Stack MissMatched: %p != %p\n", curSnail, stack[len(stack)-1])
			panicLater = true
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) != 0 {
		panic("Didn't make it back")
	}

	curSnail = sn
	stack = append(stack, curSnail.parent)
	util.Dprintf("\tDown Right\n")
	for curSnail.right != nil {
		util.Dprintf("\t\t%12p %12p %12p %12p %v\n", curSnail, curSnail.left, curSnail.right, curSnail.parent, stack)
		stack = append(stack, curSnail)
		curSnail = curSnail.right
	}
	util.Dprintf("\t\t%12p 0x0000000000 0x0000000000 %12p %d %v\n", curSnail, curSnail.parent, curSnail.number, stack)
	util.Dprintf("\tUp Right\n")
	for curSnail != nil {
		util.Dprintf("\t\t%12p %12p %12p %12p %v\n", curSnail, curSnail.left, curSnail.right, curSnail.parent, stack)
		curSnail = curSnail.parent
		if curSnail != stack[len(stack)-1] {
			util.Dprintf("Stack MissMatched: %p != %p\n", curSnail, stack[len(stack)-1])
			panicLater = true
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) != 0 {
		panic("Didn't make it back")
	}

	if panicLater {
		panic(1)
	}
}

func printSnail(sn *snail) string {
	var curSnail *snail = sn

	str := ""
	if curSnail.right == nil && curSnail.left == nil {
		str += fmt.Sprintf("%d", curSnail.number)
		return str
	}
	if curSnail.left != nil {
		str += "[" + printSnail(curSnail.left)
	}
	if curSnail.right != nil {
		str += ","
		str += printSnail(curSnail.right)
	}
	str += "]"
	return str
}

func snailParse(st string) *snail {
	util.Setdebug(false)
	curSnail := &snail{}

	util.Dprintf("%p %p\n", curSnail, curSnail.parent)

	var err error
	for _, c := range st {
		util.Dprintf("%c ", c)
		if c == '[' {
			newSnail := &snail{}
			newSnail.parent = curSnail
			curSnail.left = newSnail
			curSnail = newSnail
			util.Dprintf("left ")
		} else if unicode.IsDigit(c) {
			num := 0
			num, err = strconv.Atoi(string(c))
			curSnail.number *= 10
			curSnail.number += num
			util.Check_error(err)
			util.Dprintf("num %d ", curSnail.number)
		} else if c == ',' {
			curSnail = curSnail.parent
			newSnail := &snail{}
			newSnail.parent = curSnail
			curSnail.right = newSnail
			curSnail = newSnail
			util.Dprintf("right ")
		} else if c == ']' {
			curSnail = curSnail.parent
		}
		util.Dprintf("%p %p\n", curSnail, curSnail.parent)
	}
	util.Dprintf("%s\n", printSnail(curSnail))
	integrityCheck(curSnail)
	return curSnail
}

func snailFindRight(sn *snail) (*snail, bool) {
	if sn.parent == nil {
		panic("Can't find right on orphan")
	}

	prevSnail := sn
	curSnail := sn.parent

	for (curSnail.right == nil || curSnail.right == prevSnail) && curSnail.parent != nil {
		prevSnail = curSnail
		curSnail = curSnail.parent
	}
	if (curSnail.right == nil || curSnail.right == prevSnail) && curSnail.parent == nil {
		return nil, false
	}
	curSnail = curSnail.right
	for curSnail.left != nil {
		curSnail = curSnail.left
	}
	return curSnail, true
}

func snailFindLeft(sn *snail) (*snail, bool) {
	if sn.parent == nil {
		panic("Can't find right on orphan")
	}
	prevSnail := sn
	curSnail := sn.parent

	for (curSnail.left == nil || curSnail.left == prevSnail) && curSnail.parent != nil {
		prevSnail = curSnail
		curSnail = curSnail.parent
	}
	if curSnail.left == nil || curSnail.left == prevSnail {
		return nil, false
	}
	curSnail = curSnail.left
	for curSnail.right != nil {
		curSnail = curSnail.right
	}
	return curSnail, true
}

func snailFindDepth(sn *snail) int {
	depth := 0
	for curSnail := sn; curSnail.parent != nil; curSnail = curSnail.parent {
		depth++
	}
	return depth
}

func snailSplit(sn *snail) {
	if sn.left != nil || sn.right != nil {
		panic("Cannot split real number\n")
	}
	splitNum := float64(sn.number) / 2
	rndDown := math.Floor(splitNum)
	rndUp := math.Ceil(splitNum)

	//util.Dprintf("%f %f %f\n", splitNum, rndDown, rndUp)
	sn.left = &snail{
		number: int(rndDown),
		left:   nil,
		right:  nil,
		parent: sn}
	sn.right = &snail{
		number: int(rndUp),
		left:   nil,
		right:  nil,
		parent: sn}

	//if snailFindDepth(sn.left) > 4 {
	//	util.Dprintf("Resplode! %s\n", printSnail(sn))
	//	snailExplode(sn)
	//}
}

func createSnailPair(x, y int) *snail {
	newSnail := &snail{}
	newSnail.left = &snail{number: x}
	newSnail.right = &snail{number: y}
	return newSnail
}

func snailExplode(sn *snail) {
	util.Setdebug(false)
	util.IncreasedebugIndent()
	defer util.DecreasedebugIndent()
	if sn.left == nil || sn.right == nil {
		panic("Can't explode none snail")
	}

	pSnail := sn.parent

	if pSnail == nil {
		panic("Parent of explode node is nil?")
	}

	sn.number = 0

	rightSnail, rfound := snailFindRight(sn)
	if !rfound {
		util.Dprintf("No right, chuck\n")
	} else {
		util.Dprintf("Targeting R: %s (%s)\n", printSnail(rightSnail), printSnail(rightSnail.parent))
		rightSnail.number += sn.right.number
		util.Dprintf("Result R: %s (%s)\n", printSnail(rightSnail), printSnail(rightSnail.parent))
	}
	sn.right = nil

	leftSnail, lfound := snailFindLeft(sn)
	if !lfound {
		util.Dprintf("No left, chuck\n")
	} else {
		util.Dprintf("Targeting L: %s (%s)\n", printSnail(leftSnail), printSnail(leftSnail.parent))
		leftSnail.number += sn.left.number
		util.Dprintf("Result L: %s (%s)\n", printSnail(leftSnail), printSnail(leftSnail.parent))
	}
	sn.left = nil

	util.Dprintf("Final: %s\n", printSnail(pSnail))
}

func snailReduce(sn *snail) (*snail, bool) {
	util.Setdebug(false)
	util.IncreasedebugIndent()
	defer util.DecreasedebugIndent()
	if sn.left == nil || sn.right == nil {
		panic("Can't reduce none snail")
	}

	curSnail := sn
	depth := 0
	for {
		util.Dprintf("Cur: %s\n", printSnail(curSnail))
		util.IncreasedebugIndent()
		for curSnail.left != nil {
			curSnail = curSnail.left
			depth++
			util.Dprintf("SL: %s (%d)\n", printSnail(curSnail), depth)
		}
		if depth > 4 {
			util.Dprintf("Exploding: %s (%s) %d\n", printSnail(curSnail.parent), printSnail(curSnail.parent.parent), depth)
			snailExplode(curSnail.parent)
			util.DecreasedebugIndent()
			return sn, true
		}
		prevSnail := curSnail
		for (curSnail.right == nil || curSnail.right == prevSnail) && curSnail.parent != nil {
			prevSnail = curSnail
			curSnail = curSnail.parent
			depth--
			util.Dprintf("SU: %s\n", printSnail(curSnail))
		}
		if curSnail.right != nil && curSnail.right != prevSnail {
			curSnail = curSnail.right
			depth++
			util.Dprintf("SR: %s\n", printSnail(curSnail))
		}
		if curSnail.parent == nil {
			util.Dprintf("At P\n")
			util.DecreasedebugIndent()
			break
		}
		util.DecreasedebugIndent()
	}

	for {
		util.Dprintf("Cur: %s\n", printSnail(curSnail))
		util.IncreasedebugIndent()
		for curSnail.left != nil {
			curSnail = curSnail.left
			depth++
			util.Dprintf("SL: %s (%d)\n", printSnail(curSnail), depth)
		}
		if curSnail.number >= 10 {
			util.Dprintf("Splitting: %s, (%s)\n", printSnail(curSnail), printSnail(curSnail.parent))
			snailSplit(curSnail)
			util.DecreasedebugIndent()
			return sn, true
		}
		prevSnail := curSnail
		for (curSnail.right == nil || curSnail.right == prevSnail) && curSnail.parent != nil {
			prevSnail = curSnail
			curSnail = curSnail.parent
			depth--
			util.Dprintf("SU: %s\n", printSnail(curSnail))
		}
		if curSnail.right != nil && curSnail.right != prevSnail {
			curSnail = curSnail.right
			depth++
			util.Dprintf("SR: %s\n", printSnail(curSnail))
		}
		if curSnail.parent == nil {
			util.Dprintf("At P\n")
			util.DecreasedebugIndent()
			break
		}
		util.DecreasedebugIndent()
	}
	return sn, false
}

func reduceAll(sn *snail) *snail {
	action := false
	for {
		//util.Dprintf("%s\n", printSnail(sn))
		sn, action = snailReduce(sn)
		if !action {
			break
		}
		//util.Dprintf("%s\n\n", printSnail(sn))
	}
	return sn
}

func snailMagnitude(sn *snail) int {
	if sn.left == nil && sn.right == nil {
		return sn.number
	} else {
		return 3*snailMagnitude(sn.left) + 2*snailMagnitude(sn.right)
	}
}

func snailAdd(s1, s2 *snail) *snail {
	newSnail := &snail{}
	s1.parent = newSnail
	s2.parent = newSnail
	newSnail.left = s1
	newSnail.right = s2
	//util.Dprintf("%p %v %s\n", s1, *s1, printSnail(*s1))
	//util.Dprintf("%p %v %s\n", s2, *s2, printSnail(*s2))
	//util.Dprintf("%p %v\n", newSnail, newSnail)
	//util.Dprintf("%p %v %s\n", newSnail.left, *newSnail.left, printSnail(*newSnail.left))
	//util.Dprintf("%p %v %s\n", newSnail.right, *newSnail.right, printSnail(*newSnail.right))
	//util.Dprintf("%s\n", printSnail(*newSnail))
	//for curSn := newSnail; curSn.left != nil; curSn = curSn.left {
	//	util.Dprintf("%p %p\n", curSn, curSn.parent)
	//}
	integrityCheck(newSnail)
	return newSnail
}

func addAll(ssn snailSet) *snail {
	util.Setdebug(false)
	if len(ssn) < 2 {
		panic("Can't add set of less than 2")
	}
	ns := ssn[0]
	for i := 1; i < len(ssn); i++ {
		util.Dprintf("%s + %s\n", printSnail(ns), printSnail(ssn[i]))
		ns = snailAdd(ns, ssn[i])
		util.Dprintf("%s\n\n", printSnail(ns))
		nns := reduceAll(ns)
		util.Dprintf("%s\n\n", printSnail(ns))
		ns = nns
	}
	return ns
}

func parseInput(file_scanner *bufio.Scanner) snailSet {
	nss := make(snailSet, 0)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		snail := snailParse(line)
		integrityCheck(snail)
		nss = append(nss, snail)
		integrityCheck(nss[len(nss)-1])
	}
	return nss
}

func parseInputForTest(filename string) snailSet {
	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()
	file_scanner := bufio.NewScanner(f)
	return parseInput(file_scanner)
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	allParsed := parseInput(file_scanner)
	added := addAll(allParsed)
	//util.Dprintf("%s\n", printSnail(added))

	return fmt.Sprintf("%d", snailMagnitude(added))
}
