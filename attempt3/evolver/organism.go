package evolver

import "github.com/unixpickle/AI-attempts/attempt2/nnn"

type Organism struct {
	*nnn.Network
	health *Health
}

func NewOrganism() *Organism {
	return &Organism{nnn.NewNetwork(), NewHealth()}
}

func (o *Organism) Add(n *nnn.Neuron) {
	n.UserInfo = false
	o.Network.Add(n)
}

func (o *Organism) AddPermanent(n *nnn.Neuron) {
	n.UserInfo = true
	o.Network.Add(n)
}

func (o *Organism) Clone() *Organism {
	return &Organism{o.Network.Clone(), o.health.Clone()}
}

func (o *Organism) Cycle() {
	o.Network.Cycle()
	o.health.Cycle()
}

func (o *Organism) Health() *Health {
	return o.health.Clone()
}

func (o *Organism) Pain(value float64) {
	o.health.ApplyPain(value)
}
