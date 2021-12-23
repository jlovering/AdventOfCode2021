package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/beefsack/go-astar"
)

var movementCosts = map[rune]uint{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

//HallwaySpots
var hallWaySpots = map[string]int{
	"LeftNook":  0,
	"LeftGap":   1,
	"roomA":     2,
	"AB":        3,
	"roomB":     4,
	"BC":        5,
	"roomC":     6,
	"CD":        7,
	"roomD":     8,
	"RightGap":  9,
	"RightNook": 10,
}

//HallwaySpots
var hallWaySpotsName = map[int]string{
	0:  "LeftNook",
	1:  "LeftGap",
	2:  "roomA",
	3:  "AB",
	4:  "roomB",
	5:  "BC",
	6:  "roomC",
	7:  "CD",
	8:  "roomD",
	9:  "RightGap",
	10: "RightNook",
}

var validHallWaySpots = []string{
	"LeftNook",
	"LeftGap",
	"AB",
	"BC",
	"CD",
	"RightGap",
	"RightNook",
}

var rooms = []string{
	"roomD",
	"roomC",
	"roomB",
	"roomA",
}

func hallwayDistance(s, e string) uint {
	return uint(math.Abs(float64(hallWaySpots[s] - hallWaySpots[e])))
}

func travelCost(s, e string, pod rune) uint {
	return hallwayDistance(s, e) * movementCosts[pod]
}

func (b burrow) validPath(s, e string) bool {
	db := util.Getdebug()
	util.Setdebug(false)
	defer util.Setdebug(db)
	util.IncreasedebugIndent()
	util.Dprintf("%s -> %s", s, e)
	if hallWaySpots[s] < hallWaySpots[e] {
		for i := hallWaySpots[s] + 1; i <= hallWaySpots[e]; i++ {
			if b.hallway[i] != 0 {
				//Something in our way
				util.Dprintf(" NG obstical at %d\n", i)
				util.DecreasedebugIndent()
				return false
			}
		}
	} else {
		for i := hallWaySpots[s] - 1; i >= hallWaySpots[e]; i-- {
			if b.hallway[i] != 0 {
				//Something in our way
				util.Dprintf(" NG obstical at %d\n", i)
				util.DecreasedebugIndent()
				return false
			}
		}
	}
	util.Dprintf(" Good!\n")
	util.DecreasedebugIndent()
	return true
}

var minCostsMap = map[string]uint{
	"A-roomA": 0,
	"B-roomA": 4*movementCosts['B'] + travelCost("roomA", "roomB", 'B'),
	"C-roomA": 4*movementCosts['C'] + travelCost("roomA", "roomC", 'C'),
	"D-roomA": 4*movementCosts['D'] + travelCost("roomA", "roomD", 'D'),

	"A-roomB": 4*movementCosts['A'] + travelCost("roomB", "roomA", 'A'),
	"B-roomB": 0,
	"C-roomB": 4*movementCosts['C'] + travelCost("roomB", "roomC", 'C'),
	"D-roomB": 4*movementCosts['D'] + travelCost("roomB", "roomD", 'D'),

	"A-roomC": 4*movementCosts['A'] + travelCost("roomC", "roomA", 'A'),
	"B-roomC": 4*movementCosts['B'] + travelCost("roomC", "roomB", 'B'),
	"C-roomC": 0,
	"D-roomC": 4*movementCosts['D'] + travelCost("roomC", "roomD", 'D'),

	"A-roomD": 4*movementCosts['A'] + travelCost("roomD", "roomA", 'A'),
	"B-roomD": 4*movementCosts['B'] + travelCost("roomD", "roomB", 'B'),
	"C-roomD": 4*movementCosts['C'] + travelCost("roomD", "roomC", 'C'),
	"D-roomD": 0,
}

func (b burrow) costEstimator() uint {
	cost := uint(0)
	for i := 0; i < b.roomA.count; i++ {
		cost += minCostsMap[string(b.roomA.room[i])+"-roomA"] //+ movementCosts[c]*uint(i)
	}
	for i := 0; i < b.roomB.count; i++ {
		cost += minCostsMap[string(b.roomB.room[i])+"-roomB"] //+ movementCosts[c]*uint(i)
	}
	for i := 0; i < b.roomC.count; i++ {
		cost += minCostsMap[string(b.roomC.room[i])+"-roomC"] //+ movementCosts[c]*uint(i)
	}
	for i := 0; i < b.roomD.count; i++ {
		cost += minCostsMap[string(b.roomD.room[i])+"-roomD"] //+ movementCosts[c]*uint(i)
	}

	mult := uint(1)
	hallCost := uint(0)
	for i, c := range b.hallway {
		if c != 0 {
			hallCost += travelCost(hallWaySpotsName[i], "room"+string(c), c) * movementCosts[c]
			mult *= 1
		}
	}
	cost += hallCost * mult
	return cost
}

func (b burrow) integrityCheck() {
	aC := 0
	bC := 0
	cC := 0
	dC := 0

	for _, c := range b.roomA.room {
		switch c {
		case 'A':
			aC++
		case 'B':
			bC++
		case 'C':
			cC++
		case 'D':
			dC++
		}
	}
	for _, c := range b.roomB.room {
		switch c {
		case 'A':
			aC++
		case 'B':
			bC++
		case 'C':
			cC++
		case 'D':
			dC++
		}
	}
	for _, c := range b.roomC.room {
		switch c {
		case 'A':
			aC++
		case 'B':
			bC++
		case 'C':
			cC++
		case 'D':
			dC++
		}
	}
	for _, c := range b.roomD.room {
		switch c {
		case 'A':
			aC++
		case 'B':
			bC++
		case 'C':
			cC++
		case 'D':
			dC++
		}
	}
	for _, c := range b.hallway {
		switch c {
		case 'A':
			aC++
		case 'B':
			bC++
		case 'C':
			cC++
		case 'D':
			dC++
		}
	}

	if aC != b.roomA.size || bC != b.roomB.size || cC != b.roomC.size || dC != b.roomC.size {
		util.Dprintf("%v\n", b)
		panic("Shrimp escape!")
	}
}

func (b burrow) deepCopy() burrow {
	nB := burrow{
		hallway: b.hallway,
		roomA: burrowRoom{
			room:  b.roomA.room,
			count: b.roomA.count,
			size:  b.roomA.size,
			lock:  b.roomA.lock,
		},
		roomB: burrowRoom{
			room:  b.roomB.room,
			count: b.roomB.count,
			size:  b.roomB.size,
			lock:  b.roomB.lock,
		},
		roomC: burrowRoom{
			room:  b.roomC.room,
			count: b.roomC.count,
			size:  b.roomC.size,
			lock:  b.roomC.lock,
		},
		roomD: burrowRoom{
			room:  b.roomD.room,
			count: b.roomD.count,
			size:  b.roomD.size,
			lock:  b.roomD.lock,
		},
	}

	return nB
}

func roomInValidity(br *burrowRoom, pod rune) bool {
	r := br.count > 0 && !br.lock
	for i := 0; i < br.count; i++ {
		r = r && (br.room[i] != 0 && br.room[i] != pod)
	}
	return r
}

func (b burrow) invalid() bool {
	A := roomInValidity(&b.roomA, 'A')
	B := roomInValidity(&b.roomB, 'B')
	C := roomInValidity(&b.roomC, 'C')
	D := roomInValidity(&b.roomD, 'D')
	//util.Dprintf("%v %v %v %v\n", A, B, C, D)

	return A || B || C || D
}

type burrowRoom struct {
	room  [4]rune
	lock  bool
	count int
	size  int
}

type burrow struct {
	hallway [11]rune
	roomA   burrowRoom
	roomB   burrowRoom
	roomC   burrowRoom
	roomD   burrowRoom
}

func (b burrow) String() string {
	var sb strings.Builder
	sb.WriteString("#############\n")
	sb.WriteString(util.GetDebugIndent() + "#")
	for _, c := range b.hallway {
		if c != 0 {
			sb.WriteRune(c)
		} else {
			sb.WriteRune('.')
		}
	}
	sb.WriteString("#\n")
	for i := b.roomA.size - 1; i >= 0; i-- {
		if i == b.roomA.size-1 {
			sb.WriteString(util.GetDebugIndent() + "###")
		} else {
			sb.WriteString(util.GetDebugIndent() + "  #")
		}
		if b.roomA.room[i] != 0 {
			sb.WriteRune(b.roomA.room[i])
		} else {
			sb.WriteRune('.')
		}
		sb.WriteRune('#')
		if b.roomB.room[i] != 0 {
			sb.WriteRune(b.roomB.room[i])
		} else {
			sb.WriteRune('.')
		}
		sb.WriteRune('#')
		if b.roomC.room[i] != 0 {
			sb.WriteRune(b.roomC.room[i])
		} else {
			sb.WriteRune('.')
		}
		sb.WriteRune('#')
		if b.roomD.room[i] != 0 {
			sb.WriteRune(b.roomD.room[i])
		} else {
			sb.WriteRune('.')
		}
		if i == b.roomA.size-1 {
			sb.WriteString("###\n")
		} else {
			sb.WriteString("#\n")
		}
	}
	sb.WriteString(util.GetDebugIndent() + "  #########\n")
	sb.WriteString(fmt.Sprintf(util.GetDebugIndent()+"   %d %d %d %d\n", b.roomA.count, b.roomB.count, b.roomC.count, b.roomD.count))
	sb.WriteString(fmt.Sprintf(util.GetDebugIndent()+"   %d %d %d %d\n", b.roomA.size, b.roomB.size, b.roomC.size, b.roomD.size))
	sb.WriteString(fmt.Sprintf(util.GetDebugIndent()+"   %1v %1v %1v %1v\n", b.roomA.lock, b.roomB.lock, b.roomC.lock, b.roomD.lock))
	sb.WriteString(fmt.Sprintf(util.GetDebugIndent()+"   %d", b.costEstimator()))
	return sb.String()
}

type graphEdge struct {
	node *graphNode
	cost uint
}

type graphNode struct {
	state burrow
	edges []graphEdge
}

func roomInMagic(room *burrowRoom, pod rune) {
	//util.Dprintf("%v %d %d %c\n", room, room.count, room.size, pod)
	(*room).room[room.count] = pod
	room.count++
	if room.count > room.size {
		panic("Overcap, call the fire marshal")
	}
}

func roomIn(room *burrowRoom, pod rune) uint {
	if room.count == room.size {
		return 0
	}

	(*room).room[room.count] = pod
	cost := movementCosts[pod] * uint(room.size-room.count)
	room.count++
	return cost
}

func roomOut(room *burrowRoom) (rune, uint) {
	//util.Dprintf("%d %d\n", room.count, room.size)
	if room.count == 0 {
		return 0, 0
	}
	if !room.lock {
		return 0, 0
	}
	pod := (*room).room[room.count-1]
	(*room).room[room.count-1] = 0
	cost := movementCosts[pod] * uint(room.size+1-room.count)
	room.count--
	return pod, cost
}

func tidyRoom(room *burrowRoom, pod rune) {
	if !room.lock {
		return
	}
	if room.count == 0 {
		room.lock = false
		return
	}

	unlock := true
	for i := 0; i < room.count; i++ {
		unlock = unlock && (*room).room[i] == pod
	}
	if unlock {
		room.lock = false
	}
}

func (b *burrow) popOutOfRoom(room string) (rune, uint) {
	switch room {
	case "roomA":
		r, c := roomOut(&b.roomA)
		tidyRoom(&b.roomA, 'A')
		return r, c
	case "roomB":
		r, c := roomOut(&b.roomB)
		tidyRoom(&b.roomB, 'B')
		return r, c
	case "roomC":
		r, c := roomOut(&b.roomC)
		tidyRoom(&b.roomC, 'C')
		return r, c
	case "roomD":
		r, c := roomOut(&b.roomD)
		tidyRoom(&b.roomD, 'D')
		return r, c
	}
	panic("Red room?")
}

func (b burrow) checkRoom(room string, pod rune) bool {
	switch room {
	case "roomA":
		if pod != 'A' {
			panic("Not X in the X's room")
		}
		return !b.roomA.lock && b.roomA.count < b.roomA.size
	case "roomB":
		if pod != 'B' {
			panic("Not X in the X's room")
		}
		return !b.roomB.lock && b.roomB.count < b.roomB.size
	case "roomC":
		if pod != 'C' {
			panic("Not X in the X's room")
		}
		return !b.roomC.lock && b.roomC.count < b.roomC.size
	case "roomD":
		if pod != 'D' {
			panic("Not X in the X's room")
		}
		return !b.roomD.lock && b.roomD.count < b.roomD.size
	}
	panic("Red room?")
}

func (b *burrow) popIntoRoom(room string, pod rune) uint {
	switch room {
	case "roomA":
		return roomIn(&b.roomA, pod)
	case "roomB":
		return roomIn(&b.roomB, pod)
	case "roomC":
		return roomIn(&b.roomC, pod)
	case "roomD":
		return roomIn(&b.roomD, pod)
	}
	panic("Red room?")
}

func (b *burrow) goToYourRoom(room string, pod rune) {
	switch room {
	case "roomA":
		roomInMagic(&b.roomA, pod)
		return
	case "roomB":
		roomInMagic(&b.roomB, pod)
		return
	case "roomC":
		roomInMagic(&b.roomC, pod)
		return
	case "roomD":
		roomInMagic(&b.roomD, pod)
		return
	}
	panic("I hate this place I'm leaving!")
}

var nastyNodeMap map[burrow]*graphNode = make(map[burrow]*graphNode)

func (b burrow) validMoves() []graphEdge {
	db := util.Getdebug()
	util.Setdebug(false)
	defer util.Setdebug(db)
	util.Dprintf("\n%v\n", b)
	edges := []graphEdge{}
	//First move any in the hallway
	for i := range b.hallway {
		if c := b.hallway[i]; c != 0 {
			util.IncreasedebugIndent()
			util.Dprintf("Considering %c in hallway: ", c)
			s := hallWaySpotsName[i]
			e := "room" + string(c)
			if b.validPath(s, e) && b.checkRoom(e, c) {
				util.Dprintf("can get to %s\n", e)
				nB := b.deepCopy()
				nB.hallway[i] = 0
				cost := travelCost(s, e, c) + nB.popIntoRoom(e, c)
				nB.integrityCheck()
				if nB.invalid() {
					util.Dprintf("Dest State Invalid\n")
					util.DecreasedebugIndent()
					continue
				}
				var nG *graphNode
				if eG, e := nastyNodeMap[nB]; e {
					nG = eG
				} else {
					nG = &graphNode{
						state: nB,
					}
					nastyNodeMap[nB] = nG
				}
				nGe := graphEdge{
					cost: cost,
					node: nG,
				}
				util.Dprintf("%v\n", nB)

				edges = append(edges, nGe)
			} else {
				util.Dprintf("%s NG\n", e)
			}
			util.DecreasedebugIndent()
		}
	}
	//Now deal with poping out of rooms
	for _, s := range rooms {
		util.IncreasedebugIndent()
		nB := b.deepCopy()
		c, cost := nB.popOutOfRoom(s)

		if c != 0 {
			util.Dprintf("Considering %c in from %s\n", c, s)

			//First check the direct root
			e := "room" + string(c)
			util.Dprintf("%v\n", nB)
			if nB.validPath(s, e) && nB.checkRoom(e, c) {
				util.Dprintf("can get to %s (direct)\n", e)
				nnB := nB.deepCopy()
				nCost := cost + travelCost(s, e, c) + nnB.popIntoRoom(e, c)
				nnB.integrityCheck()
				if nnB.invalid() {
					util.Dprintf("Dest State Invalid:\n")
					util.Dprintf("%v\n", nnB)
				} else {
					var nG *graphNode
					if eG, e := nastyNodeMap[nnB]; e {
						nG = eG
					} else {
						nG = &graphNode{
							state: nnB,
						}
						nastyNodeMap[nnB] = nG
					}
					nGe := graphEdge{
						cost: nCost,
						node: nG,
					}
					util.Dprintf("%v\n", nnB)

					edges = append(edges, nGe)
				}
			} else {
				util.Dprintf("Direct NG\n")
			}

			for _, e := range validHallWaySpots {
				if nB.validPath(s, e) {
					util.Dprintf("heading to %s\n", e)
					nnB := nB.deepCopy()
					nnB.hallway[hallWaySpots[e]] = c
					nnB.integrityCheck()
					if nnB.invalid() {
						util.Dprintf("Dest State Invalid\n")
						util.Dprintf("%v\n", nnB)
						continue
					}
					var nG *graphNode
					if eG, e := nastyNodeMap[nnB]; e {
						nG = eG
					} else {
						nG = &graphNode{
							state: nnB,
						}
						nastyNodeMap[nnB] = nG
					}
					nGe := graphEdge{
						cost: cost + travelCost(s, e, c),
						node: nG,
					}
					util.Dprintf("%v\n", nnB)
					edges = append(edges, nGe)
				} else {
					util.Dprintf("%s NG\n", e)
				}
			}
		}
		util.Dprintf("\n")
		util.DecreasedebugIndent()
	}

	return edges
}

func (b *graphNode) PathNeighbors() []astar.Pather {
	var neigh []astar.Pather = make([]astar.Pather, 0)
	b.edges = b.state.validMoves()
	for _, e := range b.edges {
		neigh = append(neigh, e.node)
	}
	return neigh
}

func (b *graphNode) PathNeighborCost(to astar.Pather) float64 {
	b2 := to.(*graphNode)
	for _, e := range b.edges {
		if b2.state == e.node.state {
			return float64(e.cost)
		}
	}
	return math.MaxFloat64
}

func (t *graphNode) PathEstimatedCost(to astar.Pather) float64 {
	return float64(t.state.costEstimator())
}

func parseInput(file_scanner *bufio.Scanner) burrow {

	barrow := burrow{}

	file_scanner.Scan()
	_ = file_scanner.Text()
	file_scanner.Scan()
	_ = file_scanner.Text()
	file_scanner.Scan()
	s1 := file_scanner.Text()
	file_scanner.Scan()
	s2 := file_scanner.Text()

	barrow.roomA.size = 2
	barrow.roomB.size = 2
	barrow.roomC.size = 2
	barrow.roomD.size = 2

	barrow.goToYourRoom("roomA", rune(s2[3]))
	barrow.goToYourRoom("roomB", rune(s2[5]))
	barrow.goToYourRoom("roomC", rune(s2[7]))
	barrow.goToYourRoom("roomD", rune(s2[9]))
	barrow.goToYourRoom("roomA", rune(s1[3]))
	barrow.goToYourRoom("roomB", rune(s1[5]))
	barrow.goToYourRoom("roomC", rune(s1[7]))
	barrow.goToYourRoom("roomD", rune(s1[9]))

	barrow.roomA.lock = true
	barrow.roomB.lock = true
	barrow.roomC.lock = true
	barrow.roomD.lock = true
	return barrow
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	b := parseInput(file_scanner)
	util.Dprintf("%v\n", b)

	start := graphNode{
		state: b,
	}

	end := graphNode{
		state: burrow{
			hallway: [11]rune{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			roomA: burrowRoom{
				room:  [4]rune{'A', 'A'},
				lock:  false,
				count: 2,
				size:  2,
			},
			roomB: burrowRoom{
				room:  [4]rune{'B', 'B'},
				lock:  false,
				count: 2,
				size:  2,
			},
			roomC: burrowRoom{
				room:  [4]rune{'C', 'C'},
				lock:  false,
				count: 2,
				size:  2,
			},
			roomD: burrowRoom{
				room:  [4]rune{'D', 'D'},
				lock:  false,
				count: 2,
				size:  2,
			},
		},
	}

	util.Dprintf("%v\n", end.state)
	nastyNodeMap[end.state] = &end

	path, distance, found := astar.Path(&start, &end)

	if !found {
		panic("No path")
	}

	for i := range path {
		gn := path[len(path)-1-i].(*graphNode)
		fmt.Printf("%v\n", gn.state)
	}

	return fmt.Sprintf("%d", uint(distance))
}
