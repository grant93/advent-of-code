package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	value int
	paths []string
}

type State struct {
	time int
	location string
	score int
	opened *list.List
}

func parse(data []string) map[string]Node {
	nodes := make(map[string]Node)
	for _, line := range data {
		node := Node{}
		tmp := strings.Fields(line)
		name := tmp[1]
		paths := tmp[9:]
		tmpVal := strings.Split(tmp[4], "=")[1]
		node.value, _ = strconv.Atoi(tmpVal[:len(tmpVal)-1])
		for _, path := range paths {
			fmt.Println(path)
			if strings.HasSuffix(path, ",") {
				path = string(path[:len(path)-1])
			}
			node.paths = append(node.paths, path)
		}
		nodes[name] = node
	}
	return nodes
}


func solve(data []string) []int {
	nodes := parse(data)
	fmt.Println(nodes)
	answers := make([]int, 2)

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
