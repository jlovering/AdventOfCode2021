package main

import (
	arrayutil "adventofcode/util/array"
	"fmt"
)

func main() {
	testSlice := arrayutil.SliceBuilder2D(10, 10)

	testSlice[0][0] = 'a'

	switch testSlice[0][0].(type) {
	case rune:
		fmt.Printf("%c\n", testSlice[0][0])
	default:
		panic(1)
	}

	testSlice[0][1] = 0

	switch testSlice[0][1].(type) {
	case int:
		fmt.Printf("%d\n", testSlice[0][1])
	default:
		panic(1)
	}

	fmt.Println(arrayutil.SPrintArrayXY(testSlice, "%5v"))
	fmt.Println(arrayutil.SPrintArrayYX(testSlice, "%5v"))

	testSlice2 := arrayutil.SliceBuilder2DString(2, 2)

	testSlice2[0][0] = "a"
	testSlice2[0][1] = "b"
	testSlice2[1][0] = "c"
	testSlice2[1][1] = "d"

	fmt.Println(arrayutil.SPrintArrayXY(testSlice2, "%2s"))
	fmt.Println(arrayutil.SPrintArrayYX(testSlice2, "%2s"))

	testSlice3 := arrayutil.SliceBuilder2DRune(2, 2)

	testSlice3[0][0] = 'a'
	testSlice3[0][1] = 'b'
	testSlice3[1][0] = 'c'
	testSlice3[1][1] = 'd'

	fmt.Println(arrayutil.SPrintArrayXY(testSlice3, "%c"))
	fmt.Println(arrayutil.SPrintArrayYX(testSlice3, "%c"))
}
