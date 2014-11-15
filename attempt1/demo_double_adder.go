package main

import (
	"./nnn/src"
	"fmt"
	"math/rand"
	"time"
)

const MAX_NEURONS = 6

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	network := nnn.NewNetwork()
	a := nnn.NewOrNeuron()
	b := nnn.NewOrNeuron()
	c0 := nnn.NewOrNeuron()
	c1 := nnn.NewOrNeuron()
	network.AddNeuron(a)
	network.AddNeuron(b)
	network.AddNeuron(c0)
	network.AddNeuron(c1)
	streak := 0
	trues := 0
	falses := 0
	for {
		result := RunNetworkRandom(network)
		if result {
			trues++
			streak++
			if streak == 1000 {
				break
			}
		} else {
			falses++
			streak = 0
		}
		if (trues+falses)%1000 == 0 {
			fmt.Println("true/false ratio:",
				float64(trues)/float64(falses), "cycles =", network.Time)
		}
	}
	fmt.Println("found circuit after", trues+falses, "iterations and",
		network.Time, "cycles.")
	fmt.Println("circuit:", network.Neurons)
}

func RunNetworkRandom(network *nnn.Network) bool {
	a := rand.Intn(2) != 0
	b := rand.Intn(2) != 0
	c0 := (a || b) && !(a && b)
	c1 := a && b
	return RunNetwork(network, []bool{a, b, c0, c1})
}

func RunNetwork(network *nnn.Network, values []bool) bool {
	for i := 0; i < 2; i++ {
		network.Neurons[i].Firing = values[i]
	}
	handleEnd := func() bool {
		// Compare the circuit's output to the given input
		status := true
		pain := -0.001
		if network.Neurons[2].Firing != values[2] {
			status = false
			pain += 0.6
		}
		if network.Neurons[3].Firing != values[3] {
			status = false
			pain += 0.4
		}
		nnn.AddPain(network, pain)
		return status
	}
	for i := 0; i < 5; i++ {
		nnn.PrunePain(network)
		nnn.PruneUseless(network, 1000)
		for len(network.Neurons) < MAX_NEURONS {
			nnn.Evolve(network, nnn.Recentness(network))
		}
		network.Cycle()
		if network.Neurons[2].Firing || network.Neurons[3].Firing {
			return handleEnd()
		}
	}
	// The circuit gave no output
	return handleEnd()
}
