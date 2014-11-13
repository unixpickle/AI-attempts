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
}

func (self *Neuron) InputCount() int {
	count := 0
	for input := range self.inputs {
		if input.firing {
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
		for input := range self.inputs {
			if !input.firing {
				return false
			}
		}
		return true
	case NEURON_OR:
		for input := range self.inputs {
			if input.firing {
				return true
			}
		}
		return false
	}
}
