package part1

func Solve(lines []string) int {
	sumOfImpossibleIds := 0

	for _, line := range lines {
		game := convertLineToGame(line)
		if game.isPossible() {
			sumOfImpossibleIds += game.Id
		}
	}

	return sumOfImpossibleIds
}
