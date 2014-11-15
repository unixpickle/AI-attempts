package nnn

import "math/rand"

func AddPain(network *Network, amount float64) {
	for _, neuron := range network.Neurons {
		for _, link := range neuron.Outputs {
			elapsed := 1 + network.Time - link.Life.LastUsed
			recentness := 1.0 / float64(elapsed)
			link.NetPain += recentness * amount
		}
	}
}

func PrunePain(network *Network) {
	Prune(network, func (link *Link) bool {
		return rand.Float64() > link.NetPain
	})
}
