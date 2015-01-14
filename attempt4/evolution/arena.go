package evolution

type Arena struct {
	adults   []Organism
	children []Organism
	maxPop   int
}

func NewArena(seed Organism, maxPop int) *Arena {
	return &Arena{[]Organism{}, []Organism{seed}, maxPop}
}

func (a *Arena) Step() {
	// Grow to adulthood in safety.
	for i, org := range a.children {
		org.Step()
		if ch.Adult() {
			a.children[i] = a.children[len(a.children)-1]
			a.children = a.children[0 : len(a.children)-1]
			a.adults = append(a.adults, ch)
		}
	}

	// Perform pregnancies.
	for i, org := range a.adults {
		// TODO: do reproduction here
	}
}
