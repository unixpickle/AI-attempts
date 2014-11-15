package main

import (
	"./nnn/src"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	network := nnn.NewNetwork()
	input0 := nnn.NewOrNeuron()
	input1 := nnn.NewOrNeuron()
	output := nnn.NewOrNeuron()
	network.AddNeuron(input0)
	network.AddNeuron(input1)
	network.AddNeuron(output)
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
	flag0 := rand.Intn(2) != 0
	flag1 := rand.Intn(2) != 0
	output := (flag0 || flag1) && !(flag0 && flag1)
	return RunNetwork(network, flag0, flag1, output)
}

func RunNetwork(network *nnn.Network, f0 bool, f1 bool, output bool) bool {
	network.Neurons[0].Firing = f0
	network.Neurons[1].Firing = f1
	for i := 0; i < 5; i++ {
		nnn.Prune(network)
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
