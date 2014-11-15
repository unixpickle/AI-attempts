package nnn

func Prune(network *Network, keep func(*Link) bool) {
	for _, neuron := range network.Neurons {
		for i := 0; i < len(neuron.Outputs); i++ {
			link := neuron.Outputs[i]
			if link.Life.Permanent {
				continue
			}
			if !keep(link) {
				link.Remove()
				i--
			}
		}
	}
	for removeUnlinked(network) {
	}
}

func removeUnlinked(network *Network) bool {
	removed := false
	for i := 0; i < len(network.Neurons); i++ {
		neuron := network.Neurons[i]
		if neuron.Life.Permanent {
			continue
		}
		if len(neuron.Outputs) == 0 {
			removed = true
			network.RemoveNeuron(neuron)
			i--
		}
	}
	return removed
}
