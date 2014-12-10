package evolver

import "github.com/unixpickle/AI-attempts/attempt2/nnn"

func NewLink(from *nnn.Neuron, to *nnn.Neuron) *nnn.Link {
	l := nnn.NewLink(from, to)
	l.UserInfo = NewHistory()
	return l
}
