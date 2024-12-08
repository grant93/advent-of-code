package dayeight

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
		{`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
			solution{14, 34}},
		{puzzleInput, solution{426, 1359}},
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
