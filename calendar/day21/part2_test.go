package adventofcode

import (
	util "adventofcode/util/common"
	"fmt"
	"testing"
)

func TestPart2Samples(t *testing.T) {
	type args struct {
		p1start int
		p2start int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AoCExample",
			args: args{
				p1start: 4,
				p2start: 8,
			},
			want: "444356092776315",
		},
	}
	util.Setmasterdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.p1start, tt.args.p2start); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestRunPart2(t *testing.T) {
	type args struct {
		p1start int
		p2start int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AoCInput",
			args: args{
				p1start: 2,
				p2start: 10,
			},
			want: "",
		},
	}
	util.Setmasterdebug(false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.args.p1start, tt.args.p2start)
			fmt.Printf("%s\n", got)
		})
	}
}
