package nnn

type LifetimeInfo struct {
	permanent bool
	lastUsed  CycleTime
}

type LinkLifetime struct {
	info      LifetimeInfo
	permanent bool
	lastUsed  CycleTime
	netPain   float64
}

type Lifetime interface {
	GetLifetimeInfo() *LifetimeInfo
}

func (self *LifetimeInfo) GetLifetimeInfo() *LifetimeInfo {
	return self
}

func (self *LinkLifetime) GetLifetimeInfo() *LifetimeInfo {
	return &self.info
}

func SetLastUsed(lt Lifetime, time CycleTime) {
	lt.GetLifetimeInfo().lastUsed = time
}

func IsPermanent(lt Lifetime) bool {
	return lt.GetLifetimeInfo().permanent
}
