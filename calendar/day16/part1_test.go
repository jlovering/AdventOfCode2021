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
				file: "test/sample_test0.1",
			},
			want: "6",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test1",
			},
			want: "16",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test2",
			},
			want: "12",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test3",
			},
			want: "23",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test4",
			},
			want: "31",
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
				file: "test/input",
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
