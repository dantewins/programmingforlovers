package main

import "fmt"

func main() {
	fmt.Println("Functions!")

	x := 3
	n := SumTwoInts(x, x)
	fmt.Println(n)

	m := 17
	fmt.Println(AddOne(m))
	fmt.Println(m) // what gets printed?
}

func AddOne(k int) int {
	return k + 1
}

// SumTwoInts takes two as input two integers and returns their sum.
func SumTwoInts(a int, b int) int {
	return a + b
}

// DoubleAndDuplicate takes as input a float64. It returns two copies of that input variable.
func DoubleAndDuplicate(x float64) (float64, float64) {
	return x * 2.0, x * 2.0
}

// Pi takes no inputs and returns the value, the mathematical constant
func Pi() float64 {
	return 3.14
}

// PrintHi simply prints a message "Hi" to the console.
func PrintHi() {
	fmt.Println("Hi")
}
