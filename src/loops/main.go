/*
	Helpful pseudocode for this lesson

	Factorial(n)
		p ← 1
		i ← 1
		while i ≤ n
			p ← p ⋅ i
			i ← i + 1
		return p

	EuclidGCD(a, b)
		while a ≠ b
			if a > b
				a ← a - b
			else b ← a
		return a
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops!")

	fmt.Println(Factorial(5))
	fmt.Println(SumFirstNIntegers(3))
}

// AnotherFactorial takes as input an integer n and returns n!
func AnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	p := 1

	for i := 1; i <= n; i++ {
		// i := 1 is called the "initialization step"
		// i <= n is called the "condition"
		// i++ is called the "postcondition"

		p *= i

		// there is a hidden line of code: i++
	}

	return p
}

func YetAnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	p := 1

	for i := n; i >= 1; i-- {
		p *= i
	}

	return p
}

// SumEven takes as input as an integer k and returns the sum of all even positive integers up to and (possibly) including k.

func SumEven(k int) int {
	sum := 0

	// for ever integer j between 2 and k, add j to sum

	for j := 2; j <= k; j += 2 {
		// j must be even :)

		sum += j
	}

	return sum
}

// AnotherSum takes an integer n as input and returns the sume of the first n positive integers, using a for loop
func AnotherSum(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	sum := 0 // will store the final answer

	for k := 1; k <= n; k++ {
		sum += k
	}

	return sum
}

func GaussSum(n int) int {
	return n * (n + 1) / 2
}

// mathematical fact: n! = n * (n-1)!
// when n = 1: 1! = 1 * 0!
// so, 1 = 1 * 0! and therefore 1 = 0!

// Factorial takes as input an integer n and returns n! = n * (n - 1) * (n-2) * ... * 2 * 1
func Factorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	p := 1 // will store product
	i := 1 // serves as a counter

	// Go has no keyword "while" and uses "for" instead for all loops
	for i <= n {
		p *= i
		i++
	}

	return p
}

// SumFirstNIntegers takes as input an integer n and returns the sum of the first n integers positive integers

func SumFirstNIntegers(n int) int {
	if n < 0 {
		panic("Error: negative input given to SumFirstNIntegers().")
	}

	sum := 0 // will store the final answer
	i := 1   // represents the current integer being added to sum

	for i <= n {
		sum += i // equivalent to sum = sum + i
		// also have sum -= i, sum *= i, sum /= i
		i++ // there is also i-- for i = i - 1
	}

	return sum
}
