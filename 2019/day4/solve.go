package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


func partOne(data string) int {
	values := strings.Split(data, "-")
	min, _ := strconv.Atoi(values[0])
	max, _ := strconv.Atoi(values[1])
	
	count := 0
	for i:=min; i<max; i++ {
		rule1 := false
		rule2 := true 
		str := strconv.Itoa(i)
		for j:=1; j<len(str); j++ { 
			if str[j] == str[j-1] {
				rule1 = true
			}
			if int(str[j-1] - '0') > int(str[j] - '0') {
				rule2 = false
			}
		}
		if rule1 && rule2 {
			count++
		}
	}
	return count
}

func checkGroups(index int, value byte, full string) bool{
	if strings.Contains(full, strings.Repeat(string(value), 3)) {
		return false	
	} else {
		return true
	}
}

func partTwo(data string) int {
	values := strings.Split(data, "-")
	min, _ := strconv.Atoi(values[0])
	max, _ := strconv.Atoi(values[1])
	
	count := 0
	for i:=min; i<max; i++ {
		rule1 := false
		rule2 := true 
		str := strconv.Itoa(i)
		for j:=1; j<len(str); j++ { 
			if str[j] == str[j-1] && checkGroups(j, str[j], str) {
				rule1 = true
			}
			if int(str[j-1] - '0') > int(str[j] - '0') {
				rule2 = false
			}
		}
		if rule1 && rule2 {
			count++
		}
	}
	return count
}


func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(data[0])
	answers[1] = partTwo(data[0])
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
