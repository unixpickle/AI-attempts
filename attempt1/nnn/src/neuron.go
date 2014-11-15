package nnn

import "fmt"

const (
	NEURON_XOR = iota
	NEURON_AND = iota
	NEURON_OR  = iota
)

type Neuron struct {
	Inputs   []*Link
	Outputs  []*Link
	Firing   bool
	willFire bool
	Function int
	Life     Lifetime
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
	for _, link := range self.Inputs {
		if link.Output.Firing {
			count++
		}
	}
	return count
}

func (self *Neuron) NextCycle() bool {
	switch self.Function {
	case NEURON_XOR:
		return self.InputCount()%2 != 0
	case NEURON_AND:
		for _, link := range self.Inputs {
			if !link.Output.Firing {
				return false
			}
		}
		return true
	case NEURON_OR:
		for _, link := range self.Inputs {
			if link.Output.Firing {
				return true
			}
		}
		return false
	}
	return false
}

func (self *Neuron) Fire() {
	self.Firing = true
}

func (self *Neuron) Inhibit() {
	self.Firing = false
}

func (self *Neuron) RemoveLinks() {
	for len(self.Inputs) > 0 {
		self.Inputs[0].Remove()
	}
	for len(self.Outputs) > 0 {
		self.Outputs[0].Remove()
	}
}

func (self *Neuron) String() string {
	var funcStr string
	switch self.Function {
	case NEURON_OR:
		funcStr = "OR"
	case NEURON_AND:
		funcStr = "AND"
	case NEURON_XOR:
		funcStr = "XOR"
	}
	var firingStr string
	if self.Firing {
		firingStr = "true"
	} else {
		firingStr = "false"
	}
	return fmt.Sprintf("Neuron(%p){Function=%s, Firing=%s, Outputs=%s}", self,
	    funcStr, firingStr, self.Outputs)
}
