package nnn

import "math/rand"

type WeightFunc func(*Neuron) float64

type weightedNeuron struct {
	weight      float64
	priorWeight float64
	neuron      *Neuron
}

// Recentness returns a WeightFunc which weights neurons based on how recently
// they fired, favoring more recent ones.
func Recentness(network *Network) WeightFunc {
	return func(neuron *Neuron) float64 {
		delay := network.Time - neuron.Life.LastUsed
		return 1.0 / float64(delay+1)
	}
}

// Choose a given number of neurons from the network using a probability
// weighting function.
func WeightedChoose(network *Network, count int, weight WeightFunc) []*Neuron {
	if count > len(network.Neurons) {
		panic("Too many random neurons requested")
	}
	// Get the initial weighted list
	list := buildWeightedList(network, weight)
	// Choose random elements from the list
	result := make([]*Neuron, 0)
	for i := 0; i < count; i++ {
		var neuron *Neuron
		neuron, list = pickNeuron(list)
		result = append(result, neuron)
	}
	return result
}

func buildWeightedList(network *Network, weight WeightFunc) []weightedNeuron {
	// Generate a list of weighted neurons which we can use for randomness.
	weighted := make([]weightedNeuron, 0)
	var prior float64 = 0.0
	for _, neuron := range network.Neurons {
		el := weightedNeuron{weight(neuron), prior, neuron}
		prior += el.weight
		weighted = append(weighted, el)
	}
	return weighted
}

func pickNeuron(list []weightedNeuron) (*Neuron, []weightedNeuron) {
	count := len(list)
	if count == 0 {
		panic("Picking neuron from empty list")
	}
	// Compute the total amount of weight in the list.
	lastElement := list[count-1]
	maximum := lastElement.priorWeight + lastElement.weight
	// Generate a random number between 0 and maximum.
	random := rand.Float64() * maximum
	// Find the element in the list that we picked
	for i, wn := range list {
		if wn.priorWeight+wn.weight >= random {
			list[i] = list[count-1]

			// Update priorWeight fields after the deletion
			priorWeight := wn.priorWeight
			for j := i; j < count-1; j++ {
				list[j].priorWeight = priorWeight
				priorWeight += list[j].weight
			}

			return wn.neuron, list[0 : count-1]
		}
	}
	// This will only occur in weird rounding situations, or never at all.
	return list[count-1].neuron, list[0 : count-1]
}
