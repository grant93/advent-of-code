package main

import (
	"bufio"
	"fmt"
	"os"
)

type Scores struct {
	score_us   int
	score_them int
}

var scoring map[string]Scores
var scoring2 map[string]Scores

/* precompute scoring maps */
func _init() {
	scoring["A X"] = Scores{score_us: 4, score_them: 4} // rock->rock draw
	scoring["B X"] = Scores{score_us: 1, score_them: 8} // rock->paper loss
	scoring["C X"] = Scores{score_us: 7, score_them: 3} // rock->scissors win
	scoring["A Y"] = Scores{score_us: 8, score_them: 1} // paper->rock win
	scoring["B Y"] = Scores{score_us: 5, score_them: 5} // paper->paper draw
	scoring["C Y"] = Scores{score_us: 2, score_them: 9} // paper->scissors loss
	scoring["A Z"] = Scores{score_us: 3, score_them: 7} // scissors->rock loss
	scoring["B Z"] = Scores{score_us: 9, score_them: 2} // scissors->paper win
	scoring["C Z"] = Scores{score_us: 6, score_them: 6} // scissors->scissors draw

	scoring2["A X"] = Scores{score_us: 3, score_them: 7} // scissors->rock lose
	scoring2["B X"] = Scores{score_us: 1, score_them: 8} // rock->paper lose
	scoring2["C X"] = Scores{score_us: 2, score_them: 9} // paper->scissors lose
	scoring2["A Y"] = Scores{score_us: 4, score_them: 4} // rock->rock draw
	scoring2["B Y"] = Scores{score_us: 5, score_them: 5} // paper->paper draw
	scoring2["C Y"] = Scores{score_us: 6, score_them: 6} // scissors->scissors draw
	scoring2["A Z"] = Scores{score_us: 8, score_them: 1} // paper->rock win
	scoring2["B Z"] = Scores{score_us: 9, score_them: 2} // scissors->paper win
	scoring2["C Z"] = Scores{score_us: 7, score_them: 3} // rock->scissors win
}

func solve(data []string) []int {
	answers := make([]int, 2)
	_init()

	for _, line := range data {
		scores := scoring[line]
		answers[0] += scores.score_us

		scores = scoring2[line]
		answers[1] += scores.score_us

	}
	return answers
}

/* standard boilerplate */
func main() {
	var data []string
	scoring = make(map[string]Scores)
	scoring2 = make(map[string]Scores)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	fmt.Println(solve(data))
}
