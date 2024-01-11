package part2

func Solve(lines []string) int {
	cardsMap := generatePointsMap(lines)
	copyCards(cardsMap)

	return totalCards(cardsMap)
}
