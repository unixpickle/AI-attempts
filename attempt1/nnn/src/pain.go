package nnn

func AddPain(network *Network, amount float64) {
	for _, neuron := range network.Neurons {
		for _, link := range neuron.Outputs {
			elapsed := 1 + network.Time - link.Life.LastUsed
			recentness := 1.0 / float64(elapsed)
			link.NetPain += recentness * amount
		}
	}
}

func Prune(network *Network) {
	for _, neuron := range network.Neurons {
		for i := 0; i < len(neuron.Outputs); i++ {
			// TODO: use a better pruning function here with randomness
			link := neuron.Outputs[i]
			if link.NetPain > 1 && !link.Life.Permanent {
				link.Remove()
				i--
			}
		}
	}
	// TODO: see if I need these curly braces
	for removeUnlinked(network) {}
}

func removeUnlinked(network *Network) bool {
	removed := false
	for i := 0; i < len(network.Neurons); i++ {
		neuron := network.Neurons[i]
		if len(neuron.Outputs) == 0 && !neuron.Life.Permanent {
			removed = true
			network.RemoveNeuron(neuron)
			i--
		}
	}
	return removed
}
