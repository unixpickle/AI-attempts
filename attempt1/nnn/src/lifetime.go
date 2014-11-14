package nnn

type Lifetime struct {
	permanent bool
	lastUsed  CycleTime
}

func NewLifetime() Lifetime {
	return Lifetime{true, 0}
}
