package nnn

const (
	NEURON_XOR = iota
	NEURON_AND = iota
	NEURON_OR  = iota
)

type Neuron struct {
	inputs   []*Link
	outputs  []*Link
	firing   bool
	willFire bool
	function int
	lifetime Lifetime
}

func NewNeuron(function int) *Neuron {
	return &Neuron{[]*Link{}, []*Link{}, false, false, function, NewLifetime()}
}

func NewOrNeuron() *Neuron {
	return NewNeuron(NEURON_OR)
}

func NewAndNeuron() *Neuron {
	return NewNeuron(NEURON_AND)
}

func NewXorNeuron() *Neuron {
	return NewNeuron(NEURON_XOR)
}

func (self *Neuron) InputCount() int {
	count := 0
	for _, link := range self.inputs {
		if link.output.firing {
			count++
		}
	}
	return count
}

func (self *Neuron) NextCycle() bool {
	switch self.function {
	case NEURON_XOR:
		return self.InputCount()%2 != 0
	case NEURON_AND:
		for _, link := range self.inputs {
			if !link.output.firing {
				return false
			}
		}
		return true
	case NEURON_OR:
		for _, link := range self.inputs {
			if link.output.firing {
				return true
			}
		}
		return false
	}
	return false
}

func (self *Neuron) Fire() {
	self.firing = true
}

func (self *Neuron) Inhibit() {
	self.firing = false
}

func (self *Neuron) RemoveLinks() {
	for len(self.inputs) > 0 {
		self.inputs[0].Remove()
	}
	for len(self.outputs) > 0 {
		self.outputs[0].Remove()
	}
}
