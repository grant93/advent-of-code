package dayseven

import (
	"strconv"
	"strings"
)

type solution struct {
	partOne, partTwo int64
}

type equation struct {
	result  int64
	numbers [][]int64
}

// Evaluate all possibilities left to right, excluding anything that goes over the target number.
func evaluate(a, b []int64, result int64, partTwo bool) []int64 {
	ret := []int64{}
	for _, i := range a {
		if b == nil {
			if i == result {
				ret = append(ret, i)
				break
			}
		}
		for _, j := range b {
			if z := i * j; z <= result {
				ret = append(ret, z)
			}
			if z := i + j; z <= result {
				ret = append(ret, z)
			}
			if partTwo {
				pow := int64(1)
				for range len(strconv.FormatInt(j, 10)) {
					pow *= 10
				}
				if z := (i * pow) + j; z <= result {
					ret = append(ret, z)
				}
			}
		}
	}
	return ret
}

func calculate(equations []equation, partTwo bool) int64 {
	result := int64(0)
	for _, eqn := range equations {
		new := eqn.numbers[0]
		for i := range len(eqn.numbers) - 1 {
			new = evaluate(new, eqn.numbers[i+1], eqn.result, partTwo)
		}

		if tmp := evaluate(new, nil, eqn.result, partTwo); len(tmp) > 0 {
			result += tmp[0]
		}
	}
	return result
}

func parse(input []string) []equation {
	ret := []equation{}
	for _, line := range input {
		parts := strings.Split(line, ": ")
		result, _ := strconv.ParseInt(parts[0], 10, 64)
		nums := [][]int64{}
		for _, n := range strings.Fields(parts[1]) {
			num, _ := strconv.ParseInt(n, 10, 64)
			nums = append(nums, []int64{num})
		}
		ret = append(ret, equation{result, nums})
	}
	return ret
}

func solve(input []string) solution {
	a := parse(input)
	return solution{calculate(a, false), calculate(a, true)}
}
