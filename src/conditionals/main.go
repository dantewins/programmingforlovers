package main

import "fmt"

func main() {
	fmt.Println("The minimum of 3 and 4 is", Min2(3, 4))

	fmt.Println(WhichIsGreater(3, 5))
	fmt.Println(WhichIsGreater(3, 5))
}

// Min2 takes two integers as input and returns their minimum
func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// PositiveDifference takes as input two integers. It returns the absolute value of the difference between these integers.
func PositiveDifference(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return a - b
	} else {
		// can safely infer that b > a
		return b - a
	}
}

// SameSign takes as input two integers and it returns true if the two integers have the same sign, and false if they have different signs
func SameSign(x, y int) bool {
	return x*y >= 0
}

// WhichIsGreater compares two input integers and returns 1 if the first input is larger, -1 if the second input is larger, and 0 if they are equal
func WhichIsGreater(x, y int) int {
	if x == y {
		return 0
	} else if x > y {
		return 1
	} else {
		return -1
	}
}

// index of comparsion operators

/*
	> : more than
	< : less than
	>= : greater than or equal to
	<= : less than or equal to
	== : equal to
	!= : not equal to
	! : "not". For example, if x is Boolean, then !x is false when x is true and true when x is false
*/
