package part1

import (
	"math"
	"strings"
)

type ScratchCard struct {
	winningNumbers map[string]bool
	yourNumbers    map[string]bool
}

func convertLineToScratchCard(line string) ScratchCard {
	splittedLine := strings.Split(line, ":")
	numbers := strings.TrimSpace(splittedLine[1])
	splittedNumbers := strings.Split(numbers, " | ")

	return ScratchCard{
		winningNumbers: bindNumsToMap(strings.Fields(splittedNumbers[0])),
		yourNumbers:    bindNumsToMap(strings.Fields(splittedNumbers[1])),
	}
}

func (card ScratchCard) points() int {
	matchCount := 0

	for winningNum, _ := range card.winningNumbers {
		if card.yourNumbers[winningNum] {
			matchCount++
		}
	}

	if matchCount == 0 {
		return 0
	}

	return int(math.Exp2(float64(matchCount) - 1))
}

func bindNumsToMap(arr []string) map[string]bool {
	m := map[string]bool{}

	for _, num := range arr {
		m[num] = true
	}

	return m
}
