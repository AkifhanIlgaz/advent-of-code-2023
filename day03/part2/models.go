package part2

type PartNumber struct {
	number   int
	row      int
	startCol int
	endCol   int
}

func createEngineSchematic(lines []string) [][]rune {
	schematic := make([][]rune, len(lines))

	for row, line := range lines {
		schematic[row] = make([]rune, len(line))
		for col, char := range line {
			schematic[row][col] = char
		}
	}

	return schematic
}

func getPartNumbers(schematic [][]rune) []PartNumber {
	partNumbers := make([]PartNumber, 0)

	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			currentNumber, start, end := 0, 0, 0

			for ; col < len(schematic[row]); col++ {
				if schematic[row][col] < '0' || schematic[row][col] > '9' {
					break
				}

				if currentNumber == 0 {
					start = col
					end = col
				} else {
					end = col
				}

				currentNumber = currentNumber*10 + int(schematic[row][col]-'0')
			}

			if currentNumber > 0 {
				for r := row - 1; r <= row+1; r++ {
					for c := start - 1; c <= end+1; c++ {
						if r < 0 || r >= len(schematic) || c < 0 || c >= len(schematic[r]) {
							continue
						}

						if schematic[r][c] != '.' && (schematic[r][c] < '0' || schematic[r][c] > '9') {
							partNumbers = append(partNumbers, PartNumber{currentNumber, row, start, end})
						}
					}
				}
			}
		}
	}

	return partNumbers
}

func sumGearRatios(grid [][]rune, partNumbers []PartNumber) int {
	result := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '*' {
				count := 0
				adjacentNumbers := make([]int, 0)

				for _, number := range partNumbers {
					if number.row >= row-1 && number.row <= row+1 && number.endCol >= col-1 && number.startCol <= col+1 {
						count++
						adjacentNumbers = append(adjacentNumbers, number.number)
					}
				}

				if count == 2 {
					result += adjacentNumbers[0] * adjacentNumbers[1]
				}
			}
		}
	}

	return result
}
