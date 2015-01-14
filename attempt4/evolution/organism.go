package evolution

// An organism can be put through time and have its fitness evaluated.
type Organism interface {
	Adult() bool
	Die()
	Fitness() float64
	Reproduce() Organism
	Step()
}
