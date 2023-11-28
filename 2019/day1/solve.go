package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(mass int) int{
	if mass <= 0 {
		return 0 
	}
	base_fuel := int(math.Floor(float64(mass)/3.0)) - 2
	fuel := calculateFuel(base_fuel)
	if fuel < 0 {
		fuel = 0
	}
	return base_fuel + fuel
}

func partTwo(data []string) int {
	total := 0
	for _, mass := range data {
		tmp, _ := strconv.Atoi(mass)
		total += calculateFuel(tmp)
	}
	return total
}

func partOne(data []string) int {
	total := 0
	for _, mass := range data {
		tmp, _ := strconv.Atoi(mass)
		total += int(math.Floor(float64(tmp)/3.0)) - 2
	}
	return total
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
