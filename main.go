package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(RollString("5d20"))
}

type DicePool struct {
	Rollset []Rollset
}

// Roll and resolve a DicePool
func (dp *DicePool) DirectRoll(random *rand.Rand) (sum int) {
	for i := 0; i < len(dp.Rollset); i++ {
		s, _ := dp.Rollset[i].DirectRoll(random)
		sum = sum + s
	}
	return sum
}
func (dp *DicePool) Roll() (sum int) {
	timesource := rand.NewSource(time.Now().UnixNano())
	rando := rand.New(timesource)
	return dp.DirectRoll(rando)
}

type Rollset struct {
	dice     []Die
	modifier int    // Plus or minus an ammount
	keep     string // high,low or ""
}

// Resolves a Rollset with a seed
func (rs *Rollset) DirectRoll(random *rand.Rand) (sum int, rolls []int) {
	for i := 0; i < len(rs.dice); i++ {
		rolls = append(rolls, rs.dice[i].DirectRole(random))
	}
	for i := 0; i < len(rolls); i++ {
		sum = sum + rolls[i]
	}
  sum = sum + rs.modifier
	return sum, rolls
}

func (rs *Rollset) Roll() (sum int, rolls []int) {
	timesource := rand.NewSource(time.Now().UnixNano())
	rando := rand.New(timesource)
	return rs.DirectRoll(rando)
}

type Die struct {
	size int
}

// Rolls a die and returns the results based on the current time
func (d *Die) Roll() (result int) {
	timesource := rand.NewSource(time.Now().UnixNano())
	rando := rand.New(timesource)
	return d.DirectRole(rando)
}

// Rolls the die but accepts a rand object pointer, useful for repetable generation
func (d *Die) DirectRole(random *rand.Rand) (results int) {
	return random.Intn(d.size) + 1
}

// Takes a single dienotation and returns a Rollset
func SingleDieNotationToRollset(dienotation string) (results Rollset) {

	working := strings.TrimSpace(dienotation)
	working = strings.ToLower(working)
	stringmultiplier, stringsize, found := strings.Cut(working, "d")

	if !found {
		panic("I do not understand this notation")

	} else {

		muliplier, _ := strconv.Atoi(stringmultiplier)
		sides, _ := strconv.Atoi(stringsize)

		for i := 0; i < muliplier; i++ {
			newdie := Die{
				size: sides,
			}
			results.dice = append(results.dice, newdie)
		}
		return results
	}
}

// Takes a dienotation string and rolls it
func RollString(roll string) (results int, rolls []int) {
	working := SingleDieNotationToRollset(roll)
	return working.Roll()
}
