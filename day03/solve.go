package day03

import (
	"fmt"

	"github.com/AkifhanIlgaz/advent-of-code-2023/day03/part1"
	"github.com/AkifhanIlgaz/advent-of-code-2023/day03/part2"
	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

func Solve() {
	lines := utils.ReadLines("./day03/input.txt")

	fmt.Println("Part 1:", part1.Solve(lines))
	fmt.Println("Part 2:", part2.Solve(lines))
}
