package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps in Go.")

	var polls map[string]float64
	polls = make(map[string]float64)

	// in practice, polls := make(map[string]float64)

	polls["Pennsylvania"] = 0.517
	polls["Ohio"] = 0.488
	polls["Texas"] = 0.378
	polls["Florida"] = 0.5

	fmt.Println("the number of states in the map is", len(polls))

	// getting rid of Florida

	delete(polls, "Florida")

	fmt.Println("the number of states in the map is", len(polls))

	// array and slice literals
	/*
		dnaAlphabet := [4]byte{'A', 'C', 'G', 'T'}
		primes := []int{2, 3, 5, 7, 11}
	*/

	// map literals
	electoralVotes := map[string]uint{
		"Pennsylvania": 20,
		"Ohio":         18,
		"Texas":        38, // GO demands final comma for consistency
	}

	UpdateVotes(electoralVotes)
	fmt.Println("The number of electoral votes in Pennsylvania is", electoralVotes["Pennsylvania"])

	for state, votes := range electoralVotes {
		fmt.Println("The number of electoral votes in", state, "is", votes)
	}

	PrintMapAlphabetical(electoralVotes)
}

func PrintMapAlphabetical(dict map[string]uint) {
	// sort the keys of the map, then range over the sorted keys to print key-value pairs
	i := 0
	keys := make([]string, len(dict))

	for key := range dict {
		keys[i] = key
		i++
	}

	// sort the keys
	sort.Strings(keys)

	// let's range over keys and print the associated dictionary values
	for _, key := range keys {
		fmt.Println("The value associated with", key, "is", dict[key])
	}
}

func UpdateVotes(electoralVotes map[string]uint) {
	electoralVotes["Pennsylvania"] = 19
	electoralVotes["Ohio"] = 17
	electoralVotes["Texas"] = 40
}
