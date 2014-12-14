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
	output := nnn.NewOrNeuron()
	organism.AddPermanent(input0) // set (flip bit)
	organism.AddPermanent(output) // output
	organism.UserInfo = false
	arena = evolver.NewArena(RunMemoryCase, Birth, Death, 2000, 100, organism)
	arena.Wait()
	if bestOrganism == nil {
		fmt.Println("Did not get a good organism.")
	} else {
		fmt.Println("Got organism:", bestOrganism, "value",
			bestOrganism.Health().Value())
	}
}

func Birth(o *evolver.Organism) {
	o.UserInfo = false
}

func Death(o *evolver.Organism) {
}

func RunMemoryCase(o *evolver.Organism) {
	if o.Health().Cycles > 2000 && o.Health().Value() > 0.0 {
		bestOrganism = o
		arena.Stop()
		return
	}
	// Set should be less likely than get
	if rand.Intn(3) == 0 {
		RunSet(o)
	} else {
		RunGet(o)
	}
}

func RunSet(o *evolver.Organism) {
	o.UserInfo = !o.UserInfo.(bool)
	o.Get(0).Fire()
	for i := 0; i < 5; i++ {
		o.Cycle()
	}
}

func RunGet(o *evolver.Organism) {
	// Wait for an output
	for i := 0; i < 5; i++ {
		o.Cycle()
		if o.Get(1).Firing() {
			if o.UserInfo.(bool) {
				o.Pain(-0.01)
				return
			} else {
				o.Pain(1.0)
				return
			}
		}
	}
	if !o.UserInfo.(bool) {
		o.Pain(-0.01)
	} else {
		o.Pain(1.0)
	}
}
