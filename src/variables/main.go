package main

import "fmt"

// Go won't read this line

/*
This is a multi-line comment here.
These are good for lengthier discussions/explanations of your code.
*/

func main() {
	fmt.Println("Variables and types.")

	var j int = 14            // default: 0
	var x float64 = -2.3      // default: 0.0
	var yo string = "Hi"      // default: ""
	var u uint = 14           // default: 0
	var symbol byte = 'H'     // default: 0?
	var statement bool = true //default: false

	// shorthand declarations avoid lengthy var j int statements
	i := 6
	hi := "Yo"
	k := 34
	y := 3.7
	secondStatement := true

	// let's see what values these variables have
	fmt.Println(j)
	fmt.Println(x)
	fmt.Println(yo)
	fmt.Println(u)
	fmt.Println(symbol)
	fmt.Println(statement)

	fmt.Println(i, hi, k, y, secondStatement)

	fmt.Println(2 * (i + 5) * k)
	fmt.Println(2*y - 3.16)

	fmt.Println(float64(k) * y) // k is still an integer, but we "cast" its value to be a float64 for the purpose of the multiplication

	fmt.Println(k / 14) // might expect 2.4285 but we see 2!

	// dividing two integers corresponds to integer divison (throw away the remainder)

	fmt.Println(float64(k) / 14.0)

	// not every type conversion will work ... things like bool(0) might be false in another language, but Go doesn't allow it

	var p int = -187
	var s uint = uint(p)
	fmt.Println(s)

	m := 9223372036854775807
	fmt.Println(m + 1)
}
