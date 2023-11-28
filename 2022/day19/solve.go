package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Costs struct {
	ore, clay, obsidian int
}

type BluePrint struct {
	id                                         int
	oreCost, clayCost, obsidianCost, geodeCost Costs
}

type State struct {
	inventory [4]int
	bots      [4]int
	time      int
}

func toInt(data []string) []int {
	converted := []int{}
	for _, item := range data {
		tmp, _ := strconv.Atoi(item)
		converted = append(converted, tmp)
	}
	return converted
}

func parse(data []string, blueprints *[]BluePrint) {
	r, _ := regexp.Compile(`(\d+)`)
	for i, line := range data {
		costs := toInt(r.FindAllString(line, -1))
		blueprint := BluePrint{
			i + 1,
			Costs{costs[1], 0, 0},
			Costs{costs[2], 0, 0},
			Costs{costs[3], costs[4], 0},
			Costs{costs[5], 0, costs[6]},
		}
		*blueprints = append(*blueprints, blueprint)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func move(state State) State {
	state.inventory[0] += state.bots[0]
	state.inventory[1] += state.bots[1]
	state.inventory[2] += state.bots[2]
	state.inventory[3] += state.bots[3]
	state.time--

	return state
}

func bfs(blueprint BluePrint, timeout int) int {
	max := 0
	queue := []State{State{[4]int{0, 0, 0, 0}, [4]int{1, 0, 0, 0}, timeout}}
	visited := map[State]bool{}
	maxOre := maxInt(maxInt(blueprint.oreCost.ore, blueprint.clayCost.ore), maxInt(blueprint.obsidianCost.ore, blueprint.geodeCost.ore))

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		max = maxInt(max, state.inventory[3])

		if state.time == 0 {
			continue
		}

		/* don't track more than we can spend - minimises searchspace */
		if state.inventory[0] > (state.time*maxOre)-((state.time-1)*state.bots[0]) {
			state.inventory[0] = (state.time * maxOre) - ((state.time - 1) * state.bots[0])
		}

		if visited[state] {
			continue
		}
		visited[state] = true

		if state.inventory[0] >= blueprint.geodeCost.ore && state.inventory[2] >= blueprint.geodeCost.obsidian {
			newState := state
			newState.inventory[0] -= blueprint.geodeCost.ore
			newState.inventory[2] -= blueprint.geodeCost.obsidian
			newState = move(newState)
			newState.bots[3]++
			queue = append(queue, newState)
		} else {
			if state.inventory[0] >= blueprint.oreCost.ore && state.bots[0] < maxOre {
				newState := state
				newState.inventory[0] -= blueprint.oreCost.ore
				newState = move(newState)
				newState.bots[0]++
				queue = append(queue, newState)
			}
			if state.inventory[0] >= blueprint.clayCost.ore && state.bots[1] < blueprint.obsidianCost.clay {
				newState := state
				newState.inventory[0] -= blueprint.clayCost.ore
				newState = move(newState)
				newState.bots[1]++
				queue = append(queue, newState)
			}
			if state.inventory[0] >= blueprint.obsidianCost.ore && state.inventory[1] >= blueprint.obsidianCost.clay && state.bots[2] < blueprint.geodeCost.obsidian {
				newState := state
				newState.inventory[0] -= blueprint.obsidianCost.ore
				newState.inventory[1] -= blueprint.obsidianCost.clay
				newState = move(newState)
				newState.bots[2]++
				queue = append(queue, newState)
			}

			state = move(state)
			queue = append(queue, state)
		}
	}
	return max
}

func partOne(blueprints []BluePrint) int {
	timeout := 24
	answer := 0
	for _, bp := range blueprints {
		answer += (bp.id * bfs(bp, timeout))
	}

	return answer
}

func partTwo(blueprints []BluePrint) int {
	timeout := 32
	answer := 1
	for _, bp := range blueprints[:3] {
		answer *= bfs(bp, timeout)
	}

	return answer
}

func solve(data []string) []int {
	blueprints := []BluePrint{}
	parse(data, &blueprints)
	fmt.Println(blueprints)
	answers := make([]int, 2)
	answers[0] = partOne(blueprints)
	answers[1] = partTwo(blueprints)

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
