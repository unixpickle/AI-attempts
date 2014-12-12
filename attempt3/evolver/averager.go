package evolver

type Averager struct {
	average    float64
	lastValues []float64
}

func NewAverager(count int, start float64) *Averager {
	res := &Averager{start, make([]float64, count)}
	for i := 0; i < count; i++ {
		res.lastValues[i] = start
	}
	return res
}

func (a *Averager) Push(value float64) float64 {
	min := 1000000.0
	max := -1000000.0
	for i := 0; i < len(a.lastValues) - 1; i++ {
		moving := a.lastValues[i + 1]
		a.lastValues[i] = moving
		if moving < min {
			min = moving
		}
		if moving > max {
			max = moving
		}
	}
	a.lastValues[len(a.lastValues)-1] = value
	if value >= max {
		return 1.0
	} else if value <= min {
		return 0.0
	}
	return (value - min) / (max - min)
}
