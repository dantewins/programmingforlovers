package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	fmt.Println("Finding primes!")

	n := 1000000
	start := time.Now()
	fmt.Println(TrivialPrimeFinder(n))
	elapsed := time.Since(start)
	log.Printf("TrivialPrimeFinder took %s", elapsed)

	start2 := time.Now()
	fmt.Println(SieveOfEratosthenes(n))
	elapsed2 := time.Since(start2)
	log.Printf("TrivialPrimeFinder took %s", elapsed2)

	//fmt.Println(ListPrimes(n))
}

// TrivalPrimeFinder as input an integer n and returns a slice of boolean variables storing the primarlity of each nonnegative integer up to and including n.

func TrivialPrimeFinder(n int) []bool {
	primeBooleans := make([]bool, n+1)

	for p := 2; p <= n; p++ {
		primeBooleans[p] = IsPrime(p)
	}

	return primeBooleans
}

// IsPrime takes as input an integer p and returns true if p is prime and false otherwise.

func IsPrime(p int) bool {
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		// is k a divisor of p? If so, return false
		if p%k == 0 {
			return false // composite
		}
	}

	// surviving all of these checks means that p is prime
	return true
}

func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)

	// everyone starts as false by default now, set everything from 2 onward equal to true
	for p := 2; p <= n; p++ {
		primeBooleans[p] = true
	}

	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		// is prime? If so, cross of its multiples

		if primeBooleans[p] {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}
	}

	return primeBooleans
}

// CrossOffMultiples takes as input a slice of boolean variables primeBooleans and an integer p and returns a slice of boolean variables where all multiples of p have been set to false.
func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1

	for k := 2 * p; k <= n; k += p {
		primeBooleans[k] = false
	}

	return primeBooleans
}

// ListPrimes all prime numbers up to and (possibly) including n.

func ListPrimes(n int) []int {
	// first, create a slice of length 0]
	primeList := make([]int, 0)

	primeBooleans := SieveOfEratosthenes(n)

	for p := range primeBooleans {
		if primeBooleans[p] {
			// append p to our list
			primeList = append(primeList, p)
		}
	}

	return primeList
}
