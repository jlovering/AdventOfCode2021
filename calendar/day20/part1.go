package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func printImage(image [][]rune) {
	db := util.Getdebug()
	util.Setdebug(true)
	defer util.Setdebug(db)
	for _, l := range image {
		for i := range l {
			util.Dprintf("%c", l[i])
		}
		util.Dprintf("\n")
	}
}

func printKernel(image [3][3]rune) {
	for _, l := range image {
		for i := range l {
			util.Dprintf("%c", l[i])
		}
		util.Dprintf("\n")
	}
}

func parseInput(file_scanner *bufio.Scanner) (string, [][]rune) {
	file_scanner.Scan()
	enhancer := file_scanner.Text()

	file_scanner.Text()
	file_scanner.Scan()

	inImage := make([][]rune, 0)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		//util.Dprintf("%v\n", inImage)
		inImage = append(inImage, make([]rune, len(line)))
		for i, c := range line {
			inImage[len(inImage)-1][i] = c
		}
	}

	return enhancer, inImage
}

func computeKernel(kernel [3][3]rune) int {
	//printKernel(kernel)
	bitmask := 0
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			if kernel[j][i] == '#' {
				bitmask |= 0x1
			}
			bitmask <<= 1
		}
	}
	bitmask >>= 1
	//util.Dprintf(" %d\n", bitmask)
	return bitmask
}

func constructKernel(inImage [][]rune, i, j int) [3][3]rune {
	kernel := [3][3]rune{}

	//util.Dprintf("%d x %d, %d, %d\n", len(inImage[0]), len(inImage), i, j)
	for kj, dy := range []int{-1, 0, 1} {
		for ki, dx := range []int{-1, 0, 1} {
			x := i + dx
			y := j + dy
			util.IncreasedebugIndent()
			if x < 0 || x >= len(inImage[0]) || y < 0 || y >= len(inImage) {
				//util.Dprintf("Out: x: %2d y: %2d ki: %2d kj: %2d\n", x, y, ki, kj)
				panic("Out of bounds")
			} else {
				//util.Dprintf(" In: x: %2d y: %2d ki: %2d kj: %2d\n", x, y, ki, kj)
				kernel[kj][ki] = inImage[y][x]
			}
			util.DecreasedebugIndent()
		}
	}

	return kernel
}

func edgeExpand(inImage [][]rune, exp int, fill rune) [][]rune {
	outImage := make([][]rune, len(inImage)+exp*2)
	for i := 0; i < len(outImage); i++ {
		outImage[i] = make([]rune, len(inImage[0])+exp*2)
	}

	for j := 0; j < len(outImage); j++ {
		for i := 0; i < len(outImage[0]); i++ {
			in_i := i - exp
			in_j := j - exp
			if in_i < 0 || in_i >= len(inImage[0]) || in_j < 0 || in_j >= len(inImage) {
				outImage[j][i] = fill
			} else {
				outImage[j][i] = inImage[in_j][in_i]
			}
		}
	}

	return outImage
}

func enhance(inImage [][]rune, enhancer string) [][]rune {
	outImage := make([][]rune, len(inImage))
	for i := 0; i < len(outImage); i++ {
		outImage[i] = make([]rune, len(inImage[0]))
	}

	for j := 1; j < len(outImage)-1; j++ {
		for i := 1; i < len(outImage[j])-1; i++ {
			kernel := constructKernel(inImage, i, j)
			idx := computeKernel(kernel)
			outImage[j][i] = rune(enhancer[idx])
		}
		outImage[j] = outImage[j][1 : len(outImage[j])-1]
	}
	//util.Dprintf("%v\n", outImage)
	//printImage(outImage)
	return outImage[1 : len(outImage)-1]
}

func countPixels(image [][]rune) int {
	count := 0
	for _, l := range image {
		for i := range l {
			if l[i] == '#' {
				count++
			}
		}
	}
	return count
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	enhancer, inImage := parseInput(file_scanner)

	inImage = edgeExpand(inImage, 3, '.')
	printImage(inImage)

	out1 := enhance(inImage, enhancer)
	printImage(out1)

	out1 = edgeExpand(out1, 2, out1[0][0])
	printImage(out1)

	out2 := enhance(out1, enhancer)
	printImage(out2)

	return fmt.Sprintf("%d", countPixels(out2))
}
