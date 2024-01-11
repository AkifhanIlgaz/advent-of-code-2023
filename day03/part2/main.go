package part2

func Solve(lines []string) int {
	grid := createEngineSchematic(lines)
	partNumbers := getPartNumbers(grid)

	return sumGearRatios(grid, partNumbers)
}
