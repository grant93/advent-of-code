package daytwo

type Scores struct {
	score_us   int
	score_them int
}

var (
	scoring = map[string]Scores{
		"A X": Scores{4, 4}, // rock->rock draw
		"B X": Scores{1, 8}, // rock->paper loss
		"C X": Scores{7, 3}, // rock->scissors win
		"A Y": Scores{8, 1}, // paper->rock win
		"B Y": Scores{5, 5}, // paper->paper draw
		"C Y": Scores{2, 9}, // paper->scissors loss
		"A Z": Scores{3, 7}, // scissors->rock loss
		"B Z": Scores{9, 2}, // scissors->paper win
		"C Z": Scores{6, 6}, // scissors->scissors draw
	}
	scoring2 = map[string]Scores{
		"A X": Scores{3, 7}, // scissors->rock lose
		"B X": Scores{1, 8}, // rock->paper lose
		"C X": Scores{2, 9}, // paper->scissors lose
		"A Y": Scores{4, 4}, // rock->rock draw
		"B Y": Scores{5, 5}, // paper->paper draw
		"C Y": Scores{6, 6}, // scissors->scissors draw
		"A Z": Scores{8, 1}, // paper->rock win
		"B Z": Scores{9, 2}, // scissors->paper win
		"C Z": Scores{7, 3}, // rock->scissors win
	}
)

func solve(data []string) (int, int) {
	partOne := 0
	partTwo := 0

	for _, line := range data {
		scores := scoring[line]
		partOne += scores.score_us

		scores = scoring2[line]
		partTwo += scores.score_us

	}
	return partOne, partTwo
}
