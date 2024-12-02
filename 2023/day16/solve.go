package daysixteen

type Direction int

type coordinates struct {
	x, y int
}

func (c coordinates) Add(b coordinates) coordinates {
	return coordinates{c.x + b.x, c.y + b.y}
}

const (
	DIRECTION_LEFT Direction = iota
	DIRECTION_UP
	DIRECTION_RIGHT
	DIRECTION_DOWN
)

var (
	movementMap = map[byte]map[Direction][]Direction{
		'|': {
			DIRECTION_LEFT:  {DIRECTION_UP, DIRECTION_DOWN},
			DIRECTION_RIGHT: {DIRECTION_UP, DIRECTION_DOWN},
			DIRECTION_UP:    {DIRECTION_UP},
			DIRECTION_DOWN:  {DIRECTION_DOWN},
		},
		'-': {
			DIRECTION_LEFT:  {DIRECTION_LEFT},
			DIRECTION_RIGHT: {DIRECTION_RIGHT},
			DIRECTION_UP:    {DIRECTION_RIGHT, DIRECTION_LEFT},
			DIRECTION_DOWN:  {DIRECTION_RIGHT, DIRECTION_LEFT},
		},
		'\\': {
			DIRECTION_LEFT:  {DIRECTION_UP},
			DIRECTION_RIGHT: {DIRECTION_DOWN},
			DIRECTION_UP:    {DIRECTION_LEFT},
			DIRECTION_DOWN:  {DIRECTION_RIGHT},
		},
		'/': {
			DIRECTION_LEFT:  {DIRECTION_DOWN},
			DIRECTION_RIGHT: {DIRECTION_UP},
			DIRECTION_UP:    {DIRECTION_RIGHT},
			DIRECTION_DOWN:  {DIRECTION_LEFT},
		},
		'.': {
			DIRECTION_LEFT:  {DIRECTION_LEFT},
			DIRECTION_RIGHT: {DIRECTION_RIGHT},
			DIRECTION_UP:    {DIRECTION_UP},
			DIRECTION_DOWN:  {DIRECTION_DOWN},
		},
	}

	dirChange = map[Direction]coordinates{
		DIRECTION_LEFT:  {-1, 0},
		DIRECTION_UP:    {0, -1},
		DIRECTION_RIGHT: {1, 0},
		DIRECTION_DOWN:  {0, 1},
	}
)

func run(input []string, pos coordinates, dir Direction, m *map[coordinates]map[Direction]bool) {
	if pos.y < 0 || pos.y >= len(input) || pos.x < 0 || pos.x >= len(input[0]) {
		return
	}
	// Cache the position and the direction so we don't re-run, if we run all the branches we blow up the stack.
	if a, has := (*m)[pos]; has {
		if a[dir] {
			// We've been here before.
			return
		} else {
			a[dir] = true
		}
	} else {
		(*m)[pos] = map[Direction]bool{}
	}
	tile := input[pos.y][pos.x]
	next := movementMap[tile][dir]
	for _, d := range next {
		newPos := pos.Add(dirChange[d])
		run(input, newPos, d, m)
	}
}

func solvePartTwo(input []string) int {
	highest := 0
	xMax := len(input[0]) - 1
	yMax := len(input) - 1
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if x != 0 && x != xMax && y != 0 && y != yMax {
				continue
			}
			if x == 0 {
				m := make(map[coordinates]map[Direction]bool)
				run(input, coordinates{x, y}, DIRECTION_RIGHT, &m)
				if s := len(m); s > highest {
					highest = s
				}
			}
			if x == xMax {
				m := make(map[coordinates]map[Direction]bool)
				run(input, coordinates{x, y}, DIRECTION_LEFT, &m)
				if s := len(m); s > highest {
					highest = s
				}
			}
			if y == 0 {
				m := make(map[coordinates]map[Direction]bool)
				run(input, coordinates{x, y}, DIRECTION_DOWN, &m)
				if s := len(m); s > highest {
					highest = s
				}
			}
			if x == 0 {
				m := make(map[coordinates]map[Direction]bool)
				run(input, coordinates{x, y}, DIRECTION_UP, &m)
				if s := len(m); s > highest {
					highest = s
				}
			}
		}
	}
	return highest
}

func solve(input []string) (int, int) {
	energisedMap := map[coordinates]map[Direction]bool{}
	run(input, coordinates{0, 0}, DIRECTION_RIGHT, &energisedMap)
	partTwo := solvePartTwo(input)
	return len(energisedMap), partTwo
}
