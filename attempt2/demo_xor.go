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
	output := nnn.NewOrNeuron()
	organism.Add(input0)
	organism.Add(input1)
	organism.Add(output)
	organism.KeepAt(0)
	organism.KeepAt(1)
	organism.KeepAt(2)
	arena := evolver.NewArena(0.99, 0.01, 1000, 50, RunXorCase, organism)
	o := arena.Wait()
	fmt.Println("Got organism of age", o.Age())
}

func RunXorCase(o *evolver.Organism) {
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
				return
			} else {
				o.Pain(-0.1)
				return
			}
		}
	}
	if output {
		o.Pain(1.0)
	} else {
		o.Pain(-0.1)
	}
}
