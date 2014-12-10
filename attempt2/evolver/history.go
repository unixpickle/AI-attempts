package evolver

import (
	"math"
	"math/rand"
)

type History struct {
	LastFired uint64
	Pain      float64
	Permanent bool
}

func NewHistory() *History {
	return &History{0, 0.0, false}
}

func (h *History) ApplyPain(pain float64, time uint64) {
	// Effect of pain is inversely proportional to its delay
	h.Pain += pain / float64(time - h.LastFired)
}

func (h *History) Clone() *History {
	cpy := *h
	return &cpy
}

func (h *History) RandomKeep() bool {
	if h.Permanent {
		return true
	}
	return rand.Float64() > h.Value()
}

func (h *History) Value() float64 {
	// This is a nice asymptotic function which approaches 1 for negative values
	// and 0 for positive values.
	return 1.0 - (math.Atan(h.Pain) + math.Pi/2)/math.Pi
}
