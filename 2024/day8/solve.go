package dayeight

import "fmt"

type solution struct {
	partOne, partTwo int64
}

type coord struct {
	x, y int
}

type antennaMap struct {
	antennas   map[rune][]coord
	xMax, yMax int
}

func (c *coord) Add(a coord) coord {
	return coord{c.x + a.x, c.y + a.y}
}

func (c *coord) Subtract(a coord) coord {
	return coord{c.x - a.x, c.y - a.y}
}

func isAlphaNumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func parse(input []string) antennaMap {
	a := antennaMap{map[rune][]coord{}, len(input[0]), len(input)}
	for y, line := range input {
		for x, t := range line {
			if isAlphaNumeric(t) {
				c := coord{x, y}
				if _, ok := a.antennas[t]; ok {
					a.antennas[t] = append(a.antennas[t], c)
					continue
				}
				a.antennas[t] = []coord{{x, y}}
			}
		}
	}
	return a
}

func generatePairs(n []coord) [][]coord {
	ret := [][]coord{}
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			ret = append(ret, []coord{n[i], n[j]})
		}
	}
	return ret
}

// Only need to add the gradient to the first point, the other direction will get handled in the
// opposite permutation (a,b) and (b,a).
func generateAntinode(am antennaMap, a, gradient coord, partTwo bool) []coord {
	ret := []coord{}
	next := a.Add(gradient)
	if next.x >= 0 && next.x < am.xMax && next.y >= 0 && next.y < am.yMax {
		ret = append(ret, next)
		if partTwo {
			ret = append(ret, generateAntinode(am, next, gradient, partTwo)...)
		}
		return ret
	} else {
		return ret
	}
}

// Loop through all antenna coords and deduce antinodes.
func findAntinodes(a antennaMap, partTwo bool) int64 {
	z := map[coord]bool{}
	for _, v := range a.antennas {
		perms := generatePairs(v)
		for _, p := range perms {
			gradient := coord{p[0].x - p[1].x, p[0].y - p[1].y}
			nodes := generateAntinode(a, p[0], gradient, partTwo)
			// All antenna points are antinodes in part 2, shove them in.
			if partTwo {
				z[p[0]] = true
			}
			for _, vv := range nodes {
				z[vv] = true
			}
		}
	}
	return int64(len(z))
}

// Dumb helper to print the map because I had a bug.
func printMap(a antennaMap, b map[coord]bool) {
	for i := range a.yMax {
		line := ""
		for j := range a.xMax {
			if b[coord{j, i}] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func solve(input []string) solution {
	n := parse(input)
	return solution{findAntinodes(n, false), findAntinodes(n, true)}
}
