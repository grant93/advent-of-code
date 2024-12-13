package daythirteen

import (
	"fmt"
	"strconv"
	"strings"
)

type solution struct {
	partOne, partTwo int64
}

type equation struct {
	ax, ay, bx, by, px, py int64
}

func parseNums(line string) (int64, int64) {
	tmp := strings.Split(line, ": ")
	tmp = strings.Split(tmp[1], ", ")
	x, _ := strconv.ParseInt(tmp[0][2:], 10, 64)
	y, _ := strconv.ParseInt(tmp[1][2:], 10, 64)
	return x, y
}

func parseEquation(input []string, partTwo bool) []equation {
	i := 0
	ans := []equation{}
	for {
		e := equation{}
		if i >= len(input) {
			break
		}
		e.ax, e.ay = parseNums(input[i])
		e.bx, e.by = parseNums(input[i+1])
		e.px, e.py = parseNums(input[i+2])
		if partTwo {
			e.px += 10000000000000
			e.py += 10000000000000
		}
		i += 4
		ans = append(ans, e)
	}
	fmt.Println(ans)
	return ans
}

func solveEquation(e equation) int64 {
	// who would have thought maths would come in handy.
	denominator := (e.ax * e.by) - (e.bx * e.ay)
	numeratorA := (e.px * e.by) - (e.bx * e.py)
	numeratorB := (e.ax * e.py) - (e.px * e.ay)
	if numeratorA%denominator != 0 || numeratorB%denominator != 0 {
		return 0
	}
	return ((numeratorA * 3) + numeratorB) / denominator
}

func solve(input []string) solution {
	s := solution{}
	for _, e := range parseEquation(input, false) {
		s.partOne += solveEquation(e)
	}
	for _, e := range parseEquation(input, true) {
		s.partTwo += solveEquation(e)
	}
	return s
}
