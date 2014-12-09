package nnn

type Lifetime struct {
	Permanent bool
	LastUsed  CycleTime
	Creation  CycleTime
}

func NewLifetime() Lifetime {
	return Lifetime{true, 0, 0}
}
