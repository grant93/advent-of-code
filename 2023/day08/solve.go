package daysix

import (
	"regexp"
	"strings"
)

type options struct {
	left string
	right string
}

var r = regexp.MustCompile(`\w{3}`)

func parse(input []string) (string, map[string]options) {
	instructions := input[0]
	m := map[string]options{}
	for _, n := range input[2:] {
		if n == "" {
			break
		}
		l := r.FindAllString(n, -1)
		m[l[0]] = options{l[1], l[2]}
	}
	return instructions, m
}

func run(i string, m map[string]options) int {
	pos := "AAA"
	n := 0
	for pos != "ZZZ" {
		switch(i[n%len(i)]) {
		case 'L':
			pos = m[pos].left
		case 'R':
			pos = m[pos].right
		}
		n++
	}
	return n
}


// TODO: shove this in a library.
// Implementation from https://go.dev/play/p/undVrRVOgby
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func runTwo(i string, m map[string]options) int {
	nodes := []string{}
	for k := range m {
		if strings.HasSuffix(k, "A") {
			nodes = append(nodes, k)
		}
	}
	n := 0
	counts := make([]int, len(nodes))
	found := 0
	done := false
	for !done {
		newNodes := []string{}
		for z, pos := range nodes {
			switch(i[n%len(i)]) {
			case 'L':
				pos = m[pos].left
			case 'R':
				pos = m[pos].right
			}
			if counts[z] == 0 && strings.HasSuffix(pos, "Z") {
				counts[z] = n+1
				found ++
			}
			newNodes = append(newNodes, pos)
		}
		if found == len(nodes) {
			break
		}
		n++
		nodes = newNodes
	}
	return LCM(counts[0], counts[1], counts[2:]...)
}

func solve(input []string) (int, int) {
	partOne := 0
	partTwo := 0
	i, m := parse(input)
	partOne = run(i, m)
	partTwo = runTwo(i, m)
	return partOne, partTwo
}
