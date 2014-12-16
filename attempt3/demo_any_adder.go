package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt3/evolver"
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
	"time"
)

const DIGITS = 2
const MAXAGE = 3000

var bestOrganism *evolver.Organism
var arena *evolver.Arena
var highestHealth float64 = -1000.0

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	organism := evolver.NewOrganism()
	for i := 0; i < DIGITS * 3; i++ {
		organism.AddPermanent(nnn.NewOrNeuron())
	}
	arena = evolver.NewArena(RunAddCase, Birth, Death, MAXAGE, 100, organism)
	arena.Wait()
	if bestOrganism == nil {
		fmt.Println("Did not get a good organism.")
	} else {
		fmt.Println("Got organism:", bestOrganism, "value",
			bestOrganism.Health().Value())
	}
}

func Birth(o *evolver.Organism) {
	o.UserInfo = map[string]int64{}
}

func Death(o *evolver.Organism) {
}

func RunAddCase(o *evolver.Organism) {
	if o.Health().Cycles > MAXAGE {
		if o.Health().Value() > 0.0 {
			bestOrganism = o
			arena.Stop()
			return
		} else if o.Health().Value() > highestHealth {
			highestHealth = o.Health().Value()
			fmt.Println("New highest health:", highestHealth, "-- successes",
				o.UserInfo)
		}
	}
	in1 := rand.Intn(1 << DIGITS)
	in2 := rand.Intn(1 << DIGITS)
	out := (in1 + in2) % (1 << DIGITS)
	success := RunCase(o, in1, in2, out)
	m := o.UserInfo.(map[string]int64)
	str := fmt.Sprintf("%d+%d", in1, in2)
	if success {
		m[str]++
	} else {
		m[str]--
	}
}

func RunCase(o *evolver.Organism, a, b, out int) bool {
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
	handleEnd := func() bool {
		// Compare the circuit's output to the given input
		pain := -0.01
		for i := 0; i < DIGITS; i++ {
			shouldFire := 0 != (out & (1 << uint(i)))
			if o.Get(i + 2 * DIGITS).Firing() != shouldFire {
				pain += 0.4
			}
		}
		o.Pain(pain)
		return pain < 0.0
	}
	for i := 0; i < 5; i++ {
		o.Cycle()
		for j := 0; j < DIGITS; j++ {
			if o.Get(j + 2 * DIGITS).Firing() {
				return handleEnd()
			}
		}
	}
	return handleEnd()
}
