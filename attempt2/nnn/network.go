package nnn

type Network struct {
	neurons []*Neuron
}

func NewNetwork() *Network {
	return &Network{[]*Neuron{}}
}

func (n *Network) Add(neuron *Neuron) {
	if neuron.network != nil {
		neuron.Remove()
	}
	neuron.network = n
	n.neurons = append(n.neurons, neuron)
}

func (n *Network) Cycle() {
	for _, neuron := range n.neurons {
		neuron.willFire = neuron.cycle()
	}
	for _, neuron := range n.neurons {
		neuron.firing = neuron.willFire
	}
}

func (n *Network) Get(idx int) *Neuron {
	return n.neurons[idx]
}

func (n *Network) Len() int {
	return len(n.neurons)
}
