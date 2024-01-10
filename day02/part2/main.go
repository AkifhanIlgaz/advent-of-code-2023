package part2

func Solve(lines []string) int {
	sumOfPowers := 0

	for _, line := range lines {
		game := convertLineToGame(line)
		sumOfPowers += game.findFewestNumOfCubes().power()
	}

	return sumOfPowers
}
