package AdventOfCode

import (
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
				file: "test/sample_part2_test",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.file); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}
