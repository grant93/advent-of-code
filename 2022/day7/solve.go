package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name    string
	root    *dir
	subdirs []*dir
	size    int
}

const DiskSpace int = 70000000
const RequiredSpace int = 30000000

/* forgot about todays challenge until just before bed, then decided to give it a go... */

/* denug helper for tree visualisation */
func printTree(root *dir, level int) {
	fmt.Println(strings.Repeat("  ", level), root.name, root.size)
	level++
	for _, subdir := range root.subdirs {
		printTree(subdir, level)
	}
}

/* part two */
func findSpace(root *dir, space int) *dir {
	var smallest *dir = nil
	for _, subdir := range root.subdirs {
		if subdir.size >= space && (smallest == nil || subdir.size <= smallest.size) {
			smallest = subdir
		}
		tmp := findSpace(subdir, space)
		if tmp != nil && tmp.size < smallest.size {
			smallest = tmp
		}
	}
	return smallest
}

/* part one */
func sumTree(root *dir) int {
	total := 0
	for _, subdir := range root.subdirs {
		if subdir.size <= 100000 {
			total += subdir.size
		}
		total += sumTree(subdir)
	}
	return total
}

/* kinda gross, shouldn't write code after midnight.. (adds values back up the tree) */
func sumUp(directory *dir, size int) {
	current := directory.root
	for current != nil {
		current.size += size
		current = current.root
	}
}

func solve(data []string) []int {
	answers := make([]int, 2)
	root := dir{"/", nil, make([]*dir, 0), 0}
	current := &root

	for _, line := range data {
		// cd is the only command we really care about
		if strings.HasPrefix(line, "$ cd") {
			if strings.HasPrefix(line[5:], "/") {
				current = &root
			} else if strings.HasPrefix(line[5:], "..") {
				current = current.root
			} else {
				name := line[5:len(line)]
				for _, entry := range current.subdirs {
					if entry.name == name {
						current = entry
						break
					}
				}
			}
		} else {
			if strings.HasPrefix(line, "dir") {
				entry := dir{line[4:len(line)], current, make([]*dir, 0), 0}
				current.subdirs = append(current.subdirs, &entry)
			} else {
				tmp := strings.Fields(line)
				size, _ := strconv.Atoi(tmp[0])
				current.size += size

				sumUp(current, size)
			}

		}
	}
	//printTree(&root, 0)
	answers[0] = sumTree(&root)
	requiredSpace := RequiredSpace - (DiskSpace - root.size)
	answers[1] = findSpace(&root, requiredSpace).size

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
