package evolver

import (
	"math/rand"
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
}

func NewArena(loopFunc, birthFunc, deathFunc OrganismFunc,
	          adult, maxPopulation uint64, o *Organism) *Arena {
	averageSize := int(maxPopulation)
	if averageSize < 100 {
		averageSize = 100
	}
	a := &Arena{NewAverager(int(maxPopulation) * 10, -10000.0), adult,
		maxPopulation, loopFunc, deathFunc, birthFunc, sync.RWMutex{},
		sync.WaitGroup{}, false, 1}
	a.waitGroup.Add(1)
	go a.organismLoop(o)
	return a
}

func (a *Arena) Population() uint64 {
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
	percentile := a.average.Push(v)
	
	live, reproduce := organismDestination(percentile)
	
	if !live {
		return a.population == 1
	}
	
	if reproduce && a.population < a.maxPopulation {
		dup := o.Reproduce()
		a.waitGroup.Add(1)
		a.population++
		go a.organismLoop(dup)
	}
	
	return true
}

func organismDestination(val float64) (bool, bool) {
	if val < 0.5 {
		// No reproduction, but according to probability we might want to keep
		// the organism alive.
		return rand.Float64()/2.0 >= val, false
	} else {
		return true, (0.5 + rand.Float64()/2.0) < val
	}
}
