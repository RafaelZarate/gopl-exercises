package main


import (
	"os"
	"fmt"
)

func isAnagram(mapOne map[rune]int, mapTwo map[rune]int) bool {
	lenMapOne := len(mapOne)
	if lenMapOne != len(mapTwo) {
		return false
	}

	for key, mapOneVal := range mapOne {
		mapTwoVal, ok := mapTwo[key]
		if !ok {
			return false
		} else if mapOneVal != mapTwoVal{
			return false
		}
	}

	return true
}

func createAnagramMapping(str string) map[rune]int {
	mapping := make(map[rune]int)

	for _, char := range str {
		val, ok := mapping[char]
		if ok {
			mapping[char] = val + 1
		} else {
			mapping[char] = 1
		}
	}

	return mapping
}


func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Printf("Expected at least 2 args")
		os.Exit(1)
	}
	firstStringMapping := createAnagramMapping(args[1])

	for pos, str := range args[1:] {
		validityString := " "
		if !isAnagram(firstStringMapping, createAnagramMapping(str)) {
			validityString = " not "
		}
		fmt.Printf("%s (pos %d) is%sa valid anagram!\n", str, pos + 2, validityString)
	}
	// fmt.Printf("%v", firstStringMapping)
}
