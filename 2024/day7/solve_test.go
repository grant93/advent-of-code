package dayseven

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
		{`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			solution{3749, 11387}},
		{puzzleInput, solution{663613490587, 110365987435001}},
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
