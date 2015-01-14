package solve

import (
	"github.com/unixpickle/AI-attempts/attempt4/brain"
	"github.com/unixpickle/AI-attempts/attempt4/evolution"
)

func Solve(box BlackBox, sampleCount int) brain.Network {
	// Create a new network with inputs and outputs.
	network := brain.Network{}
	for i := 0; i < box.AnswerLen()+box.QuestionLen(); i++ {
		node := brain.Node{brain.NodeTypeXOR, true, []int{}}
		network = append(network, node)
	}

	// Create a starting point.
	var solution *Organism
	report := func(o *Organism) {
		solution = o
	}
	seed := &Organism{brain.NewBrain(network), box, 0, 0, sampleCount,
		3, 20, box.Initial(), report}

	arena := evolution.NewArena(seed, 100, 100)
	for solution == nil {
		arena.Step()
	}

	return solution.brain.Network()
}
