package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt3/evolver"
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
	"time"
)

var bestOrganism *evolver.Organism
var arena *evolver.Arena

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	organism := evolver.NewOrganism()
	input0 := nnn.NewOrNeuron()
	input1 := nnn.NewOrNeuron()
	output0 := nnn.NewOrNeuron()
	output1 := nnn.NewOrNeuron()
	organism.AddPermanent(input0)
	organism.AddPermanent(input1)
	organism.AddPermanent(output0)
	organism.AddPermanent(output1)
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
	a := rand.Intn(2) != 0
	b := rand.Intn(2) != 0
	c0 := (a || b) && !(a && b)
	c1 := a && b
	RunCase(o, a, b, c0, c1)
}

func RunCase(o *evolver.Organism, a, b, c0, c1 bool) {
	if a {
		o.Get(0).Fire()
	} else {
		o.Get(0).Inhibit()
	}
	if b {
		o.Get(1).Fire()
	} else {
		o.Get(1).Inhibit()
	}
	handleEnd := func() {
		// Compare the circuit's output to the given input
		pain := -0.01
		if o.Get(2).Firing() != c0 {
			pain += 0.6
		}
		if o.Get(3).Firing() != c1 {
			pain += 0.4
		}
		o.Pain(pain)
	}
	for i := 0; i < 5; i++ {
		o.Cycle()
		if o.Get(2).Firing() || o.Get(3).Firing() {
			handleEnd()
			return
		}
	}
	handleEnd()
}
