package four

import (
	"2022/go/four/camp"
	"fmt"
)

func Run(filePath string) {
	pairs := camp.GetPairs(filePath)
	sum := camp.NumberAssignmentsWithRangeFullyContainOther(pairs)
	fmt.Println("Total pairs: ", len(pairs))
	fmt.Println("Number of pairs with one range containing the other: ", sum)
	sum = camp.NumberAssignmentsWithOverlappingRanges(pairs)
	fmt.Println("Number of pairs with overlapping ranges: ", sum)
}
