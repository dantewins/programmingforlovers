package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Strings.")

	fmt.Println(string('A'))
	fmt.Println(string(45))

	fmt.Println(strconv.Itoa(45))

	j, err := strconv.Atoi("23")

	// the conversion is OK when err variable is equal to nil
	if err != nil {
		// a problem happened
		panic(err)
	}

	fmt.Println(j)

	pi, err2 := strconv.ParseFloat("3.14", 64)

	if err2 != nil {
		// a problem happened
		panic(err2)
	}

	fmt.Println("The value of pi is", pi)

	var s string = "Hi"
	t := "Lovers"

	// concentrate these strings with + operation

	u := s + t
	fmt.Println(u)

	fmt.Println("The first symbol of u is", string(u[0]))
	fmt.Println("The final symbol of u is", string(u[len(u)-1]))

	if t[2] == 'v' {
		fmt.Println("The symbol at position 2 of t is v.")
	}

	dna := "ACCGATCGATCGTTTAGCTGGCGTATCGATCGATCTCTCTCGAGATATATAGGAGCGCGTACTACGACGACTACGATCGATCGTACTAGCTAGCTAGTCGATCGATCGTACGTAGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCTAGCTGTACGTAGTTTCGCTGCTCGCGCGCGTAGAGAGATAGTAGTAGCGCCCGCGATCGATCGTACGTAGCTACGTACGATCTACGTACGTACGTACGTACGTACGTACGT"
	fmt.Println(Complement(dna))        // should print "TGGCTA"
	fmt.Println(Reverse(dna))           // should print "TAGCCA"
	fmt.Println(ReverseComplement(dna)) // should print "ATCGGT"
}

// ReverseComplement takes as input a string pattern of DNA symbols; it returns the reverse complementary strand's pattern of the string
func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

// Complement takes as input a string pattern of DNA symbols; it returns the string formed by complemnting each position of the input string ('A' -> 'T', 'C' -> 'G')
func Complement(dna string) string {
	result := make([]byte, len(dna)) // or result := ""
	for i, symbol := range dna {
		switch symbol {
		case 'A':
			result[i] = 'T'
		case 'C':
			result[i] = 'G'
		case 'G':
			result[i] = 'C'
		case 'T':
			result[i] = 'A'
		default:
			panic("Invalid symbol given to Compliment().")
		}
	}

	return string(result)
}

// Reverse takes as input a string pattern. It returns the string formed by reversing the positions of all symbools in pattern.

func Reverse(pattern string) string {
	rev := make([]byte, len(pattern))

	n := len(pattern)

	for i := range pattern {
		rev[i] = pattern[n-1-i]
	}

	return string(rev)
}
