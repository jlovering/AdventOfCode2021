package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type subPackFrame2 struct {
	op       int
	lty      int
	bitEx    int
	bitUsed  int
	packEx   int
	packUsed int
	values   []uint64
}

func sum(values []uint64) uint64 {
	sum := uint64(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

func product(values []uint64) uint64 {
	prod := uint64(1)
	for _, v := range values {
		prod *= v
	}
	return prod
}

func min(values []uint64) uint64 {
	min := uint64(math.MaxUint64)
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func max(values []uint64) uint64 {
	max := uint64(0)
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func grtn(values []uint64) uint64 {
	if len(values) != 2 {
		panic(3)
	}
	if values[0] > values[1] {
		return 1
	}
	return 0
}

func ltn(values []uint64) uint64 {
	if len(values) != 2 {
		panic(3)
	}
	if values[0] < values[1] {
		return 1
	}
	return 0
}

func equal(values []uint64) uint64 {
	if len(values) != 2 {
		panic(3)
	}
	if values[0] == values[1] {
		return 1
	}
	return 0
}

func performOp(op int, values []uint64) uint64 {
	util.Dprintf("Performing %d on %v\n", op, values)
	switch op {
	case 0:
		return sum(values)
	case 1:
		return product(values)
	case 2:
		return min(values)
	case 3:
		return max(values)
	case 5:
		return grtn(values)
	case 6:
		return ltn(values)
	case 7:
		return equal(values)
	}
	panic(4)
}

func parsePacket2(line string) int {
	const (
		HEADER = iota
		LITERAL
		OPERATOR
		LEN_OP_CLEAN
		COUNT_OP_CLEAN
	)

	const (
		LENGTH_TYPE = iota
		COUNT_TYPE
	)

	subPacketLenstack := make([]subPackFrame2, 1)
	subPacketLenstack[0] = subPackFrame2{-1, -1, 0, 0, 0, 0, make([]uint64, 0)}
	stack := make([]int, 1)
	stack[0] = HEADER
	bits := bitstream{}
	i := 0
	prevStackDepth := len(stack)
	for len(stack) > 0 {
		if i < len(line) && !(stack[len(stack)-1] == LEN_OP_CLEAN || stack[len(stack)-1] == COUNT_OP_CLEAN) {
			c := line[i]
			i++
			b, err := strconv.ParseUint(string(c), 16, 64)
			util.Check_error(err)
			nibble := bitstream{b, 4}
			bits = addNibble(bits, nibble)
		} else if prevStackDepth == len(stack) {
			panic(3)
		}
		prevStackDepth = len(stack)

		bitsConsumed := 0
		packetsConsumed := 0
		util.Dprintf("0x%x %d %d/%d\n", bits.bits, bits.len, i, len(line))
		switch stack[len(stack)-1] {
		case HEADER:
			version := 0
			typeid := 0
			parsed := false
			if bits, version, typeid, bitsConsumed, parsed = parseHeader(bits); parsed {
				util.Dprintf("V:%d T:%d\n", version, typeid)
				stack = stack[:len(stack)-1]
				switch typeid {
				case 4:
					stack = append(stack, LITERAL)
				default:
					stack = append(stack, typeid)
					stack = append(stack, OPERATOR)
				}
			}
		case LITERAL:
			value := 0
			parsed := false
			if bits, value, bitsConsumed, parsed = parseLiteral(bits); parsed {
				util.Dprintf("Lv: %d\n", value)
				stack = stack[:len(stack)-1]
				packetsConsumed++
				subPacketLenstack[len(subPacketLenstack)-1].values = append(subPacketLenstack[len(subPacketLenstack)-1].values, uint64(value))
			}
		case OPERATOR:
			lty, length := 0, 0
			parsed := false
			if bits, lty, length, bitsConsumed, parsed = parseSubPacketHeader(bits); parsed {
				util.Dprintf("Sp: %d %d\n", lty, length)
				stack = stack[:len(stack)-1]
				op := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if length > 0 {
					switch lty {
					case LENGTH_TYPE:
						stack = append(stack, LEN_OP_CLEAN)
						subPacketLenstack[len(subPacketLenstack)-1].bitUsed += bitsConsumed
						subPacketLenstack = append(subPacketLenstack, subPackFrame2{op, lty, length, 0, 0, 0, make([]uint64, 0)})
						stack = append(stack, HEADER)
					case COUNT_TYPE:
						stack = append(stack, COUNT_OP_CLEAN)
						subPacketLenstack[len(subPacketLenstack)-1].bitUsed += bitsConsumed
						subPacketLenstack = append(subPacketLenstack, subPackFrame2{op, lty, 0, 0, length, 0, make([]uint64, 0)})
						stack = append(stack, HEADER)
					}
					bitsConsumed = 0
				}
			}
		case LEN_OP_CLEAN:
			// For now ignore padding on lengthed packet
			//parsed := false
			//if bits, parsed = trimBits(bits, subPacketLenstack[len(subPacketLenstack)-1].lenEx-subPacketLenstack[len(subPacketLenstack)-1].lenUsed); parsed {
			if subPacketLenstack[len(subPacketLenstack)-1].bitEx > subPacketLenstack[len(subPacketLenstack)-1].bitUsed {
				util.Dprintf("Sp: more (len)\n")
				stack = append(stack, HEADER)
			} else if subPacketLenstack[len(subPacketLenstack)-1].bitEx < subPacketLenstack[len(subPacketLenstack)-1].bitUsed {
				panic(1)
			} else {
				util.Dprintf("Sp: exit (len)\n")

				bitsConsumed = subPacketLenstack[len(subPacketLenstack)-1].bitUsed
				packetsConsumed = 1

				value := performOp(subPacketLenstack[len(subPacketLenstack)-1].op, subPacketLenstack[len(subPacketLenstack)-1].values)

				subPacketLenstack = subPacketLenstack[:len(subPacketLenstack)-1]
				subPacketLenstack[len(subPacketLenstack)-1].values = append(subPacketLenstack[len(subPacketLenstack)-1].values, value)
				stack = stack[:len(stack)-1]
			}
		case COUNT_OP_CLEAN:
			// For now ignore padding on multipacket
			//parsed := false
			//if bits, parsed = trimBits(bits, subPacketLenstack[len(subPacketLenstack)-1].lenUsed); parsed {
			if subPacketLenstack[len(subPacketLenstack)-1].packEx > subPacketLenstack[len(subPacketLenstack)-1].packUsed {
				util.Dprintf("Sp: more (pack)\n")
				stack = append(stack, HEADER)
			} else if subPacketLenstack[len(subPacketLenstack)-1].packEx < subPacketLenstack[len(subPacketLenstack)-1].packUsed {
				panic(2)
			} else {
				util.Dprintf("Sp: exit (pack)\n")

				bitsConsumed = subPacketLenstack[len(subPacketLenstack)-1].bitUsed
				packetsConsumed = 1

				value := performOp(subPacketLenstack[len(subPacketLenstack)-1].op, subPacketLenstack[len(subPacketLenstack)-1].values)

				subPacketLenstack = subPacketLenstack[:len(subPacketLenstack)-1]
				subPacketLenstack[len(subPacketLenstack)-1].values = append(subPacketLenstack[len(subPacketLenstack)-1].values, value)
				stack = stack[:len(stack)-1]
			}
		}

		subPacketLenstack[len(subPacketLenstack)-1].bitUsed += bitsConsumed
		subPacketLenstack[len(subPacketLenstack)-1].packUsed += packetsConsumed

		util.Dprintf("%v %v\n\n", stack, subPacketLenstack)
	}
	return int(subPacketLenstack[len(subPacketLenstack)-1].values[0])
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	file_scanner.Scan()
	line := file_scanner.Text()
	o := parsePacket2(line)

	return fmt.Sprintf("%d", o)
}
