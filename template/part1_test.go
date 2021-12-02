package AdventOfCode

import (
	"testing"
)

func TestPart1Samples(t *testing.T) {
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
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.file); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}
