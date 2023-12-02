package dayone

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

func TestPartOne(t *testing.T) {
	firstInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	var tests = []struct {
		input string
		want  int
	}{
		{firstInput, 142},
		{secondInput, 57346},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			ans := partOne(strings.Split(tt.input, "\n"))
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	firstInput := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	var tests = []struct {
		input string
		want  int
	}{
		{firstInput, 281},
		{secondInput, 57345},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("Test Number %d", i)
		t.Run(name, func(t *testing.T) {
			ans := partTwo(strings.Split(tt.input, "\n"))
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
