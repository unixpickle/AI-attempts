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
	var sum float64 = 0.0
	for i := 0; i < len(a.lastValues) - 1; i++ {
		a.lastValues[i] = a.lastValues[i + 1]
		sum += a.lastValues[i]
	}
	a.lastValues[len(a.lastValues)-1] = value
	sum += value
	a.average = sum / float64(len(a.lastValues))
	return a.average
}
