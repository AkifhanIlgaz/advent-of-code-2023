package part1

import (
	"strings"

	"github.com/AkifhanIlgaz/advent-of-code-2023/utils"
)



type Ranges struct {
	DestinationRange []int
	SourceRange      []int
}

func convertLineToVals(line string) Ranges {
	ranges := strings.Fields(line)
	destinationStart, sourceStart, rangeLength := utils.Atoi(ranges[0]), utils.Atoi(ranges[1]),utils.Atoi(ranges[2])

	destinationRange := applyRange(destinationStart, rangeLength)
	sourceRange := applyRange(sourceStart, rangeLength)

	return Ranges{
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