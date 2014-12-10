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
	Function int
	UserInfo interface{}
	firing   bool
	willFire bool
	network  *Network
}

func NewNeuron(function int) *Neuron {
	return &Neuron{[]*Link{}, []*Link{}, function, nil, false, false, nil}
}

func NewAndNeuron() *Neuron {
	return NewNeuron(NEURON_AND)
}

func NewOrNeuron() *Neuron {
	return NewNeuron(NEURON_OR)
}

func NewXorNeuron() *Neuron {
	return NewNeuron(NEURON_XOR)
}

func (n *Neuron) Fire() {
	n.firing = true
}

func (n *Neuron) Inhibit() {
	n.firing = false
}

func (n *Neuron) Firing() bool {
	return n.firing
}

func (n *Neuron) Network() *Network {
	return n.network
}

func (n *Neuron) Remove() {
	for len(n.Inputs) > 0 {
		n.Inputs[0].Remove()
	}
	for len(n.Outputs) > 0 {
		n.Outputs[0].Remove()
	}
	if n.network != nil {
		n.removeFromList(&n.network.neurons)
		n.network = nil
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
	if self.Firing() {
		firingStr = "true"
	} else {
		firingStr = "false"
	}
	return fmt.Sprintf("Neuron(%p){Function=%s, Firing=%s, Outputs=%s}", self,
		funcStr, firingStr, self.Outputs)
}

func (n *Neuron) countInputs() int {
	count := 0
	for _, link := range n.Inputs {
		if link.From.Firing() {
			count++
		}
	}
	return count
}

func (n *Neuron) cycle() bool {
	switch n.Function {
	case NEURON_XOR:
		return n.countInputs()%2 != 0
	case NEURON_AND:
		for _, link := range n.Inputs {
			if !link.From.Firing() {
				return false
			}
		}
		return true
	case NEURON_OR:
		for _, link := range n.Inputs {
			if link.From.Firing() {
				return true
			}
		}
		return false
	}
	return false
}

func (n *Neuron) removeFromList(list *[]*Neuron) {
	for i, x := range *list {
		if x == n {
			(*list)[i] = (*list)[len(*list)-1]
			*list = (*list)[0 : len(*list)-1]
			break
		}
	}
}
