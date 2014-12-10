package evolver

import "github.com/unixpickle/AI-attempts/attempt2/nnn"

type Organism struct {
	*nnn.Network
	history *History
	age     uint64
}

func NewOrganism() *Organism {
	return &Organism{nnn.NewNetwork(), NewHistory(), 0}
}

func (o *Organism) Age() uint64 {
	return o.age
}

func (o *Organism) Cycle() {
	// Update the timestamps on firing neurons and their outputs
	for i := 0; i < o.Len(); i++ {
		neuron := o.Get(i)
		if !neuron.Firing() {
			continue
		}
		if neuron.UserInfo == nil {
			neuron.UserInfo = NewHistory()
		}
		neuron.UserInfo.(*History).LastFired = o.Age()
		for _, link := range neuron.Outputs {
			if link.UserInfo == nil {
				link.UserInfo = NewHistory()
			}
			link.UserInfo.(*History).LastFired = o.Age()
		}
	}
	// Run a cycle
	o.Network.Cycle()
	o.history.LastFired = o.Age()
	o.age++
}

func (o *Organism) History() History {
	return *o.history
}

func (o *Organism) KeepAt(idx int) {
	neuron := o.Get(idx)
	if neuron.UserInfo == nil {
		neuron.UserInfo = NewHistory()
	}
	neuron.UserInfo.(*History).Permanent = true
}

func (o *Organism) Pain(value float64) {
	// Trigger pain using the history of each neuron and link
	for i := 0; i < o.Len(); i++ {
		neuron := o.Get(i)
		if neuron.UserInfo == nil {
			continue
		}
		neuron.UserInfo.(*History).ApplyPain(value, o.Age())
		for _, link := range neuron.Outputs {
			if link.UserInfo == nil {
				continue
			}
			link.UserInfo.(*History).LastFired = o.Age()
		}
	}
	// Trigger pain for the entire organism
	o.history.ApplyPain(value, o.Age())
}
