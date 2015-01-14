package evolution

import (
	"math"
	"math/rand"
	"sort"
)

type Arena struct {
	adults   []Organism
	children []Organism
	adultPop int
	childPop int
}

func NewArena(seed Organism, adultPop, childPop int) *Arena {
	return &Arena{[]Organism{}, []Organism{seed}, adultPop, childPop}
}

func (a *Arena) Step() {
	// Children grow to adulthood in safety.
	for i := 0; i < len(a.children); i++ {
		org := a.children[i]
		org.Step()
		if org.Adult() {
			a.children[i] = a.children[len(a.children)-1]
			a.children = a.children[0 : len(a.children)-1]
			a.adults = append(a.adults, org)
			i--
		}
	}

	// If there are no adults, we can't do much of anything.
	if len(a.adults) == 0 {
		return
	}

	// Sort adults in order of increasing fitness.
	sort.Sort(fitnessOrganismList(a.adults))

	// Kill adults in order to keep the population under the limit.
	for len(a.adults) > a.adultPop {
		idx := killIndex(len(a.adults))
		a.adults[idx].Die()
		copy(a.adults[idx:], a.adults[idx+1:])
		a.adults = a.adults[0 : len(a.adults)-1]
	}

	// Reproduce as much as possible.
	for len(a.children) < a.childPop {
		idx := birthIndex(len(a.adults))
		org := a.adults[idx]
		a.children = append(a.children, org.Reproduce())
	}
}

type fitnessOrganismList []Organism

func (f fitnessOrganismList) Len() int {
	return len(f)
}

func (f fitnessOrganismList) Less(i, j int) bool {
	return f[i].Fitness() < f[j].Fitness()
}

func (f fitnessOrganismList) Swap(i, j int) {
	x := f[i]
	f[i] = f[j]
	f[j] = x
}

func birthIndex(count int) int {
	if count == 1 {
		return 0
	}

	rawNum := rand.Float64()
	weighted := 1 - math.Pow(rawNum, 3)
	res := int(float64(count) * weighted)

	// TODO: I don't think this is necessary, but rounding is evil.
	if res < 0 {
		res = 0
	} else if res >= count {
		res = count - 1
	}

	return res
}

func killIndex(count int) int {
	if count == 1 {
		return 0
	}

	rawNum := rand.Float64()
	weighted := math.Pow(rawNum, 3)
	res := int(float64(count) * weighted)

	// TODO: I don't think this is necessary, but rounding is evil.
	if res < 0 {
		res = 0
	} else if res >= count {
		res = count - 1
	}

	return res
}
