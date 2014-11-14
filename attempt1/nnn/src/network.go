package nnn

type Network struct {
	neurons    []*Neuron
	time       CycleTime
}

func NewNetwork() *Network {
	return &Network{[]*Neuron{}, 0}
}

func (self *Network) AddNeuron(neuron *Neuron) {
	self.neurons = append(self.neurons, neuron)
}

func (self *Network) RemoveNeuron(neuron *Neuron) {
	for i, x := range self.neurons {
		if x == neuron {
			self.neurons[i] = self.neurons[len(self.neurons)-1]
			self.neurons = self.neurons[0 : len(self.neurons)-1]
			break
		}
	}
}

func (self *Network) Cycle() {
	self.time++
	for _, neuron := range self.neurons {
		neuron.willFire = neuron.NextCycle()
	}
	for _, neuron := range self.neurons {
		neuron.firing = neuron.willFire
		if neuron.firing {
			neuron.lifetime.lastUsed = self.time
			for _, output := range neuron.outputs {
				output.lifetime.lastUsed = self.time
			}
		}
	}
}

func (self *Network) CountFiring() uint {
	var count uint = 0
	for _, x := range self.neurons {
		if x.firing {
			count++
		}
	}
	return count
}

func (self *Network) GetNeuron(idx int) *Neuron {
	return self.neurons[idx]
}

func (self *Network) CountNeurons() int {
	return len(self.neurons)
}
