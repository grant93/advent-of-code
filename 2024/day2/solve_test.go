package daytwo

import (
	"fmt"
	"strings"
	"testing"

	_ "embed"
)

var (
	//go:embed input.txt
	puzzleInput string
)

func TestSolve(t *testing.T) {
	var tests = []struct {
		input string
		want  int
		want2 int
	}{
		{input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want: 2, want2: 4},
		{puzzleInput, 585, 626},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			p1, p2 := solve(strings.Split(tt.input, "\n"))
			if p1 != tt.want {
				t.Errorf("got %d, want %d", p1, tt.want)
			}
			if p2 != tt.want2 {
				t.Errorf("got %d, want %d", p2, tt.want2)
			}
		})
	}
}
