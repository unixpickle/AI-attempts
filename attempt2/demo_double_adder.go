package main

import (
	"fmt"
	"github.com/unixpickle/AI-attempts/attempt2/evolver"
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	organism := evolver.NewOrganism()
	input0 := nnn.NewOrNeuron()
	input1 := nnn.NewOrNeuron()
	output0 := nnn.NewOrNeuron()
	output1 := nnn.NewOrNeuron()
	organism.Add(input0)
	organism.Add(input1)
	organism.Add(output0)
	organism.Add(output1)
	organism.KeepAt(0)
	organism.KeepAt(1)
	organism.KeepAt(2)
	organism.KeepAt(3)
	arena := evolver.NewArena(0.99, 0.01, 100, 8, RunAddCase, organism)
	o := arena.Wait()
	if o == nil {
		fmt.Println("Evolution failed :(")
	} else {
		fmt.Println("Got organism of age", o.Age(), "; totalDeaths =",
			arena.TotalDeaths(), "; organism size =", o.Len(),
			"; network =", o.Network.String())
	}
}

func RunAddCase(o *evolver.Organism) bool {
	/*if o.Len() > 6 {
		return false
	}*/
	a := rand.Intn(2) != 0
	b := rand.Intn(2) != 0
	c0 := (a || b) && !(a && b)
	c1 := a && b
	RunCase(o, a, b, c0, c1)
	return true
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
		pain := -0.05
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
