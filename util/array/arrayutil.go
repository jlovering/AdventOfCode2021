package arrayutil

import (
	"fmt"
	"strconv"
	"strings"
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

func Dynamic2DIntSliceBuilder(
	input string,
	line_delimiter string,
	cell_delimiter string,
) [][]int {
	lines := strings.Split(input, line_delimiter)
	grid := make([][]int, 0)
	for i, line := range lines {
		grid = append(grid, make([]int, 0))
		cells := strings.Split(line, cell_delimiter)
		for _, cell := range cells {
			v, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			grid[i] = append(grid[i], v)
		}
	}
	return grid
}

func SPrintArrayYX(array interface{}, format string) string {
	var sb strings.Builder
	switch arr := array.(type) {
	case [][]string:
		for j := range arr {
			for i := range arr[0] {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	case [][]int:
		for j := range arr {
			for i := range arr[0] {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	case [][]rune:
		for j := range arr {
			for i := range arr[0] {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	case [][]bool:
		for j := range arr {
			for i := range arr[0] {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func SPrintArrayXY(array interface{}, format string) string {
	var sb strings.Builder
	switch arr := array.(type) {
	case [][]string:
		for i := range arr[0] {
			for j := range arr {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	case [][]int:
		for i := range arr[0] {
			for j := range arr {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	case [][]rune:
		for i := range arr[0] {
			for j := range arr {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	case [][]bool:
		for i := range arr[0] {
			for j := range arr {
				sb.WriteString(fmt.Sprintf(format+" ", arr[j][i]))
			}
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}
