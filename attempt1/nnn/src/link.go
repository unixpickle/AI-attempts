package nnn

import "fmt"

type Link struct {
	Input    *Neuron
	Output   *Neuron
	Life     Lifetime
	NetPain  float64
}

func NewLink(sender *Neuron, receiver *Neuron) *Link {
	result := &Link{receiver, sender, NewLifetime(), 0.0}
	receiver.Inputs = append(receiver.Inputs, result)
	sender.Outputs = append(sender.Outputs, result)
	return result
}

func (self *Link) Remove() {
	self.removeFromList(&self.Input.Outputs)
	self.removeFromList(&self.Output.Inputs)
}

func (self *Link) removeFromList(list *[]*Link) {
	for i, x := range *list {
		if x == self {
			(*list)[i] = (*list)[len(*list)-1]
			*list = (*list)[0 : len(*list)-1]
			break
		}
	}
}

func (self *Link) String() string {
	return fmt.Sprintf("Link(%p){Input=%p, Output=%p, NetPain=%f}", self,
		self.Input, self.Output, self.NetPain)
}
