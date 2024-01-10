package firstpart

import (
	"unicode"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)

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
	for i, char := range s {
		if unicode.IsDigit(char) {
			return utils.StringDigitToInt(s[i])
		}
	}
	panic("No digit found")
}

func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return utils.StringDigitToInt(s[i])
		}
	}
	panic("No digit found")
}
