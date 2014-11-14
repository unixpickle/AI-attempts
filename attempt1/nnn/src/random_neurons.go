package nnn

import "math/rand"

type weightedNeuron struct {
	weight      float64
	priorWeight float64
	neuron      *Neuron
}

func RandomNeurons(network *Network, requestedCount int,
		weightFunc func(*Neuron) float64) []*Neuron {
    // Get the initial weighted list
	list := buildWeightedList(network, weightFunc)
	// Choose random elements from the list
	count := requestedCount
	if count > len(list) {
		count = len(list)
	}
	result := make([]*Neuron, count)
	for i := 0; i < count; i++ {
		var neuron *Neuron
		neuron, list = pickNeuron(list)
		result = append(result, neuron)
	}
	return result
}

func buildWeightedList(network *Network,
		weight func(*Neuron) float64) []weightedNeuron {
    // Generate a list of weighted neurons which we can use for randomness.
	weighted := make([]weightedNeuron, len(network.neurons))
	var prior float64 = 0.0
	for _, neuron := range network.neurons {
		el := weightedNeuron{weight(neuron), prior, neuron}
		prior += el.weight
		weighted = append(weighted, el)
	}
	return weighted
}

func pickNeuron(list []weightedNeuron) (*Neuron, []weightedNeuron) {
	if len(list) == 0 {
		return nil, list
	}
	// Compute the total amount of weight in the list.
	lastElement := list[len(list) - 1]
	maximum := lastElement.priorWeight + lastElement.weight
	// Generate a random number between 0 and maximum.
	random := rand.Float64() * maximum
	// Find the element in the list that we picked
	for i, wn := range list {
		if i == len(list) - 1 || wn.priorWeight + wn.weight >= random {
			list[i] = list[len(list) - 1]
			return wn.neuron, list[0 : len(list) - 1]
		}
	}
	return nil, list
}
