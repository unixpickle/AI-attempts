package evolver

import (
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
)

func (o *Organism) Reproduce() *Organism {
	child := o.Clone()

	// Either perform an addition or a deletion.
	for {
		num := rand.Float64()
		if num < 0.4 {
			child.MutateAddNeuron()
		} else if num < 0.8 {
			child.MutateRemoveNeuron()
			child.MutatePruneNeurons()
		} else if num < 0.9 {
			child.MutateAddLink()
		} else {
			child.MutateRemoveLink()
			child.MutatePruneNeurons()
		}
		// 1/2^n are the odds that it will perform n+1 mutations.
		if rand.Intn(2) != 0 {
			break
		}
	}

	// Stop all neural signals and reset the organism's health
	child.health = NewHealth()
	for i := 0; i < child.Len(); i++ {
		child.Get(i).Inhibit()
	}
	
	return child
}

// MutateAddNeuron adds a random neuron to an organism and connects it to
// randomly chosen neurons.
func (o *Organism) MutateAddNeuron() *nnn.Neuron {
	// Generate and add a neuron with a random function
	neuron := nnn.NewNeuron(rand.Intn(3))
	o.Add(neuron)

	// Inputs (1 or 2)
	from1 := o.Get(rand.Intn(o.Len()))
	nnn.NewLink(from1, neuron)
	if rand.Intn(2) == 0 {
		// Two inputs
		from2 := o.Get(rand.Intn(o.Len()))
		nnn.NewLink(from2, neuron)
	}

	// Output
	dest := o.Get(rand.Intn(o.Len()))
	nnn.NewLink(neuron, dest)
	
	return neuron
}

// MutateRemoveNeuron randomly removes a neuron from a neural network.
func (o *Organism) MutateRemoveNeuron() {
	// Generate the full list of neurons
	neurons := make([]*nnn.Neuron, 0, o.Len())
	for i := 0; i < o.Len(); i++ {
		neuron := o.Get(i)
		if !neuron.UserInfo.(bool) {
			neurons = append(neurons, neuron)
		}
	}
	if len(neurons) == 0 {
		return
	}
	// Choose a random neuron and remove it
	toRemove := rand.Intn(len(neurons))
	neurons[toRemove].Remove()
}

// MutateAddLink adds a link between two random neurons.
func (o *Organism) MutateAddLink() *nnn.Link {
	n1 := o.Get(rand.Intn(o.Len()))
	n2 := o.Get(rand.Intn(o.Len()))
	return nnn.NewLink(n1, n2)
}

// MutateRemoveLink removes a random link from the network.
func (o *Organism) MutateRemoveLink() {
	// Generate a full list of links
	links := make([]*nnn.Link, 0)
	for i := 0; i < o.Len(); i++ {
		for _, link := range o.Get(i).Inputs {
			links = append(links, link)
		}
	}
	// Choose a random link and remove it
	if len(links) == 0 {
		return
	}
	toRemove := rand.Intn(len(links))
	links[toRemove].Remove()
}

// MutatePruneNeurons removes all the neurons in the organism which have no
// inputs and outputs (i.e. neurons which are fully disconnected).
// This will not remove neurons which are marked as permanent.
// Returns the number of neurons which were removed.
func (o *Organism) MutatePruneNeurons() int {
	// Keep removing unused neurons until nothing changes.
	removed := 0
	changed := true
	for changed {
		changed = false
		for i := 0; i < o.Len(); i++ {
			neuron := o.Get(i)
			if len(neuron.Inputs) == 0 || len(neuron.Outputs) == 0 {
				if !neuron.UserInfo.(bool) {
					neuron.Remove()
					changed = true
					removed++
				}
			}
		}
	}
	return removed
}
