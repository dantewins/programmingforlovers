package main

import "fmt"

func main() {
	fmt.Println("Array (and slices).")

	var list [6]int
	fmt.Println(list)

	i := 3

	list[0] = -8
	list[2*i-4] = 17

	list[len(list)-1] = 43

	// these commands lead to compiler errors
	// list[len(list)] = 7
	// list [-4] = 2

	fmt.Println(list)

	var a []int
	n := 4

	a = make([]int, n)

	a[0] = 6
	a[2] = -33

	fmt.Println(a)

	b := make([]int, n+2)
	fmt.Println(b)

	fmt.Println(FactorialArray(6))

	var c [6]int

	d := make([]int, 6)

	ChangeFirstElementArray(c)
	ChangeFirstElementSlice(d)

	fmt.Println(c)
	fmt.Println(d)
}

// variadic function take a varaible number if inputs

// MinIntegers takes as input an arbitrary number of integers and returns their minimum value
func MinIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: empty slice given to MinIntegerArray")
	}

	return MinIntegerArray(numbers)
}

// MinIntegerArray takes as input a slice of integers and returns the minimum value in that slice.

func MinIntegerArray(a []int) int {
	if len(a) == 0 {
		panic("Error: empty slice given to MinIntegerArray.")
	}

	min := a[0]

	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min
}

func ChangeFirstElementArray(a [6]int) {
	a[0] = 1
}

func ChangeFirstElementSlice(a []int) {
	a[0] = 1
}

// FactorialArray takes as an input as an integer n, and it returns an array of length n+1 whose k-th element is k!
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: negative input given to FactorialArray.")
	}

	fact := make([]int, n+1)

	fact[0] = 1

	for k := 1; k <= n; k++ {
		fact[k] = k * fact[k-1]
	}

	return fact
}
