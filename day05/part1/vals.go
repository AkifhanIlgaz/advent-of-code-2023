package part1

import (
	"strings"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)



type Vals struct {
	DestinationRange []int
	SourceRange      []int
}

func convertLineToVals(line string) Vals {
	vals := strings.Fields(line)
	destinationStart, sourceStart, rangeLength := utils.Atoi(vals[0]), utils.Atoi(vals[1]),utils.Atoi(vals[2])

	destinationRange := applyRange(destinationStart, rangeLength)
	sourceRange := applyRange(sourceStart, rangeLength)

	return Vals{
		DestinationRange: destinationRange,
		SourceRange: sourceRange,
	}
}



// r => range
func applyRange(start, rangeLength int) []int {
	vals := []int{start}

	for i := 1; i <= rangeLength; i++ {
		vals = append(vals, start + i)
	} 

	return vals
}