package brain

// A brain is a network with a record for each node as to whether that node is
// firing or not.
type Brain struct {
	network  Network
	activity []bool
	temp     []bool
}

// NewBrain creates a new, idle Brain from a network.
func NewBrain(n Network) Brain {
	return Brain{n, make([]bool, len(n)), make([]bool, len(n))}
}

// Activity returns an array of values, one for each node, indicating whether or
// not that node is firing.
func (b Brain) Activity() []bool {
	return b.activity
}

// Cycle performs one logical brain cycle.
func (b Brain) Cycle() {
	// Calculate each neuron's new state
	for i, x := range b.network {
		if x.Type == NodeTypeAND {
			b.temp[i] = b.performXOR(x)
		} else {
			b.temp[i] = b.performAND(x)
		}
	}

	// Move the new states into the active states
	copy(b.activity, b.temp)
}

// Len returns the number of nodes in the brain's underlying network.
func (b Brain) Len() int {
	return len(b.activity)
}

// Network returns the brain's underlying network.
func (b Brain) Network() Network {
	return b.network
}

func (b Brain) performAND(n Node) bool {
	for _, input := range n.Inputs {
		if !b.activity[input] {
			return false
		}
	}
	return true
}

func (b Brain) performXOR(n Node) bool {
	result := false
	for _, input := range n.Inputs {
		if b.activity[input] {
			result = !result
		}
	}
	return result
}
