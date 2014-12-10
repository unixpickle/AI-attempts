package evolver

import (
	"github.com/unixpickle/AI-attempts/attempt2/nnn"
	"math/rand"
)

func (o *Organism) Reproduce() *Organism {
	child := o.Clone()

	// Perform mutations on the child
	prune(child)
	addNeurons(child)

	// Stop all neural signals and reset all histories
	child.history = resetHistory(child.history)
	child.age = 0
	for i := 0; i < child.Len(); i++ {
		neuron := child.Get(i)
		neuron.Inhibit()
		neuron.UserInfo = resetHistory(neuron.UserInfo.(*History))
		for _, link := range neuron.Inputs {
			link.UserInfo = resetHistory(link.UserInfo.(*History))
		}
	}
	return child
}

func addNeurons(o *Organism) {
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
}

func isPermanent(n *nnn.Neuron) bool {
	if n.UserInfo == nil {
		return false
	}
	return n.UserInfo.(*History).Permanent
}

func prune(o *Organism) {
	// Remove links randomly as needed
	for i := 0; i < o.Len(); i++ {
		neuron := o.Get(i)
		removeLinksInNeuron(neuron)
	}

	// Keep removing unused neurons until nothing changes.
	// This is probably not really necessary.
	changed := true
	for changed {
		changed = false
		for i := 0; i < o.Len(); i++ {
			neuron := o.Get(i)
			if len(neuron.Inputs) == 0 || len(neuron.Outputs) == 0 {
				if !isPermanent(neuron) {
					neuron.Remove()
					changed = true
				}
			}
		}
	}
}

func removeLinksInNeuron(n *nnn.Neuron) {
	for i := 0; i < len(n.Inputs); i++ {
		link := n.Inputs[i]
		if link.UserInfo == nil {
			continue
		}
		if !link.UserInfo.(*History).RandomKeep() {
			link.Remove()
			i--
		}
	}
}

func resetHistory(h *History) *History {
	if h == nil || !h.Permanent {
		return nil
	}
	res := NewHistory()
	res.Permanent = true
	return res
}
