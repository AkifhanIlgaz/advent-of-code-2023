package utils

import (
	"log"
	"strconv"
)

func Atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(s, "is not a number")
	}
	return val
}