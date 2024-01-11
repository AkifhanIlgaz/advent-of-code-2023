package part2

import (
	"strings"
)

type ScratchCard struct {
	Point  int
	Copies int
}

func generatePointsMap(lines []string) map[int]ScratchCard {
	pointsMap := map[int]ScratchCard{}

	for i, line := range lines {
		card := ScratchCard{
			Copies: 1,
		}
		id := i + 1
		splittedLine := strings.Split(line, ":")

		numbers := strings.TrimSpace(splittedLine[1])
		splittedNumbers := strings.Split(numbers, " | ")

		winningNumbers := bindNumsToMap(strings.Fields(splittedNumbers[0]))
		yourNumbers := bindNumsToMap(strings.Fields(splittedNumbers[1]))

		card.Point = matchCount(winningNumbers, yourNumbers)
		pointsMap[id] = card

	}

	return pointsMap
}

func copyCards(cards map[int]ScratchCard) {
	for id := 1; id < len(cards); id++ {
		card := cards[id]
		for i := 1; i <= card.Point; i++ {
			c := cards[id+i]
			c.Copies += card.Copies
			cards[id+i] = c
		}
	}
}

func totalCards(cards map[int]ScratchCard) int {
	sum := 0

	for _, card := range cards {
		sum += card.Copies
	}

	return sum
}

func matchCount(winningNumbers, yourNumbers map[string]bool) int {
	matchCount := 0

	for winningNum, _ := range winningNumbers {
		if yourNumbers[winningNum] {
			matchCount++
		}
	}

	return matchCount
}

func bindNumsToMap(arr []string) map[string]bool {
	m := map[string]bool{}

	for _, num := range arr {
		m[num] = true
	}

	return m
}
