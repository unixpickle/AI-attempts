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
		return 1.0 / float64(delay + 1)
	}
}

// Generates a neuron with a random type.
// The new neuron will not be marked "permanent"
func RandomNeuron() *Neuron {
	// TODO: see if I should weight this more heavily towards OR neurons...
	var result *Neuron
	switch rand.Intn(3) {
	case 0:
		result = NewOrNeuron()
	case 1:
		result = NewAndNeuron()
	default:
		result = NewXorNeuron()
	}
	result.Life.Permanent = false
	return result
}

// Adds a random neuron to the network with one or two inputs and one output.
// Both the inputs and the output will be chosen using WeightedChoose with the
// specified linkWeight function. The created links will not be permanent.
func Evolve(network *Network, linkWeight WeightFunc) *Neuron {
	if len(network.Neurons) < 2 {
		return nil
	}
	
	neuron := RandomNeuron()
	if len(network.Neurons) == 2 || rand.Intn(2) == 0 {
		// One input, one output
		conns := WeightedChoose(network, 2, linkWeight)
		NewLink(conns[0], neuron).Life.Permanent = false
		NewLink(neuron, conns[1]).Life.Permanent = false
	} else {
		// Two inputs, one output
		conns := WeightedChoose(network, 3, linkWeight)
		NewLink(conns[0], neuron).Life.Permanent = false
		NewLink(conns[1], neuron).Life.Permanent = false
		NewLink(neuron, conns[2]).Life.Permanent = false
	}
	return neuron
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
	result := make([]*Neuron, count)
	for i := 0; i < count; i++ {
		var neuron *Neuron
		neuron, list = pickNeuron(list)
		result = append(result, neuron)
	}
	return result
}

func buildWeightedList(network *Network, weight WeightFunc) []weightedNeuron {
    // Generate a list of weighted neurons which we can use for randomness.
	weighted := make([]weightedNeuron, len(network.Neurons))
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
	lastElement := list[count - 1]
	maximum := lastElement.priorWeight + lastElement.weight
	// Generate a random number between 0 and maximum.
	random := rand.Float64() * maximum
	// Find the element in the list that we picked
	for i, wn := range list {
		if wn.priorWeight + wn.weight >= random {
			list[i] = list[count - 1]
			return wn.neuron, list[0 : count - 1]
		}
	}
	// This will only occur in weird rounding situations, or never at all.
	return list[count - 1].neuron, list[0 : count - 1]
}

