package nnn

type Link struct {
	input  *Neuron
	output *Neuron
}

func NewLink(sender *Neuron, receiver *Neuron) *Link {
	result := &Link{receiver, sender}
	receiver.inputs = append(receiver.inputs, result)
	sender.outputs = append(sender.outputs, result)
	return result
}

func (self *Link) Remove() {
	self.removeFromList(&self.input.outputs)
	self.removeFromList(&self.output.inputs)
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
