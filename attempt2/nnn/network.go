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

func (n *Network) Clone() *Network {
	res := NewNetwork()
	// Duplicate neurons themselves
	for _, neuron := range n.neurons {
		cpy := NewNeuron(neuron.Function)
		cpy.UserInfo = neuron.UserInfo
		cpy.firing = neuron.firing
		res.Add(cpy)
	}
	// Duplicate links between neurons
	for _, neuron := range n.neurons {
		for _, link := range neuron.Inputs {
			from, to := n.indexOf(link.From, link.To)
			link := NewLink(res.neurons[from], res.neurons[to])
			link.UserInfo = link.UserInfo
		}
	}
	return res
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

func (n *Network) indexOf(n1 *Neuron, n2 *Neuron) (int, int) {
	res1 := -1
	res2 := -1
	for i, neuron := range n.neurons {
		if neuron == n1 {
			res1 = i
		}
		if neuron == n2 {
			res2 = i
		}
	}
	if res1 < 0 || res2 < 0 {
		panic("neuron not found")
	}
	return res1, res2
}
