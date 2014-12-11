package evolver

import "sync"

type ArenaWorld func(o *Organism) bool

type Arena struct {
	average       *Averager
	highest       float64
	adultAge      uint64
	maxPopulation uint64
	world         ArenaWorld
	
	mutex        sync.RWMutex
	outputStream chan *Organism
	stopped      bool
	population   uint64
	totalDeaths  uint64
}

func NewArena(world ArenaWorld, adult, maxPopulation uint64,
	o *Organism) *Arena {
	a := &Arena{NewAverager(int(maxPopulation), -10000.0), -10000.0, adult,
		maxPopulation, world, sync.RWMutex{}, make(chan *Organism), false, 1, 0}
	// TODO: start goroutine for first organism
	return a
}

func (a *Arena) OutputStream() chan *Organism {
	return a.outputStream
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
