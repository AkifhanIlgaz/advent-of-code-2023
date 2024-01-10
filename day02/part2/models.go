package part2

import (
	"log"
	"strconv"
	"strings"
)

type Game struct {
	Sets []Set
}

func parseGame(g string) Game {
	parsed := strings.Split(g, ":")

	return Game{
		Sets: parseSets(parsed[1]),
	}
}

func (game Game) findFewestNumOfCubes() Set {
	var set Set

	for _, s := range game.Sets {
		set.Red = max(set.Red, s.Red)
		set.Green = max(set.Green, s.Green)
		set.Blue = max(set.Blue, s.Blue)
	}

	return set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (s Set) Power() int {
	return s.Red * s.Green * s.Blue
}

func parseSets(s string) []Set {
	s = strings.TrimSpace(s)
	splittedSets := strings.Split(s, ";")

	var sets []Set

	for _, set := range splittedSets {
		sets = append(sets, parseSet(set))
	}

	return sets
}

func parseSet(s string) Set {
	var set Set
	s = strings.TrimSpace(s)

	cubes := strings.Split(s, ",")

	for _, cube := range cubes {
		setCube(&set, cube)
	}

	return set
}

func setCube(set *Set, cube string) {
	count, color := splitCube(cube)
	switch color {
	case "red":
		set.Red = count
	case "green":
		set.Green = count
	case "blue":
		set.Blue = count
	}
}

func splitCube(s string) (int, string) {
	splitted := strings.Fields(s)

	count, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Fatal("color count is not a number")
	}

	return count, splitted[1]
}
