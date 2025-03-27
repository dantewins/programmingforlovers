package main

import (
	"fmt"
)

func main() {
	fmt.Println("Substrings and (subslices).")

	s := "Hi lovers!"

	fmt.Println(s[1:5])
	fmt.Println(s[:7])
	fmt.Println(s[4:])

	a := make([]int, 6)
	for i := range a {
		a[i] = 2*i + 1
	}

	fmt.Println(a)
	fmt.Println(a[3:5])
	fmt.Println(a[3:])
	fmt.Println(a[:4])

	pattern := "ATA"
	text := "ATATA"

	fmt.Println(PatternCount(pattern, text))
	fmt.Println(StartingIndices(pattern, text))
}

// PatternCount takes as input two strings pattern and text. It returns the number of times that pattern occurs as a substring of text, including overlaps

func PatternCount(pattern, text string) int {

	return len(StartingIndices(pattern, text))

	/*
		count := 0
		n := len(text)
		k := len(pattern)

		for i := 0; i < n-k+1; i++ {
			if pattern == text[i:i+k] {
				count++
			}
		}

		return count
	*/
}

// StartingIndices takes as input two strings pattern and text. It returns the collection of all starting positions of a pattern as a substrings of text, including overlaps
func StartingIndices(pattern, text string) []int {
	positions := make([]int, 0)

	n := len(text)
	k := len(pattern)

	for i := 0; i < n-k+1; i++ {
		if pattern == text[i:i+k] {
			positions = append(positions, i)
		}
	}

	return positions
}
