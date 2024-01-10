package part2

import (
	"strings"
	"unicode"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

var spelledDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Solve(lines []string) int {
	return getCalibrationValues(lines)
}

func getCalibrationValues(documents []string) int {
	totalCalibrationValue := 0

	for _, doc := range documents {
		var (
			firstDigit = getFirstDigit(doc)
			lastDigit  = getLastDigit(doc)
		)
		totalCalibrationValue += firstDigit * 10
		totalCalibrationValue += lastDigit
	}

	return totalCalibrationValue
}

func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if d, found := containsSpelledDigit(s[:i]); found {
			return d
		} else if unicode.IsDigit(rune(s[i])) {
			return utils.StringDigitToInt(s[i])
		}
	}

	panic("No digit found")
}

func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if d, found := containsSpelledDigit(s[i:]); found {
			return d
		} else if unicode.IsDigit(rune(s[i])) {
			return utils.StringDigitToInt(s[i])
		}
	}

	panic("No digit found")
}

func containsSpelledDigit(s string) (int, bool) {
	for k, v := range spelledDigits {
		if strings.Contains(s, k) {
			return v, true
		}
	}
	return 0, false
}
