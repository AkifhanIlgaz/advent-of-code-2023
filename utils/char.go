package utils

func StringDigitToInt[T rune | byte](s T) int {
	return int(s - '0')
}
