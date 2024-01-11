package part1

func Solve(lines []string) int {
	engineSchematic := createEngineSchematic(lines)
	partNumbers := getPartNumbers(engineSchematic)

	return sumPartNumbers(partNumbers)

}
