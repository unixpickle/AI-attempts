package nnn

type Network struct {
	Neurons    []*Neuron
	Time       CycleTime
}

func NewNetwork() *Network {
	return &Network{[]*Neuron{}, 0}
}

func (self *Network) AddNeuron(neuron *Neuron) {
	self.Neurons = append(self.Neurons, neuron)
}

func (self *Network) RemoveNeuron(neuron *Neuron) {
	for i, x := range self.Neurons {
		if x == neuron {
			self.Neurons[i] = self.Neurons[len(self.Neurons)-1]
			self.Neurons = self.Neurons[0 : len(self.Neurons)-1]
			break
		}
	}
}

func (self *Network) Cycle() {
	self.Time++
	for _, neuron := range self.Neurons {
		neuron.willFire = neuron.NextCycle()
	}
	for _, neuron := range self.Neurons {
		neuron.Firing = neuron.willFire
		if neuron.Firing {
			neuron.Life.LastUsed = self.Time
			for _, output := range neuron.Outputs {
				output.Life.LastUsed = self.Time
			}
		}
	}
}

func (self *Network) CountFiring() uint {
	var count uint = 0
	for _, x := range self.Neurons {
		if x.Firing {
			count++
		}
	}
	return count
}

func (self *Network) GetNeuron(idx int) *Neuron {
	return self.Neurons[idx]
}

func (self *Network) CountNeurons() int {
	return len(self.Neurons)
}
