package d2

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func getSurfaceArea(l, w, h int) (surfaceArea int) {
	return (2 * l * w) + (2 * w * h) + (2 * h * l)
}

func getExtraPaper(l, w, h int) (extraPaper int) {
	return min(l*w, w*h, h*l)
}

func convertStrSliceToInt(strSlice []string) (intSlice []int) {
	for _, s := range strSlice {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, i)
	}

	return
}

func getRibbonLength(l, w, h int) (ribbonLength int) {
	// 2 * smallest side + 2 * second smallest side
	sides := []int{l, w, h}
	slices.Sort(sides) //asc order

	return (2*sides[0] + 2*sides[1])
}

func getBowLength(l, w, h int) (bowLength int) {
	return l * w * h
}

// solution for d2 part 1
// 1588178
func P1(content string) {
	// each present calculation is separated by newline
	presents := strings.Split(content, "\n")

	total := 0
	for _, p := range presents {
		parts := strings.Split(p, "x")
		if len(parts) != 3 {
			continue
		}
		intParts := convertStrSliceToInt(parts)
		l, w, h := intParts[0], intParts[1], intParts[2]

		total += getSurfaceArea(l, w, h)
		total += getExtraPaper(l, w, h)

	}
	fmt.Println(total)
}

// solution for d2 part 2
// 3783758
func P2(content string) {

	// each present calculation is separated by newline
	presents := strings.Split(content, "\n")

	total := 0
	for _, p := range presents {
		parts := strings.Split(p, "x")
		if len(parts) != 3 {
			continue
		}
		intParts := convertStrSliceToInt(parts)
		l, w, h := intParts[0], intParts[1], intParts[2]

		total += getRibbonLength(l, w, h)
		total += getBowLength(l, w, h)

	}
	fmt.Println(total)
}
