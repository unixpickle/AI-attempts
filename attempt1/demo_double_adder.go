package main

import (
	"math/rand"
	"fmt"
	"./nnn/src"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	network := nnn.NewNetwork()
	a0 := nnn.NewOrNeuron()
	a1 := nnn.NewOrNeuron()
	b0 := nnn.NewOrNeuron()
	b1 := nnn.NewOrNeuron()
	c0 := nnn.NewOrNeuron()
	c1 := nnn.NewOrNeuron()
	network.AddNeuron(a0)
	network.AddNeuron(a1)
	network.AddNeuron(b0)
	network.AddNeuron(b1)
	network.AddNeuron(c0)
	network.AddNeuron(c1)
	streak := 0
	for {
		result := RunNetworkRandom(network)
		fmt.Println("result:", result)
		if result {
			streak++
			if streak == 100 {
				break
			}
		} else {
			streak = 0
		}
	}
}

func RunNetworkRandom(network *nnn.Network) bool {
	a0 := rand.Intn(2) != 0
	a1 := rand.Intn(2) != 0
	b0 := rand.Intn(2) != 0
	b1 := rand.Intn(2) != 0
	c0 := (a0 && !b0) || (b0 && !a0)
	carry := a0 && b0
	c1 := (a1 && !b1) || (b1 && !a1)
	if carry {
		c1 = !c1
	}
	return RunNetwork(network, []bool{a0, a1, b0, b1, c0, c1})
}

func RunNetwork(network *nnn.Network, values []bool) bool {
	for i := 0; i < 4; i++ {
		network.Neurons[i].Firing = values[i]
	}
	// Give the circuit 10 clockcycles to do it
	for i := 0; i < 10; i++ {
		nnn.Prune(network)
		// Allow up to 20 neurons
		if len(network.Neurons) < 20 {
			nnn.Evolve(network, nnn.Recentness(network))
		}
		network.Cycle()
		if network.Neurons[4].Firing || network.Neurons[5].Firing {
			// Compare the circuit's output to the given input
			pain := -0.02
			if network.Neurons[4].Firing != values[4] {
				pain += 0.05
			}
			if network.Neurons[5].Firing != values[5] {
				pain += 0.05
			}
			nnn.AddPain(network, pain)
			if pain > 0 {
				return false
			} else {
				return true
			}
		}
	}
	// The circuit gave no output
	pain := -0.02
	if values[4] {
		pain += 0.05
	}
	if values[5] {
		pain += 0.05
	}
	nnn.AddPain(network, pain)
	if pain > 0 {
		return false
	} else {
		return true
	}
}
