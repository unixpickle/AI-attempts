package brain

import (
	"math/rand"
)

// Mutate performs one random insertion or deletion on a network.
// The input network may be modified as a result of the mutation, similarly to
// the append() builtin.
func Mutate(n Network) Network {
	if len(n) == 0 {
		return AddRandomNode(n)
	}

	// 0 = add node, 1 = add link, 2 = remove node, 3 = remove link
	var operation int
	num := rand.Float64()
	if CanDeleteLink(n) && CanDeleteNode(n) {
		operation = int(num / 4.0)
	} else if CanDeleteLink(n) {
		operation = int(num / 3.0)
		if operation == 3 {
			operation = 4
		}
	} else if CanDeleteNode(n) {
		operation = int(num / 3.0)
	} else {
		operation = int(num / 2.0)
	}

	if operation == 0 {
		return AddRandomNode(n)
	} else if operation == 1 {
		return AddRandomLink(n)
	} else if operation == 2 {
		return RemoveRandomNode(n)
	} else {
		return RemoveRandomLink(n)
	}
}

// AddRandomLink generates a single link in the network.
func AddRandomLink(n Network) Network {
	length := len(n)

	// Cannot add a link to an empty network
	if length == 0 {
		return n
	}

	r1 := rand.Intn(length)
	r2 := rand.Intn(length)
	n[r1].Inputs = append(n[r1].Inputs, r2)

	return n
}

// AddRandomNode generates a node and adds it to the network.
// The new node will be connected semi-randomly to other nodes in the network.
func AddRandomNode(n Network) Network {
	oldLen := len(n)

	// Add zero, one, or two random inputs.
	var inputs []int
	if oldLen == 0 {
		inputs = []int{}
	} else if oldLen == 1 || rand.Intn(2) == 0 {
		inputs = []int{rand.Intn(oldLen)}
	} else {
		// Generate two non-overlapping random numbers.
		inOne := rand.Intn(oldLen)
		inTwo := rand.Intn(oldLen - 1)
		if inTwo >= inOne {
			inTwo++
		}
		inputs = []int{inOne, inTwo}
	}

	if oldLen != 0 {
		// Add a random output
		output := rand.Intn(oldLen)
		n[output].Inputs = append(n[output].Inputs, oldLen)
	}

	node := Node{rand.Intn(2), false, inputs}
	return append(n, node)
}

// CanDeleteLink determines if a network has any inter-node connections.
func CanDeleteLink(n Network) bool {
	for _, x := range n {
		if len(x.Inputs) > 0 {
			return true
		}
	}
	return false
}

// CanDeleteNode determines if a network has any non-permanent nodes.
func CanDeleteNode(n Network) bool {
	for _, x := range n {
		if !x.Permanent {
			return true
		}
	}
	return false
}
