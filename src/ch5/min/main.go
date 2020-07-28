
package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {

	var ints []int
	for _, arg := range os.Args[1:] {
		newInt, _ := strconv.Atoi(arg)
		ints = append(ints, newInt)
	}

	minValue := min(ints...)
	fmt.Printf("Result %d", minValue)
}

func min(nums ...int) int {
	const MaxInt = int(^uint(0) >> 1)

	min := MaxInt
	for _, n := range nums {
		if (n < min) {
			min = n
		}
	}

	return min
}