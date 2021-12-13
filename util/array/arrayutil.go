package arrayutil

import (
	"fmt"
)

type SliceValue interface{}

func SliceBuilder2D(i, j int) [][]SliceValue {
	var slice [][]SliceValue = make([][]SliceValue, i)
	for x := range slice {
		slice[x] = make([]SliceValue, j)
	}
	return slice
}

func SliceBuilder2DString(i, j int) [][]string {
	var slice [][]string = make([][]string, i)
	for x := range slice {
		slice[x] = make([]string, j)
	}
	return slice
}

func SliceBuilder2DInt(i, j int) [][]int {
	var slice [][]int = make([][]int, i)
	for x := range slice {
		slice[x] = make([]int, j)
	}
	return slice
}

func SliceBuilder2DRune(i, j int) [][]rune {
	var slice [][]rune = make([][]rune, i)
	for x := range slice {
		slice[x] = make([]rune, j)
	}
	return slice
}

func SliceBuilder2DBool(i, j int) [][]bool {
	var slice [][]bool = make([][]bool, i)
	for x := range slice {
		slice[x] = make([]bool, j)
	}
	return slice
}

func SPrintArrayYX(array interface{}, format string) string {
	out := ""
	switch arr := array.(type) {
	case [][]string:
		for j := range arr {
			for i := range arr[0] {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	case [][]int:
		for j := range arr {
			for i := range arr[0] {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	case [][]rune:
		for j := range arr {
			for i := range arr[0] {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	case [][]bool:
		for j := range arr {
			for i := range arr[0] {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	}
	return out
}

func SPrintArrayXY(array interface{}, format string) string {
	out := ""
	switch arr := array.(type) {
	case [][]string:
		for i := range arr[0] {
			for j := range arr {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	case [][]int:
		for i := range arr[0] {
			for j := range arr {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	case [][]rune:
		for i := range arr[0] {
			for j := range arr {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	case [][]bool:
		for i := range arr[0] {
			for j := range arr {
				out += fmt.Sprintf(format+" ", arr[j][i])
			}
			out += "\n"
		}
	}
	return out
}
