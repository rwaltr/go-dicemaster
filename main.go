package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(RollString("2d20"))
}

type die struct {
	size int
}

// Takes a single dienotation and returns a dicepool
func dienotationdigest(dienotation string) (results []die) {

	working := strings.TrimSpace(dienotation)
	working = strings.ToLower(working)
	stringmultiplier, stringsize, found := strings.Cut(working, "d")

	if !found {
		panic("I do not understand this notation")

	} else {

		muliplier, _ := strconv.Atoi(stringmultiplier)
		sides, _ := strconv.Atoi(stringsize)

		for i := 0; i < muliplier; i++ {
			newdie := die{
				size: sides,
			}
			results = append(results, newdie)
		}
		return results
	}
}

// Rolls a dice pool and shows their sums and rolls
func RollDicePool(dice []die) (sum int, rolls []int) {
	for i := 0; i < len(dice); i++ {
		results := RollDie(dice[i])
		sum = sum + results
		rolls = append(rolls, results)
	}
	return sum, rolls
}

// Takes a dienotation string and rolls it
func RollString(roll string) int {
	rollresults, _ := RollDicePool(dienotationdigest(roll))
	return rollresults
}

// RollDie takes a die and returns the results based on the current time
func RollDie(die die) (results int) {
	timesource := rand.NewSource(time.Now().UnixNano())
	rando := rand.New(timesource)
	return DirectRoleDie(die, rando)
}

// Rolls a Die but accepted a rand object pointer, useful for repetable generation
func DirectRoleDie(die die, random *rand.Rand) (results int) {
	return random.Intn(die.size) + 1
}
