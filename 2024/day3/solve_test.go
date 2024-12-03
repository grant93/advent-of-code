package daythree

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
		{input: `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			want: 161, want2: 161},
		{input: `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			want: 161, want2: 48},
		{puzzleInput, 183669043, 59097164},
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
