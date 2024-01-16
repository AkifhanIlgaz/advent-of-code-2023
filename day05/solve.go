package day05

import (
	"log"
	"os"
	"strings"

	"github.com/AkifhanIlgaz/advent-of-code-2023/day05/part1"
)

func Solve() {
	b, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	part1.Solve(lines)
}
