package dayfive

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
		{input: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
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
