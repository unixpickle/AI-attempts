package main

import (
	"fmt"
	"math/rand"
	"./nnn/src"
)

func main() {
	network := nnn.NewNetwork()
	input0 := nnn.NewOrNeuron()
	input1 := nnn.NewOrNeuron()
	output := nnn.NewOrNeuron()
	network.AddNeuron(input0)
	network.AddNeuron(input1)
	network.AddNeuron(output)
	for i := 0; i < 100; i++ {
		fmt.Println("result:", RunNetworkRandom(network))
	}
}

func RunNetworkRandom(network *nnn.Network) bool {
	flag0 := rand.Intn(2) != 0
	flag1 := rand.Intn(2) != 0
	output := (flag0 || flag1) && !(flag0 && flag1)
	return RunNetwork(network, flag0, flag1, output)
}

func RunNetwork(network *nnn.Network, f0 bool, f1 bool, output bool) bool {
	network.Neurons[0].Firing = f0
	network.Neurons[1].Firing = f1
	for i := 0; i < 5; i++ {
		network.Prune()
		if len(network.Neurons) < 4 {
			Evolve(network, Recentness(network))
		}
		network.Cycle()
		if network.Neurons[2].Firing {
			if !output {
				AddPain(network, 0.5)
				return false
			} else {
				AddPain(network, -0.1)
				return true
			}
		}
	}
	return !output
}
