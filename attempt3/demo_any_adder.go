package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt3/evolver"
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
	"time"
)

const DIGITS = 2

var bestOrganism *evolver.Organism
var arena *evolver.Arena

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	organism := evolver.NewOrganism()
	for i := 0; i < DIGITS * 3; i++ {
		organism.AddPermanent(nnn.NewOrNeuron())
	}
	arena = evolver.NewArena(RunAddCase, Birth, Death, 1000, 100, organism)
	arena.Wait()
	if bestOrganism == nil {
		fmt.Println("Did not get a good organism.")
	} else {
		fmt.Println("Got organism:", bestOrganism, "value",
			bestOrganism.Health().Value())
	}
}

func Birth(o *evolver.Organism) {
}

func Death(o *evolver.Organism) {
}

func RunAddCase(o *evolver.Organism) {
	if o.Health().Cycles > 1000 && o.Health().Value() > 0.0 {
		bestOrganism = o
		arena.Stop()
		return
	}
	in1 := rand.Intn(1 << DIGITS)
	in2 := rand.Intn(1 << DIGITS)
	out := (in1 + in2) % (1 << DIGITS)
	RunCase(o, in1, in2, out)
}

func RunCase(o *evolver.Organism, a, b, out int) {
	for i := 0; i < DIGITS; i++ {
		mask := 1 << uint(i)
		if 0 != (a & mask) {
			o.Get(i).Fire()
		} else {
			o.Get(i).Inhibit()
		}
		if 0 != (b & mask) {
			o.Get(i + DIGITS).Fire()
		} else {
			o.Get(i + DIGITS).Inhibit()
		}
	}
	handleEnd := func() {
		// Compare the circuit's output to the given input
		pain := -0.01
		for i := 0; i < DIGITS; i++ {
			shouldFire := 0 != (out & (1 << uint(i)))
			if o.Get(i + 2 * DIGITS).Firing() != shouldFire {
				pain += 0.4
			}
		}
		o.Pain(pain)
	}
	for i := 0; i < 5; i++ {
		o.Cycle()
		for j := 0; j < DIGITS; j++ {
			if o.Get(j + 2 * DIGITS).Firing() {
				handleEnd()
				return
			}
		}
	}
	handleEnd()
}
