package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type bitstream struct {
	bits uint64
	len  uint64
}

func addNibble(bits bitstream, nibble bitstream) bitstream {
	//util.Dprintf("%X %X %d %d\n", bits.bits, nibble.bits, bits.len, nibble.len)
	return bitstream{(bits.bits << nibble.len) | nibble.bits, bits.len + nibble.len}
}

func trimBits(bits bitstream, count uint64) (bitstream, bool) {
	if bits.len < count {
		return bits, false
	}
	nB := bits.bits & (0x1 << (bits.len - count))
	return bitstream{nB, bits.len - count}, true
}

func parseLiteral(bits bitstream) (bitstream, int, int, bool) {
	org := bits
	out := 0
	cont := 1
	bC := 0
	//util.Dprintf("PL 0x%x %d\n", bits.bits, bits.len)
	for cont > 0 {
		if bits.len < 5 {
			return org, 0, 0, false
		}
		out <<= 4
		chunk := bits.bits & (0x1F << (bits.len - 5)) >> (bits.len - 5)
		cont = int(chunk & 0x10)
		out |= int(chunk & 0xF)
		bits.bits &= (0x1 << (bits.len - 5)) - 1
		bits.len -= 5
		bC += 5
		//util.Dprintf("\tv: %x\n", out)
	}
	return bits, out, bC, true
}

func parseHeader(bits bitstream) (bitstream, int, int, int, bool) {
	if bits.len < 6 {
		return bits, 0, 0, 0, false
	}
	version := bits.bits & (0x7 << (bits.len - 3)) >> (bits.len - 3)
	typeid := bits.bits & (0x7 << (bits.len - 6)) >> (bits.len - 6)

	nBits := bits.bits & (0x1<<(bits.len-6) - 1)
	return bitstream{nBits, bits.len - 6}, int(version), int(typeid), 6, true
}

func parseSubPacketHeader(bits bitstream) (bitstream, int, int, int, bool) {
	lty := (bits.bits >> (bits.len - 1)) & 0x1

	headerLen := 0
	switch lty {
	case 0:
		headerLen = 16
	case 1:
		headerLen = 12
	}
	if bits.len < uint64(headerLen) {
		return bits, 0, 0, 0, false
	}

	shift := bits.len - uint64(headerLen)
	mask := uint64((0x1 << (headerLen - 1)) - 1)
	length := int((bits.bits >> shift) & mask)
	nBits := bitstream{}
	nBits.bits = bits.bits & ((0x1 << (bits.len - uint64(headerLen))) - 1)
	nBits.len = bits.len - uint64(headerLen)
	return nBits, int(lty), length, headerLen, true
}

type subPackFrame struct {
	lty      int
	bitEx    int
	bitUsed  int
	packEx   int
	packUsed int
}

func parsePacket(line string) int {
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

	subPacketLenstack := make([]subPackFrame, 1)
	subPacketLenstack[0] = subPackFrame{-1, 0, 0, 0, 0}
	stack := make([]int, 1)
	stack[0] = HEADER
	bits := bitstream{}
	versionSum := 0
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
				versionSum += version
				stack = stack[:len(stack)-1]
				switch typeid {
				case 4:
					stack = append(stack, LITERAL)
				default:
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
			}
		case OPERATOR:
			lty, length := 0, 0
			parsed := false
			if bits, lty, length, bitsConsumed, parsed = parseSubPacketHeader(bits); parsed {
				util.Dprintf("Sp: %d %d\n", lty, length)
				stack = stack[:len(stack)-1]
				if length > 0 {
					switch lty {
					case LENGTH_TYPE:
						stack = append(stack, LEN_OP_CLEAN)
						subPacketLenstack[len(subPacketLenstack)-1].bitUsed += bitsConsumed
						subPacketLenstack = append(subPacketLenstack, subPackFrame{lty, length, 0, 0, 0})
						stack = append(stack, HEADER)
					case COUNT_TYPE:
						stack = append(stack, COUNT_OP_CLEAN)
						subPacketLenstack[len(subPacketLenstack)-1].bitUsed += bitsConsumed
						subPacketLenstack = append(subPacketLenstack, subPackFrame{lty, 0, 0, length, 0})
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

				subPacketLenstack = subPacketLenstack[:len(subPacketLenstack)-1]
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

				subPacketLenstack = subPacketLenstack[:len(subPacketLenstack)-1]
				stack = stack[:len(stack)-1]
			}
		}

		subPacketLenstack[len(subPacketLenstack)-1].bitUsed += bitsConsumed
		subPacketLenstack[len(subPacketLenstack)-1].packUsed += packetsConsumed

		util.Dprintf("%v %v\n\n", stack, subPacketLenstack)
	}
	return versionSum
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	file_scanner.Scan()
	line := file_scanner.Text()
	t := parsePacket(line)

	return fmt.Sprintf("%d", t)
}
