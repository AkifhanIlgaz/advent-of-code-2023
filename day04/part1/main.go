package part1

func Solve(lines []string) int {
	sumPoints := 0

	for _, line := range lines {
		scratchCard := convertLineToScratchCard(line)
		sumPoints += scratchCard.points()
	}

	return sumPoints
}
