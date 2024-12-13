package daythirteen

import (
	"fmt"
	"strings"
	"testing"

	_ "embed"

	"github.com/google/go-cmp/cmp"
)

var (
	//go:embed input.txt
	puzzleInput string
)

func TestSolve(t *testing.T) {
	var tests = []struct {
		input string
		want  solution
	}{
		{`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			solution{1, 0}},
		{puzzleInput, solution{0, 0}},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			got := solve(strings.Split(tt.input, "\n"))
			if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(solution{})); diff != "" {
				t.Errorf("got diff: %s", diff)
			}
		})
	}
}
