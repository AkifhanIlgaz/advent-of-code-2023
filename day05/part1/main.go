package part1

import (
	"fmt"
	"strings"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

// TODO: Parse



func Solve(lines []string) int{
	seeds := parseSeeds(lines[0])

	fmt.Println(seeds)

	// for _ , line := range lines {
	// 	// log.Printf("%+v", line)
	// }



	return 0
}


func parseSeeds(line string) []int {
	seedsStr := strings.Fields(strings.Split(line, ":")[1])

	var seeds []int

	for _, str := range seedsStr {
		seeds = append(seeds, utils.Atoi(str))
	}

	return seeds
}
