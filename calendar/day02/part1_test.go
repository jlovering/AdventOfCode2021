package adventofcode

import (
	util "adventofcode/util/common"
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_part1_test",
			},
			want: "150",
		},
	}
	util.Setdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.file); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestRunPart1(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AoCInput",
			args: args{
				file: "test/part1_input",
			},
			want: "",
		},
	}
	util.Setdebug(false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.args.file)
			fmt.Printf("%s\n", got)
		})
	}
}
