package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
)

func Part2(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)
	enhancer, inImage := parseInput(file_scanner)

	inImage = edgeExpand(inImage, 3, '.')
	image := enhance(inImage, enhancer)
	printImage(image)

	for n := 0; n < 49; n++ {
		image = edgeExpand(image, 2, image[0][0])
		image = enhance(image, enhancer)
		printImage(image)
	}
	return fmt.Sprintf("%d", countPixels(image))
}
