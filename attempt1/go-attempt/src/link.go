package nnn

type Link struct {
	input  *Neuron
	output *Neuron
}

func NewLink(input *Neuron, output *Neuron) *Link {
	result := &Link{input, output}
	input.outputs = append(input.outputs, result)
	output.inputs = append(output.inputs, result)
	return result
}

func (self *Link) Remove() {
	removeFromList(&self.input.outputs)
	removeFromList(&self.output.inputs)
}

func (self *Link) removeFromList(list *[]*Link) {
	for i, x := range list {
		if x == self {
			list[i] = list[len(list)-1]
			*list = list[0 : len(list)-1]
		}
	}
}
