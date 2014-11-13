package nnn

type Network struct {
	neurons []*Neuron
}

func NewNetwork() *Network {
	return &Network{[]*Neuron{}}
}

func (self *Network) AddNeuron(neuron *Neuron) {
	self.neurons = append(self.neurons, neuron)
}

func (self *Network) RemoveNeuron(neuron *Neuron) {
	for i, x := range self.neurons {
		if x == neuron {
			self.neurons[i] = self.neurons[len(self.neurons) - 1]
			self.neurons = self.neurons[0 : len(self.neurons) - 1]
			break
		}
	}
}

func (self *Network) Cycle() {
	for i := 0; i < len(self.neurons); i++ {
		self.neurons[i].willFire = self.neurons[i].NextCycle()
	}
	for i := 0; i < len(self.neurons); i++ {
		self.neurons[i].firing = self.neurons[i].willFire
	}
}
