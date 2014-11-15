package nnn

import "fmt"

func (self *Link) String() string {
	return fmt.Sprintf("Link(%p){Receiver=%p, Sender=%p, NetPain=%f}", self,
		self.Receiver, self.Sender, self.NetPain)
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
