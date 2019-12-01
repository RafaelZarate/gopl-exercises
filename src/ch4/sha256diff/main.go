package main

import (
	"fmt"
	"os"
	"crypto/sha256"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Expected 2 args!")
		os.Exit(1)
	}

	hash1 := sha256.Sum256([]byte(os.Args[1]))
	hash2 := sha256.Sum256([]byte(os.Args[2]))

	var diffIndexes []int
	for i := 0; i < 256; i++ {
		if hash1[i] != hash2[i] {
			diffIndexes = append(diffIndexes, i)
		}
	}

	print("%v", diffIndexes)
}
