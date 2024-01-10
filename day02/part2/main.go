package part2

func Solve(lines []string) int {
	sumOfPowers := 0

	for _, line := range lines {
		game := parseGame(line)
		sumOfPowers += game.findFewestNumOfCubes().Power()
	}

	return sumOfPowers
}
