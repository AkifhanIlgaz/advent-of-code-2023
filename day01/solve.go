package day01

import (
	"fmt"

	firstpart "github.com/AkifhanIlgaz/advent-of-code-2023/day01/first_part"
	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

func Solve() {
	lines := utils.ReadLines("./day01/input.txt")

	fmt.Println("Part 1: ", firstpart.Solve(lines))
}
