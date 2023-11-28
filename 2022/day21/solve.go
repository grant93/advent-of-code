package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var store map[string]interface{}

func resolve(key string) int {
	item := store[key]
	switch item.(type) {
	case []string:
		i := item.([]string)
		switch i[1] {
		case "/":
			return resolve(i[0]) / resolve(i[2])
		case "+":
			return resolve(i[0]) + resolve(i[2])
		case "-":
			return resolve(i[0]) - resolve(i[2])
		case "*":
			return resolve(i[0]) * resolve(i[2])
		case "=":
			fmt.Println(resolve(i[0]), resolve(i[2]))
			if resolve(i[0]) == resolve(i[2]) {
				return 1
			} else {
				return 0
			}
		}
	case [1]int:
		i := item.([1]int)
		return i[0]
	}
	return -1
}

func resolveUp(item string, target int) {
	
}

func parse(data []string) map[string]interface{} {
	store = make(map[string]interface{})
	for _, line := range data {
		entries := strings.Split(line, " ")
		if len(entries) == 2 {
			tmp, _ := strconv.Atoi(entries[1])
			name := entries[0][:len(entries[0])-1]
			store[name] = [1]int{tmp}
		} else {
			name := entries[0][:len(entries[0])-1]
			store[name] = entries[1:]
		}
	}
	return store
}

func findEntry(entry string) (string, int) {
	for key, value := range store {
		if val, ok := value.([]string); ok {
			if val[0] == entry {
				return key, 0
			} else if val[2] == entry {
				return key, 2
			}
		}
	}
	return "", -1
}

func partTwo() int {
	// resolve the right side of root as "target"
	target := resolve(store["root"].([]string)[2])
	fmt.Println(target)
	key, _ := findEntry("humn")
	fmt.Println(store[key].([]string))
	right := resolve(store[key].([]string)[2])
	fmt.Println(right)
	store[key].([]string)[2] = strconv.Itoa(right)
	fmt.Println(store[key].([]string))

	return 0
}

func solve(data []string) []int {
	answers := make([]int, 2)
	store := parse(data);
	fmt.Println(store)
	answers[0] = resolve("root") 
	a := store["root"].([]string)
	a[1] = "="
	store["root"] = a
	answers[1] = partTwo()

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
