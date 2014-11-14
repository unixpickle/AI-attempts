package nnn

type Lifetime struct {
	Permanent bool
	LastUsed  CycleTime
}

func NewLifetime() Lifetime {
	return Lifetime{true, 0}
}
