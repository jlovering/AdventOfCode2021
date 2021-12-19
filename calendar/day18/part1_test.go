package adventofcode

import (
	util "adventofcode/util/common"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		snail string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "SnailExample",
			args: args{
				snail: "[1,2]",
			},
			want: "[1,2]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[1,2],3]",
			},
			want: "[[1,2],3]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[9,[8,7]]",
			},
			want: "[9,[8,7]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[1,9],[8,5]]",
			},
			want: "[[1,9],[8,5]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
			},
			want: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
			},
			want: "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
			},
			want: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
		},
	}
	util.Setmasterdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer util.SdoutFlush()
			if got := printSnail(snailParse(tt.args.snail)); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		snail string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[[9,8],1],2],3],4]",
			},
			want: "[[[[0,9],2],3],4]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[7,[6,[5,[4,[3,2]]]]]",
			},
			want: "[7,[6,[5,[7,0]]]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[6,[5,[4,[3,2]]]],1]",
			},
			want: "[[6,[5,[7,0]]],3]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			},
			want: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			},
			want: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			},
			want: "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
			},
			want: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			},
			want: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		//{
		//	name: "SnailExample",
		//	args: args{
		//		snail: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		//	},
		//	want: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		//},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			},
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[[1,2],[3,4]],[5,6]],[7,8]],[9,1]]",
			},
			want: "[[[[0,[5,4]],[5,6]],[7,8]],[9,1]]",
		},
	}
	util.Setmasterdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer util.SdoutFlush()
			parsed := snailParse(tt.args.snail)
			reduced, _ := snailReduce(parsed)
			if got := printSnail(reduced); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestAddAll(t *testing.T) {
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
				file: "test/addition_test1",
			},
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_test2",
			},
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_test3",
			},
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_test3.5",
			},
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.1",
			},
			want: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.2",
			},
			want: "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.3",
			},
			want: "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.4",
			},
			want: "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.5",
			},
			want: "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.6",
			},
			want: "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.7",
			},
			want: "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.8",
			},
			want: "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_testpre4.9",
			},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_test4",
			},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			name: "AoCExample",
			args: args{
				file: "test/addition_test5",
			},
			want: "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
	}
	util.Setmasterdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer util.SdoutFlush()
			allParsed := parseInputForTest(tt.args.file)
			added := addAll(allParsed)
			if got := printSnail(added); got != tt.want {
				t.Errorf("\"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	type args struct {
		snail string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "SnailExample",
			args: args{
				snail: "[[1,2],[[3,4],5]]",
			},
			want: 143,
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			},
			want: 1384,
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			},
			want: 445,
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			},
			want: 791,
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			},
			want: 1137,
		},
		{
			name: "SnailExample",
			args: args{
				snail: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			},
			want: 3488,
		},
	}
	util.Setmasterdebug(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer util.SdoutFlush()
			parsed := snailParse(tt.args.snail)
			reduced, _ := snailReduce(parsed)
			if got := snailMagnitude(reduced); got != tt.want {
				t.Errorf("\"%d\", want \"%d\"", got, tt.want)
			}
		})
	}
}

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
				file: "test/sample_test",
			},
			want: "4140",
		},
	}
	util.Setmasterdebug(true)
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
	util.Setmasterdebug(false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.args.file)
			fmt.Printf("%s\n", got)
		})
	}
}
