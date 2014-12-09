package main

import (
	"github.com/unixpickle/AI-attempts/attempt1/nnn"
	"fmt"
	"math/rand"
	"time"
)

const MAX_NEURONS = 15

var isSet bool = false

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	network := nnn.NewNetwork()
	input0 := nnn.NewOrNeuron()
	input1 := nnn.NewOrNeuron()
	input2 := nnn.NewOrNeuron()
	output := nnn.NewOrNeuron()
	network.AddNeuron(input0) // set
	network.AddNeuron(input1) // value for set
	network.AddNeuron(input2) // get
	network.AddNeuron(output) // output
	streak := 0
	for {
		result := RunNetworkRandom(network)
		if result {
			streak++
			if streak == 1000 {
				break
			}
		} else {
			streak = 0
		}
	}
	fmt.Println("found circuit after", streak, "iterations and",
		network.Time, "cycles.")
	fmt.Println("circuit:", network.Neurons)
}

func RunNetworkRandom(network *nnn.Network) bool {
	if rand.Intn(4) == 0 {
		// Perform a set operation
		RunSet(network, rand.Intn(2) != 0)
		return true
	} else {
		// Perform a get operation
		return RunGet(network)
	}
}

func RunSet(network *nnn.Network, flag bool) {
	// Perform the set operation
	isSet = flag
	network.Neurons[0].Firing = true
	network.Neurons[1].Firing = flag
	for i := 0; i < 5; i++ {
		nnn.PrunePain(network)
		if len(network.Neurons) < MAX_NEURONS {
			nnn.Evolve(network, nnn.Recentness(network))
		}
		network.Cycle()
	}
}

func RunGet(network *nnn.Network) bool {
	// Perform the get operation
	network.Neurons[2].Firing = true
	for i := 0; i < 5; i++ {
		nnn.PrunePain(network)
		if len(network.Neurons) < MAX_NEURONS {
			nnn.Evolve(network, nnn.Recentness(network))
		}
		network.Cycle()
		if network.Neurons[3].Firing {
			if isSet {
				nnn.AddPain(network, -0.1)
				return true
			} else {
				nnn.AddPain(network, 0.5)
				return false
			}
		}
	}
	// We got no result, so now we make sure that was correct.
	if isSet {
		nnn.AddPain(network, 0.5)
		return false
	} else {
		nnn.AddPain(network, -0.1)
		return true
	}
}

func RunNetwork(network *nnn.Network, f0 bool, f1 bool, output bool) bool {
	network.Neurons[0].Firing = f0
	network.Neurons[1].Firing = f1
	for i := 0; i < 5; i++ {
		nnn.PrunePain(network)
		if len(network.Neurons) < 20 {
			nnn.Evolve(network, nnn.Recentness(network))
		}
		network.Cycle()
		if network.Neurons[2].Firing {
			if !output {
				nnn.AddPain(network, 1.0)
				return false
			} else {
				nnn.AddPain(network, -0.1)
				return true
			}
		}
	}
	if output {
		nnn.AddPain(network, 1.0)
		return false
	}
	return true
}
