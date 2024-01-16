package part1

import (
	"math"
	"strings"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

func Solve(lines []string) int {
	seeds := convertLineToSeeds(lines[0])

	separatorIndexes := separatorIndexes(lines)

	maps := NewMaps()

	for i := 0; i < len(separatorIndexes)-1; i++ {
		start, end := separatorIndexes[i], separatorIndexes[i+1]

		maps.update(convertLinesToUpdateMap(lines[start+1 : end]))
	}

	maps.update(convertLinesToUpdateMap(lines[separatorIndexes[len(separatorIndexes)-1]+1:]))

	lowestLocation := math.MaxInt

	for _, seed := range seeds {
		lowestLocation = min(lowestLocation, maps.findLocation(seed))

	}

	return lowestLocation
}

func convertLineToSeeds(line string) []int {
	seedsStr := strings.Fields(strings.Split(line, ":")[1])

	var seeds []int
	for _, str := range seedsStr {
		seeds = append(seeds, utils.Atoi(str))
	}

	return seeds
}

func separatorIndexes(lines []string) []int {
	var indexes []int

	for i, line := range lines {
		if isSeparator(line) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func isSeparator(s string) bool {
	return []byte(s)[0] == 13
}
