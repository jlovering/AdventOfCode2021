package adventofcode

import (
	util "adventofcode/util/common"
	"fmt"
	"testing"
)

func TestPart2Samples(t *testing.T) {
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
			want: "46",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test0.2",
			},
			want: "38",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test1",
			},
			want: "39",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test2",
			},
			want: "39769202357779",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/sample_test3",
			},
			want: "2758514936282235",
		},
	}
	util.Setmasterdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.file); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestRunPart2(t *testing.T) {
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
	util.Setmasterdebug(false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.args.file)
			fmt.Printf("%s\n", got)
		})
	}
}
