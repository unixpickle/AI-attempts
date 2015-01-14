package brain

const (
	NodeTypeAND = 0
	NodeTypeXOR = 1
)

// A network is an ordered list of nodes.
type Network []Node

// Clone creates a deep copy of a network.
func (n Network) Clone() Network {
	res := make(Network, len(n))
	for i, x := range n {
		inputsCopy := make([]int, len(x.Inputs))
		copy(inputsCopy, x.Inputs)
		res[i] = Node{x.Type, x.Permanent, inputsCopy}
	}
	return res
}

// A Node is a typed node in a graph.
// Nodes store information about their inputs.
// Nodes may be marked as "permanent" if they should not be deleted across
// evolution cycles.
type Node struct {
	Type      int
	Permanent bool
	Inputs    []int
}
