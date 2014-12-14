package evolver

import (
	"math/rand"
	"runtime"
	"sync"
)

type OrganismFunc func(o *Organism)

type Arena struct {
	fitness       *Fitness
	adultAge      uint64
	maxPopulation int
	loopFunc      OrganismFunc
	deathFunc     OrganismFunc
	birthFunc     OrganismFunc
	
	mutex        sync.RWMutex
	waitGroup    sync.WaitGroup
	stopped      bool
	population   int
}

func NewArena(loopFunc, birthFunc, deathFunc OrganismFunc,
	          adult uint64, maxPopulation int, o *Organism) *Arena {
	a := &Arena{NewFitness(maxPopulation), adult, maxPopulation,
		loopFunc, deathFunc, birthFunc, sync.RWMutex{},
		sync.WaitGroup{}, false, 1}
	a.waitGroup.Add(1)
	go a.organismLoop(o)
	return a
}

func (a *Arena) Population() int {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.population
}

func (a *Arena) Wait() {
	a.waitGroup.Wait()
}

func (a *Arena) Stop() {
	a.mutex.Lock()
	a.stopped = true
	a.mutex.Unlock()
}

func (a *Arena) organismLoop(o *Organism) {
	defer a.waitGroup.Done()
	a.mutex.Lock()
	identifier := a.fitness.Alloc()
	a.mutex.Unlock()
	
	a.birthFunc(o)
	for {
		// Perform a single lifecycle
		a.loopFunc(o)
		a.mutex.Lock()
		keep := a.lifecycle(o, identifier)
		if !keep {
			a.deathFunc(o)
			a.population--
			a.fitness.Free(identifier)
			a.mutex.Unlock()
			return
		}
		a.mutex.Unlock()
		runtime.Gosched()
	}
}

func (a *Arena) lifecycle(o *Organism, id int) bool {
	if a.stopped {
		return false
	} else if o.health.Cycles < a.adultAge {
		return true
	}
	
	v := o.health.Value()
	percentile := a.fitness.Percentile(id, v)
	
	live, reproduce := a.organismDestination(percentile)
		
	if !live {
		return false
	} else if reproduce {
		dup := o.Reproduce()
		a.waitGroup.Add(1)
		a.population++
		go a.organismLoop(dup)
	}
	
	return true
}

func (a *Arena) organismDestination(val float64) (bool, bool) {
	if a.population == 1 {
		return true, true
	}
	if val < 0.5 {
		// No reproduction, but according to probability we might want to keep
		// the organism alive.
		return rand.Float64()/2.0 >= val, false
	} else {
		if a.population == a.maxPopulation {
			return true, false
		}
		return true, (0.5 + rand.Float64()/2.0) < val
	}
}
