package dayten

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

var (
	//go:embed input.txt
	secondInput string
)

func TestSolve(t *testing.T) {
	firstInput := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	var tests = []struct {
		input      string
		wantFirst  int
		wantSecond int
	}{
		{firstInput, 8, 0},
		{secondInput, 6931, 0},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			p1, p2 := solve(strings.Split(tt.input, "\n"))
			if p1 != tt.wantFirst {
				t.Errorf("part one: got %d, want %d", p1, tt.wantFirst)
			}
			if p2 != tt.wantSecond {
				t.Errorf("part two: got %d, want %d", p2, tt.wantSecond)
			}
		})
	}
}
