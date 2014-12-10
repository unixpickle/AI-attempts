package evolver

import "sync"

type ArenaCycle func(o *Organism) bool

type Arena struct {
	goalValue     float64
	cycleWeight   float64
	adultAge      uint64
	maxPopulation uint64
	cycleFunc     ArenaCycle

	mutex        sync.RWMutex
	waitGroup    sync.WaitGroup
	resChan      chan *Organism
	done         bool
	averageValue float64
	population   uint64
	totalDeaths  uint64
}

func NewArena(goal float64, cycleWeight float64, adultAge uint64, maxPop uint64,
	cycleFunc ArenaCycle, first *Organism) *Arena {
	first.history.Permanent = true
	res := &Arena{goal, cycleWeight, adultAge, maxPop, cycleFunc,
		sync.RWMutex{}, sync.WaitGroup{}, make(chan *Organism, 1), false, 0.0,
		1, 0}
	go res.runOrganism(first)
	res.waitGroup.Add(1)
	return res
}

func (a *Arena) Cancel() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.done {
		a.done = true
		close(a.resChan)
	}
}

func (a *Arena) TotalDeaths() uint64 {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.totalDeaths
}

func (a *Arena) Wait() *Organism {
	o := <-a.resChan
	a.waitGroup.Wait()
	return o
}

func (a *Arena) lifeLogic(o *Organism) bool {
	// If everything is done, we die
	if a.done {
		return false
	}

	// Young people cannot die
	if o.Age() < a.adultAge {
		return true
	}

	// Get the value of this organism
	history := o.History()
	value := history.Value()

	// If this organism achieved the goal, die
	if value >= a.goalValue {
		a.resChan <- o
		close(a.resChan)
		a.done = true
		return false
	}
	
	// Modify the average
	a.averageValue = (a.averageValue + a.cycleWeight*value) /
		(1.0 + a.cycleWeight)

	// If we are valueable, we will most likely reproduce
	
	if a.shouldReproduce(value) {
		// Reproduce if there is no overpopulation
		if a.population < a.maxPopulation {
			a.population++
			a.waitGroup.Add(1)
			go a.runOrganism(o.Reproduce())
		}
	} else if a.population == a.maxPopulation && !o.history.Permanent {
		// Overpopulation means someone has to die, and we are a weak link.
		return false
	}
	return true
}

func (a *Arena) runOrganism(o *Organism) {
	for {
		// Run a cycle
		if !a.cycleFunc(o) {
			// They can kill an organism automatically
			a.mutex.Lock()
			a.population--
			a.totalDeaths++
			if a.population == 0 && !a.done {
				a.done = true
				close(a.resChan)
			}
			a.mutex.Unlock()
			a.waitGroup.Done()
			return
		}

		// Perform life logic
		a.mutex.Lock()
		keepAlive := a.lifeLogic(o)
		if !keepAlive {
			a.population--
			a.totalDeaths++
			if a.population == 0 && !a.done {
				a.done = true
				close(a.resChan)
			}
			a.mutex.Unlock()
			a.waitGroup.Done()
			return
		}
		a.mutex.Unlock()
	}
}

func (a *Arena) shouldReproduce(value float64) bool {
	if a.population == 1 {
		return true
	}
	return value >= a.averageValue
}
