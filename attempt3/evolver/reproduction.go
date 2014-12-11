package evolver

func (o *Organism) Reproduce() *Organism {
	child := o.Clone()

	// TODO: perform mutations on the child by randomly adding or removing
	// neurons or links.

	// Stop all neural signals and reset all histories
	child.health = NewHealth()
	for i := 0; i < child.Len(); i++ {
		child.Get(i).Inhibit()
	}
	
	return child
}
