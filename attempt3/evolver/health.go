package evolver

type Health struct {
	Cycles uint64
	Pain   float64
}

func NewHealth() *Health {
	return &Health{0, 0.0}
}

func (h *Health) ApplyPain(pain float64) {
	h.Pain += pain
}

func (h *Health) Cycle() {
	h.Cycles++
}

func (h *Health) Clone() *Health {
	x := *h
	return &x
}

func (h *Health) Value() float64 {
	return -h.Pain / float64(h.Cycles)
}
