package solve

import (
	"github.com/unixpickle/AI-attempts/attempt4/brain"
	"github.com/unixpickle/AI-attempts/attempt4/evolution"
)

type Organism struct {
	brain     brain.Brain
	box       BlackBox
	right     int
	wrong     int
	adulthood int
	minTime   int
	maxTime   int
	lastState interface{}
	report    func(o *Organism)
}

func (o *Organism) Adult() bool {
	return o.right+o.wrong >= o.adulthood
}

func (o *Organism) Die() {
}

func (o *Organism) Fitness() float64 {
	if o.wrong == 0 {
		return float64(o.right) + 1
	}
	return float64(o.right) / float64(o.wrong)
}

func (o *Organism) Reproduce() evolution.Organism {
	newNet := brain.Mutate(o.brain.Network())
	newBrain := brain.NewBrain(newNet)
	return &Organism{newBrain, o.box, 0, 0, o.adulthood, o.minTime, o.maxTime,
		o.box.Initial(), o.report}
}

func (o *Organism) Step() {
	var question []bool
	question, o.lastState = o.box.Ask(o.lastState)
	copy(o.brain.Activity(), question)

	// Run the brain until it comes up with an answer or times out.
	for i := 0; i < o.maxTime; i++ {
		if i >= o.minTime {
			if o.doAnswer(true) {
				return
			}
		}
		o.brain.Cycle()
	}

	o.doAnswer(false)
}

func (o *Organism) doAnswer(check bool) bool {
	ql := o.box.QuestionLen()
	answer := o.brain.Activity()[ql : ql+o.box.AnswerLen()]

	// If check is false, then we assume the organism produced an answer.
	if check {
		hasAnswer := false
		for _, f := range answer {
			if f {
				hasAnswer = true
				break
			}
		}
		if !hasAnswer {
			return false
		}
	}

	var right bool
	right, o.lastState = o.box.Check(answer, o.lastState)
	if right {
		o.right++
		if o.wrong == 0 && o.Adult() {
			o.report(o)
		}
	} else {
		o.wrong++
	}
	return true
}
