package AdventOfCode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func oxygenGenRate(all_values []int, bitlen int) int {
	for i := 0; i < bitlen; i++ {
		total_values := len(all_values)
		bitcount := make(map[int]int, bitlen)
		for j := 0; j < bitlen; j++ {
			bitcount[j] = 0
		}
		for _, value := range all_values {
			for j := 0; j < bitlen; j++ {
				if (value>>j)&0x1 == 0x1 {
					bitcount[bitlen-1-j]++
				}
			}
		}
		dprintf("%d\n%v\n%v\n%d\n", i, all_values, bitcount, total_values)

		match := 0
		mask := 0x1 << (bitlen - 1 - i)
		if bitcount[i] > (total_values - bitcount[i]) {
			match = 0x1 << (bitlen - 1 - i)
		} else if bitcount[i] == (total_values - bitcount[i]) {
			match = 0x1 << (bitlen - 1 - i)
		}
		dprintf("\t%d %d %02x\n\n", bitlen, i, match)

		new_values := make([]int, 0, len(all_values))
		for _, v := range all_values {
			if (v & mask) == match {
				new_values = append(new_values, v)
			}
		}
		all_values = new_values
		if len(all_values) == 1 {
			break
		}
	}

	return all_values[0]
}

func co2Rate(all_values []int, bitlen int) int {
	for i := 0; i < bitlen; i++ {
		total_values := len(all_values)
		bitcount := make(map[int]int, bitlen)
		for j := 0; j < bitlen; j++ {
			bitcount[j] = 0
		}
		for _, value := range all_values {
			for j := 0; j < bitlen; j++ {
				if (value>>j)&0x1 == 0x1 {
					bitcount[bitlen-1-j]++
				}
			}
		}
		dprintf("%d\n%v\n%v\n%d\n", i, all_values, bitcount, total_values)

		match := 0
		mask := 0x1 << (bitlen - 1 - i)
		if bitcount[i] < (total_values - bitcount[i]) {
			match = 0x1 << (bitlen - 1 - i)
		}

		dprintf("\t%d %d %02x\n\n", bitlen, i, match)

		new_values := make([]int, 0, len(all_values))
		for _, v := range all_values {
			if (v & mask) == match {
				new_values = append(new_values, v)
			}
		}
		all_values = new_values
		if len(all_values) == 1 {
			break
		}
	}

	return all_values[0]
}

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer sdout_writer.Flush()

	f, err := os.Open(filename)
	check_error(err)
	defer f.Close()

	file_stat, err := f.Stat()
	check_error(err)
	file_size := file_stat.Size()

	file_scanner := bufio.NewScanner(f)

	all_values := make([]int, 0, file_size)

	bitlen := 0
	for file_scanner.Scan() {
		line := file_scanner.Text()
		var value string
		fmt.Sscanf(line, "%s", &value)
		bitlen = len(value)
		value_int, err := strconv.ParseUint(value, 2, 64)
		check_error(err)
		all_values = append(all_values, int(value_int))
	}

	oxygen_values := all_values
	co2_values := all_values

	oxygen_rate := oxygenGenRate(oxygen_values, bitlen)
	co2_rate := co2Rate(co2_values, bitlen)

	return fmt.Sprintf("%d", oxygen_rate*co2_rate)
}
