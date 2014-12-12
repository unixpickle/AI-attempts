package evolver

import (
	"runtime"
	"sync"
)

type OrganismFunc func(o *Organism)

type Arena struct {
	average       *Averager
	adultAge      uint64
	maxPopulation uint64
	loopFunc      OrganismFunc
	deathFunc     OrganismFunc
	birthFunc     OrganismFunc
	
	mutex        sync.RWMutex
	waitGroup    sync.WaitGroup
	stopped      bool
	population   uint64
	totalDeaths  uint64
}

func NewArena(loopFunc, birthFunc, deathFunc OrganismFunc,
	          adult, maxPopulation uint64, o *Organism) *Arena {
	a := &Arena{NewAverager(int(maxPopulation), -10000.0), adult,
		maxPopulation, loopFunc, deathFunc, birthFunc, sync.RWMutex{},
		sync.WaitGroup{}, false, 1, 0}
	a.waitGroup.Add(1)
	go a.organismLoop(o)
	return a
}

func (a *Arena) TotalDeaths() uint64 {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.totalDeaths
}

func (a *Arena) Population() uint64 {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.population
}

func (a *Arena) Stop() {
	a.mutex.Lock()
	a.stopped = true
	a.mutex.Unlock()
	a.waitGroup.Wait()
}

func (a *Arena) organismLoop(o *Organism) {
	defer a.waitGroup.Done()
	a.birthFunc(o)
	for {
		// Perform a single lifecycle
		a.loopFunc(o)
		a.mutex.Lock()
		keep := a.lifecycle(o)
		if !keep {
			a.deathFunc(o)
			a.population--
			a.mutex.Unlock()
			return
		}
		a.mutex.Unlock()
		runtime.Gosched()
	}
}

func (a *Arena) lifecycle(o *Organism) bool {
	if a.stopped {
		return false
	} else if o.health.Cycles < a.adultAge {
		return true
	}
	
	v := o.health.Value()
	avg := a.average.Push(v)
	if v < avg && a.population > 1 {
		// Kill the organism unless it's the last one left
		return false
	}
	
	// Overpopulation = no reproduction
	if a.population == a.maxPopulation {
		return true
	}
	
	// Keep the organism alive and reproduce
	dup := o.Reproduce()
	a.waitGroup.Add(1)
	go a.organismLoop(dup)
	
	return true
}
