package part1

import (
	"strings"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

// TODO: Parse



func Solve(lines []string) int{
	seeds := convertLineToSeeds(lines[0])


	// for _ , line := range lines {
	// 	// log.Printf("%+v", line)
	// }



	return 0
}


func convertLineToSeeds(line string) []int {
	seedsStr := strings.Fields(strings.Split(line, ":")[1])

	var seeds []int
	for _, str := range seedsStr {
		seeds = append(seeds, utils.Atoi(str))
	}

	return seeds
}
