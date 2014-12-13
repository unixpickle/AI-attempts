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
	output := nnn.NewOrNeuron()
	organism.AddPermanent(input0)
	organism.AddPermanent(input1)
	organism.AddPermanent(output)
	arena = evolver.NewArena(RunXorCase, Birth, Death, 1000, 100, organism)
	arena.Wait()
	if bestOrganism == nil {
		fmt.Println("Did not get a good organism.")
	} else {
		fmt.Println("Got organism:", bestOrganism, "value",
			bestOrganism.Health().Value())
	}
}

func Birth(o *evolver.Organism) {
	fmt.Println("size", o.Len())
}

func Death(o *evolver.Organism) {
}

func RunXorCase(o *evolver.Organism) {
	if o.Health().Cycles > 1000 && o.Health().Value() > 0.05 {
		bestOrganism = o
		arena.Stop()
		return
	}
	flag0 := rand.Intn(2) != 0
	flag1 := rand.Intn(2) != 0
	output := (flag0 || flag1) && !(flag0 && flag1)
	RunCase(o, flag0, flag1, output)
}

func RunCase(o *evolver.Organism, flag0, flag1, output bool) {
	if flag0 {
		o.Get(0).Fire()
	} else {
		o.Get(0).Inhibit()
	}
	if flag1 {
		o.Get(1).Fire()
	} else {
		o.Get(1).Inhibit()
	}
	for i := 0; i < 5; i++ {
		o.Cycle()
		if o.Get(2).Firing() {
			if !output {
				o.Pain(1.0)
			} else {
				o.Pain(-0.2)
			}
			return
		}
	}
	if output {
		o.Pain(1.0)
	} else {
		o.Pain(-0.2)
	}
}
