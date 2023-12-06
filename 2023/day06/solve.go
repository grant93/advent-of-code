package daysix

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`\d+`)

type races struct {
	distance, time int64
}

func parseNumbers(l string) []int64 {
	ret := []int64{}
	t := re.FindAll([]byte(l), -1)
	for _, s := range t {
		n, _ := strconv.ParseInt(string(s), 10, 64)
		ret = append(ret, n)
	}
	return ret
}

func simulate(r []races) int64 {
	total := int64(1)
	// time =  t_max - t_button
	// distance = t_button * time
	// distance = t_button(t_max - t_button)
	// distance = tm*t_button - t_button^2
	// t_button^2 - tm*t_button + distance = 0
	// solve with quadratic eqn (or wolfram..)
	// t_button = (tm - (tm^2 - 4d)^1/2)/2
	// t_button = ((tm^2 - 4d)^1/2) + tm)/2
	fmt.Println(r)
	for _, a := range r {
		if a.distance == 0 {
			continue
		}
		from := math.Floor((float64(a.time) - math.Sqrt(float64(a.time*a.time-4*a.distance))) / 2)
		to := math.Ceil((float64(a.time) + math.Sqrt(float64(a.time*a.time-4*a.distance))) / 2)
		total *= int64(to) - int64(from) - 1
	}

	return total
}

func solve(input []string) (int, int) {
	partOne := int64(0)
	partTwo := int64(0)
	data := make([]races, 4)
	data2 := make([]races, 1)
	for i, line := range input {
		vals := parseNumbers(line)
		// TODO: tidy this mess.
		for j, v := range vals {
			switch i {
			case 0:
				data[j].time = v
			case 1:
				data[j].distance = v
			default:
				continue
			}
		}
		vals = parseNumbers(strings.ReplaceAll(line, " ", ""))
		for j, v := range vals {
			switch i {
			case 0:
				data2[j].time = v
			case 1:
				data2[j].distance = v
			default:
				continue
			}
		}
	}
	partOne = simulate(data)
	partTwo = simulate(data2)
	return int(partOne), int(partTwo)
}
