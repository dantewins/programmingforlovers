package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Implementing two GCD algorithms.")

	fmt.Println(TrivialGCD(63, 42))

	fmt.Println(EuclidGCD(63, 42))

	x := 37820263334
	y := 27314793433

	start := time.Now()
	TrivialGCD(x, y)
	elapsed := time.Since(start)
	log.Printf("TrivialGCD took %s", elapsed)

	start2 := time.Now()
	EuclidGCD(x, y)
	elapsed2 := time.Since(start2)
	log.Printf("EuclidGCD took %s", elapsed2)
}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TrivialGCD(a, b int) int {
	d := 1

	m := Min2(a, b)
	for p := 1; p <= m; p++ {
		if (a%p == 0) && (b%p == 0) {
			d = p
		}
	}

	return d
}

func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
