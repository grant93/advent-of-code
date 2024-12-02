package dayone

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
		want  int64
		want2 int64
	}{
		{input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			want: 11, want2: 31},
		{puzzleInput, 1834060, 21607792},
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
