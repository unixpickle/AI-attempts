package nnn

type Link struct {
	From *Neuron
	To   *Neuron
}

func NewLink(from *Neuron, to *Neuron) *Link {
	res := &Link{from, to}
	from.Outputs = append(from.Outputs, res)
	to.Inputs = append(to.Inputs, res)
	return res
}

func (l *Link) Remove() {
	if l.From != nil {
		l.removeFromList(&l.From.Outputs)
		l.From = nil
	}
	if l.To != nil {
		l.removeFromList(&l.To.Inputs)
		l.To = nil
	}
}

func (l *Link) removeFromList(list *[]*Link) {
	for i, x := range *list {
		if x == l {
			(*list)[i] = (*list)[len(*list)-1]
			*list = (*list)[0 : len(*list)-1]
			break
		}
	}
}
