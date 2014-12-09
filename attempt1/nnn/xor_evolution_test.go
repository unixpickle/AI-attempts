package nnn

import (
	"math/rand"
	"testing"
)

func TestXorEvolution(t *testing.T) {
	network := NewNetwork()
	input0 := NewOrNeuron()
	input1 := NewOrNeuron()
	output := NewOrNeuron()
	network.AddNeuron(input0)
	network.AddNeuron(input1)
	network.AddNeuron(output)
	streak := 0
	for i := 0; i < 100000; i++ {
		result := runRandomXorCase(network)
		if result {
			streak++
			if streak == 100 {
				return
			}
		} else {
			streak = 0
		}
	}
	t.Error("Test timed out; circuit never evolved.")
}

func runRandomXorCase(network *Network) bool {
	flag0 := rand.Intn(2) != 0
	flag1 := rand.Intn(2) != 0
	output := (flag0 || flag1) && !(flag0 && flag1)
	return runXorCase(network, flag0, flag1, output)
}

func runXorCase(network *Network, f0 bool, f1 bool, output bool) bool {
	network.Neurons[0].Firing = f0
	network.Neurons[1].Firing = f1
	for i := 0; i < 5; i++ {
		PrunePain(network)
		if len(network.Neurons) < 10 {
			Evolve(network, Recentness(network))
		}
		network.Cycle()
		if network.Neurons[2].Firing {
			if !output {
				AddPain(network, 1.0)
				return false
			} else {
				AddPain(network, -0.1)
				return true
			}
		}
	}
	if output {
		AddPain(network, 1.0)
		return false
	}
	return true
}
