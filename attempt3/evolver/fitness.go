package evolver

type Fitness struct {
	healths map[int]float64
	freeIds []int
}

func NewFitness(count int) *Fitness {
	res := &Fitness{make(map[int]float64), make([]int, count)}
	for i := 0; i < count; i++ {
		res.freeIds[i] = i
	}
	return res
}

func (f *Fitness) NumFree() int {
	return len(f.freeIds)
}

func (f *Fitness) Alloc() int {
	freeNum := len(f.freeIds)
	if freeNum == 0 {
		panic("no free ids when calling GetId().")
	}
	res := f.freeIds[freeNum-1]
	f.freeIds = f.freeIds[0:freeNum-1]
	return res
}

func (f *Fitness) Free(id int) {
	delete(f.healths, id)
	f.freeIds = append(f.freeIds, id)
}

func (f *Fitness) Percentile(id int, fitness float64) float64 {
	f.healths[id] = fitness
	greater := 0
	total := 0
	for k, v := range f.healths {
		if k == id {
			continue
		}
		total++
		if v > fitness {
			greater++
		}
	}
	if total == 0 {
		return 1.0
	}
	return float64(total - greater) / float64(total)
}
