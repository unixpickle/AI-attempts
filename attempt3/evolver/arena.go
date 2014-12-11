package evolver

import "sync"

type ArenaCycle func(o *Organism) bool

type Arena struct {
	average       *Averager
	highest       float64
	adultAge      uint64
	maxPopulation uint64
	cycleFunc     ArenaCycle
	
	mutex        sync.RWMutex
	outputStream chan *Organism
	stopped      bool
	population   uint64
	totalDeaths  uint64
}