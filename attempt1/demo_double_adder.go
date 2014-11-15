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
			if streak == 100 {
				break
			}
		} else {
			falses++
			streak = 0
		}
		fmt.Println("result:", result, "ratio:", float64(trues)/float64(falses))
	}
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
	// Give the circuit 10 clockcycles to do it
	for i := 0; i < 10; i++ {
		nnn.Prune(network)
		// Allow up to 20 neurons
		if len(network.Neurons) < 8 {
			nnn.Evolve(network, nnn.Recentness(network))
		}
		network.Cycle()
		if network.Neurons[2].Firing || network.Neurons[3].Firing {
			// Compare the circuit's output to the given input
			pain := -0.05
			if network.Neurons[2].Firing != values[2] {
				pain += 1.0
			}
			if network.Neurons[3].Firing != values[3] {
				pain += 1.0
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
	pain := -0.05
	if values[2] {
		pain += 1.0
	}
	if values[3] {
		pain += 1.0
	}
	nnn.AddPain(network, pain)
	if pain > 0 {
		return false
	} else {
		return true
	}
}
