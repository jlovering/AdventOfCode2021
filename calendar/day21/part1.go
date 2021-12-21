package adventofcode

import (
	util "adventofcode/util/common"
	"fmt"
)

type detDie struct {
	last     int
	numRolls int
}

type player struct {
	num   int
	pos   int
	score int
}

func (d detDie) winStats() (int, int) {
	return d.last, d.numRolls
}

func (d *detDie) roll() int {
	d.last++
	if d.last > 100 {
		d.last = 1
	}
	d.numRolls++
	return d.last
}

func (d *detDie) rollAll() int {
	//util.Dprintf("%d %d\n", d.last, d.numRolls)
	//if d.numRolls >= 2 {
	//	ret = d.last + 1 + d.last + d.last + 1)
	//} else if d.numRolls == 1 {
	//	ret = 2 + 3 + d.last + 1
	//} else if d.numRolls == 0 {
	//	d.last = 4
	//	ret = 1 + 2 + 3
	//}

	d1 := d.roll()
	d2 := d.roll()
	d3 := d.roll()
	util.Dprintf("%d+%d+%d", d1, d2, d3)
	return d1 + d2 + d3
}

func (p *player) takeMove(d *detDie) {
	util.Dprintf("Player %d rolls ", p.num)
	roll := d.rollAll()
	p.pos = (p.pos + roll) % 10
	p.score += p.pos + 1
	util.Dprintf(" moves to %d score %d\n", p.pos+1, p.score)
}

func (p player) won() bool {
	return p.score >= 1000
}

func Part1(p1s, p2s int) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	die := detDie{}
	p1 := player{1, p1s - 1, 0}
	p2 := player{2, p2s - 1, 0}

	nextPlayer := &p1
	for !p1.won() && !p2.won() {
		nextPlayer.takeMove(&die)
		if nextPlayer == &p1 {
			nextPlayer = &p2
		} else {
			nextPlayer = &p1
		}
	}

	_, numRolls := die.winStats()
	lScore := 0
	if p1.won() {
		lScore = p2.score
	} else {
		lScore = p1.score
	}
	return fmt.Sprintf("%d", lScore*numRolls)
}
