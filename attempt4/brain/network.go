package brain

import (
	"fmt"
)

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

// String makes a beautiful human-readable string from the network.
func (n Network) String() string {
	res := "{Network:\n"
	for i, obj := range n {
		funcStr := "AND"
		if obj.Type == NodeTypeXOR {
			funcStr = "XOR"
		}
		res += fmt.Sprintln("  ", i, "-", funcStr)
		for _, input := range obj.Inputs {
			res += fmt.Sprintln("    <-", input)
		}
	}
	return res + "}"
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
