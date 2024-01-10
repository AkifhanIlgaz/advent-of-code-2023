package part1

import (
	"log"
	"strconv"
	"strings"
)

const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

type Game struct {
	Id   int
	Sets []Set
}

func (game Game) isPossible() bool {
	for _, set := range game.Sets {
		if set.isImpossible() {
			return false
		}
	}

	return true
}

func parseGame(g string) Game {
	parsed := strings.Split(g, ":")

	return Game{
		Id:   parseGameId(parsed[0]),
		Sets: parseSets(parsed[1]),
	}
}

func parseGameId(s string) int {
	s = strings.TrimSpace(s)

	splitted := strings.Fields(s)

	gameId, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic("game id is not a number")
	}

	return gameId
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (set Set) isImpossible() bool {
	return set.Red > RedCubes || set.Blue > BlueCubes || set.Green > GreenCubes
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
