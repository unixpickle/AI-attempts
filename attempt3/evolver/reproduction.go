package evolver

import (
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
)

func (o *Organism) Reproduce() *Organism {
	child := o.Clone()

	// Either perform an addition or a deletion.
	for {
		if rand.Intn(2) == 0 {
			child.MutateAddNeuron()
		} else {
			child.MutateRemoveNeuron()
			child.MutatePruneNeurons()
		}
		// 1/3^n are the odds that it will perform n+1 mutations.
		if rand.Intn(3) != 0 {
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
	// Generate the full list of links
	neurons := make([]*nnn.Neuron, 0)
	for i := 0; i < o.Len(); i++ {
		neuron := o.Get(i)
		if !neuron.UserInfo.(bool) {
			neurons = append(neurons, neuron)
		}
	}
	if len(neurons) == 0 {
		return
	}
	// Choose a random link and remove it
	toRemove := rand.Intn(len(neurons))
	neurons[toRemove].Remove()
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
