package dayfive

import (
	"strconv"
	"strings"
	"unicode"
)

type instructions struct {
	sr, dr, rl int64
}

type seedRange struct {
	start  int64
	length int64
}

type stuff struct {
	seeds  []int64
	seeds2 []seedRange
	instr  [][]instructions
}

func parse(input []string) stuff {
	s := stuff{[]int64{}, []seedRange{}, make([][]instructions, 7)}
	seeds := strings.Split(strings.TrimPrefix(input[0], "seeds: "), " ")
	for _, seed := range seeds {
		num, _ := strconv.ParseInt(seed, 10, 64)
		s.seeds = append(s.seeds, num)
	}
	for i := 0; i < len(seeds); i += 2 {
		num, _ := strconv.ParseInt(seeds[i], 10, 64)
		num2, _ := strconv.ParseInt(seeds[i+1], 10, 64)
		s.seeds2 = append(s.seeds2, seedRange{num, num2})
	}
	pos := -1
	for _, line := range input[1:] {
		if strings.HasSuffix(line, ":") {
			pos++
		}
		if line == "" {
			continue
		}
		if unicode.IsNumber(rune(line[0])) {
			a := strings.Split(line, " ")
			dr, _ := strconv.ParseInt(a[0], 10, 64)
			sr, _ := strconv.ParseInt(a[1], 10, 64)
			rl, _ := strconv.ParseInt(a[2], 10, 64)
			s.instr[pos] = append(s.instr[pos], instructions{sr, dr, rl})
		}
	}
	return s
}

func checkMap(n int64, instr []instructions) int64 {
	for _, i := range instr {
		if n >= i.sr && n <= i.sr+i.rl {
			return n - i.sr + i.dr
		}
	}
	return n
}

func explore(s stuff) (int64, int64) {
	partOne, partTwo := int64(9223372036854775807), int64(9223372036854775807)
	for _, v := range s.seeds {
		n := v
		for j := 0; j < 7; j++ {
			n = checkMap(n, s.instr[j])
		}
		if n < partOne {
			partOne = n
		}
	}

	// I tried to do some smart range based stuff, but using the original
	// solution got me an answer within 3 minutes, it was off by one for some reason though..
	for _, v := range s.seeds2 {
		for i := int64(0); i < v.length; i++ {
			n := v.start + i
			for j := 0; j < 7; j++ {
				n = checkMap(n, s.instr[j])
			}
			if n < partTwo {
				partTwo = n
			}
		}
	}

	return partOne, partTwo
}

func solve(input []string) (int, int) {
	partOne := int64(0)
	partTwo := int64(0)
	s := parse(input)
	partOne, partTwo = explore(s)
	return int(partOne), int(partTwo)
}
