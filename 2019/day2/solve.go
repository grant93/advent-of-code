package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/grant93/advent-of-code/2019/lib/helpers"
	"github.com/grant93/advent-of-code/2019/lib/intcode"
)


func partOne(data []string) int {
	instructions := helpers.DeliminatedStringToIntSlice(data[0], ",")
	instructions[1] = 12
	instructions[2] = 2
	machine := &intcode.IntCode{}
	machine.Init(instructions)
	machine.Compute()
	result, _ := machine.ReadAddress(0)
	return result
}

func partTwo(data []string) int {
	for i:=0; i < 100; i++ {
		for j:=0; j < 100; j++ {
			instructions := helpers.DeliminatedStringToIntSlice(data[0], ",")
			instructions[1] = i
			instructions[2] = j
			machine := &intcode.IntCode{}
			machine.Init(instructions)
			machine.Compute()
			result, _ := machine.ReadAddress(0)
			if result == 19690720 {
				return (100 * instructions[1]) + instructions[2]
			}
		}
	}
	return 0
}


func solve(data []string) []int {
	answers := make([]int, 2)
	answers[0] = partOne(data) 
	answers[1] = partTwo(data) 
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
