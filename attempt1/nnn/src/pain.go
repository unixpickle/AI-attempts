package nnn

func AddPain(network *Network, amount float64) {
	for _, neuron := range network.neurons {
		for _, link := range neuron.outputs {
			elapsed := 1 + network.time - link.lifetime.lastUsed
			recentness := 1.0 / float64(elapsed)
			link.netPain += recentness * amount
		}
	}
}

func Prune(network *Network) {
	for _, neuron := range network.neurons {
		for i := 0; i < len(neuron.outputs); i++ {
			// TODO: use a better pruning function here with randomness
			link := neuron.outputs[i]
			if link.netPain > 1 {
				link.Remove()
				i--
			}
		}
	}
	// TODO: see if I need these {}
	for removeUnlinked(network) {}
}

func removeUnlinked(network *Network) bool {
	removed := false
	for i := 0; i < len(network.neurons); i++ {
		neuron := network.neurons[i]
		if len(neuron.outputs) == 0 {
			removed = true
			network.RemoveNeuron(neuron)
			i--
		}
	}
	return removed
}
