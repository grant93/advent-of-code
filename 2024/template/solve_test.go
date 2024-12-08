package dayone

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
		{`3   4
4   3
2   5
1   3
3   9
3   3`,
			solution{0, 0}},
		{puzzleInput, solution{0, 0}},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			got := solve(strings.Split(tt.input, "\n"))
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got diff: %s", diff)
			}
		})
	}
}
