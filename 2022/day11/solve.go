package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Plus int = iota
	Minus
	Divide
	Multiply
)

type Operation struct {
	operator int
	value    int
	self     bool
}

type Monkey struct {
	items     *[]int
	operation Operation
	condition Operation
	success   int
	failure   int
	count     int
}

/* gross parsing because I was too lazy to regex */
func parse(data []string) []Monkey {
	var monkey Monkey
	monkeys := []Monkey{}
	index := -1

	for _, line := range data {
		if strings.HasPrefix(line, "Monkey") {
			monkey = Monkey{}
			index++
		} else if strings.HasPrefix(line, "  Starting") {
			position := strings.Index(line, ":")
			items := strings.Split(string(line[position+2:]), ", ")
			monkey.items = &[]int{}

			for _, item := range items {
				worry, _ := strconv.Atoi(item)
				tmp := append(*(monkey.items), worry)
				monkey.items = &tmp
			}

		} else if strings.HasPrefix(line, "  Operation") {
			position := strings.Index(line, "old ") + 4
			operation := strings.Fields(string(line[position:]))

			switch operation[0] {
			case "*":
				monkey.operation.operator = Multiply
			case "/":
				monkey.operation.operator = Divide
			case "+":
				monkey.operation.operator = Plus
			case "-":
				monkey.operation.operator = Minus
			}

			if operation[1] == "old" {
				monkey.operation.self = true
			} else {
				value, _ := strconv.Atoi(operation[1])
				monkey.operation.value = value
			}

		} else if strings.HasPrefix(line, "  Test") {
			tmp := strings.Split(line, "by ")
			value, _ := strconv.Atoi(tmp[1])
			monkey.condition.value = value
			monkey.condition.operator = Divide
		} else if strings.HasPrefix(line, "    If true") {
			monkey.success = int(line[len(line)-1] - '0')
		} else if strings.HasPrefix(line, "    If false") {
			monkey.failure = int(line[len(line)-1] - '0')
		} else {
			monkeys = append(monkeys, monkey)
		}
	}
	monkeys = append(monkeys, monkey)
	return monkeys
}

/* i hate math tricks. */
var denominators int = 13 * 2 * 19 * 11 * 7 * 5 * 3 * 17 * 23

func inspect(item int, operation Operation) int {
	value := item
	if !operation.self {
		value = operation.value
	}
	switch operation.operator {
	case Plus:
		value += item
	case Minus:
		value -= item
	case Multiply:
		value *= item
	case Divide:
		value /= item
	}
	return value
}

func simulate(monkeys []Monkey, partOne bool) int {
	iterations := 20
	if !partOne {
		iterations = 10000
	}
	for i := 0; i < iterations; i++ {
		for j, monkey := range monkeys {
			for _, item := range *(monkey.items) {
				value := inspect(item, monkey.operation)
				if partOne {
					value /= 3
				} else {
					value %= denominators
				}
				if value%monkey.condition.value == 0 {
					tmp := append(*(monkeys[monkey.success].items), value)
					monkeys[monkey.success].items = &tmp
				} else {
					tmp := append(*(monkeys[monkey.failure].items), value)
					monkeys[monkey.failure].items = &tmp
				}
				monkeys[j].count++
			}
			tmp := []int{}
			monkeys[j].items = &tmp
		}
	}

	first, second := 0, 0
	for _, monkey := range monkeys {
		if monkey.count > first {
			second = first
			first = monkey.count
		} else if monkey.count > second {
			second = monkey.count
		}
	}
	return first * second
}

func solve(data []string) []int {
	answers := make([]int, 2)
	monkeys := parse(data)
	monkeys2 := make([]Monkey, len(monkeys))
	copy(monkeys2, monkeys)

	answers[0] = simulate(monkeys, true)
	answers[1] = simulate(monkeys2, false)

	return answers
}

/* standard boilerplate */
func main() {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	fmt.Println(solve(data))
}
