package evolver

import "sync"

type ArenaCycle func(o *Organism)

type Arena struct {
	goalValue     float64
	cycleWeight   float64
	adultAge      uint64
	maxPopulation uint64
	cycleFunc     ArenaCycle
	
	mutex         sync.RWMutex
	waitGroup     sync.WaitGroup
	resChan       chan *Organism
	done          bool
	averageValue  float64
	population    uint64
}

func NewArena(goal float64, cycleWeight float64, adultAge uint64, maxPop uint64,
	cycleFunc ArenaCycle, first *Organism) *Arena {
	res := &Arena{goal, cycleWeight, adultAge, maxPop, cycleFunc,
		sync.RWMutex{}, sync.WaitGroup{}, make(chan *Organism, 1), false, 0.0,
		1}
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
		return false
	}
	
	// If we are valueable, we will most likely reproduce
	if value >= a.averageValue {
		// Adjust the average value
		a.averageValue = (a.averageValue + a.cycleWeight*value) /
			(1.0 + a.cycleWeight)
		// Reproduce if there is no overpopulation
		if a.population < a.maxPopulation {
			a.population++
			a.waitGroup.Add(1)
			go a.runOrganism(o.Reproduce())
		}
	} else if a.population == a.maxPopulation {
		// Overpopulation means someone has to die, and we are a weak link.
		return false
	}
	return true
}

func (a *Arena) runOrganism(o *Organism) {
	for {
		// Run a cycle
		a.cycleFunc(o)
		
		// Perform life logic
		a.mutex.Lock()
		keepAlive := a.lifeLogic(o)
		if !keepAlive {
			a.population--
			a.mutex.Unlock()
			a.waitGroup.Done()
			return
		}
		a.mutex.Unlock()
	}
}
