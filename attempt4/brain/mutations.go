package brain

import (
	"math/rand"
)

// Mutate performs one random insertion or deletion on a network.
// The input network will not be modified as a result of the mutation.
func Mutate(rawInput Network) Network {
	n := rawInput.Clone()

	if len(n) == 0 {
		return addRandomNode(n)
	}

	// 0 = add node, 1 = add link, 2 = remove node, 3 = remove link
	var operation int
	num := rand.Float64()
	if canDeleteLink(n) && canDeleteNode(n) {
		operation = int(num / 4.0)
	} else if canDeleteLink(n) {
		operation = int(num / 3.0)
		if operation == 3 {
			operation = 4
		}
	} else if canDeleteNode(n) {
		operation = int(num / 3.0)
	} else {
		operation = int(num / 2.0)
	}

	if operation == 0 {
		return addRandomNode(n)
	} else if operation == 1 {
		return addRandomLink(n)
	} else if operation == 2 {
		return removeRandomNode(n)
	} else {
		return removeRandomLink(n)
	}
}

func addRandomLink(n Network) Network {
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

func addRandomNode(n Network) Network {
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
		// Add a random output.
		output := rand.Intn(oldLen)
		n[output].Inputs = append(n[output].Inputs, oldLen)
	}

	node := Node{rand.Intn(2), false, inputs}
	return append(n, node)
}

func canDeleteLink(n Network) bool {
	for _, x := range n {
		if len(x.Inputs) > 0 {
			return true
		}
	}
	return false
}

func canDeleteNode(n Network) bool {
	for _, x := range n {
		if !x.Permanent {
			return true
		}
	}
	return false
}

func removeNode(n Network, idx int) Network {
	// Remove the node.
	res := n
	copy(res[idx:], res[idx+1:])
	res = res[0 : len(res)-1]

	// Subtract or delete appropriate input indexes.
	// TODO: make this more efficient.
	for i := 0; i < len(res); i++ {
		newInputs := make([]int, 0, len(res[i].Inputs))
		for _, input := range res[i].Inputs {
			if input == idx {
				continue
			} else if input > idx {
				newInputs = append(newInputs, input-1)
			} else {
				newInputs = append(newInputs, input)
			}
		}
		res[i].Inputs = newInputs
	}

	return res
}

func removeRandomLink(n Network) Network {
	// Gather a list of links so we can pick one randomly.
	links := make([]linkInfo, 0)
	for i, x := range n {
		for j, _ := range x.Inputs {
			link := linkInfo{i, j}
			links = append(links, link)
		}
	}

	// No links means no mutation.
	if len(links) == 0 {
		return n
	}

	// Remove the link entry
	rem := links[rand.Intn(len(links))]
	node := &n[rem.owner]
	node.Inputs[rem.idx] = node.Inputs[len(node.Inputs)-1]
	node.Inputs = node.Inputs[0 : len(node.Inputs)-1]

	return n
}

func removeRandomNode(n Network) Network {
	// Create a list of indexes for non-permanent nodes.
	nonPerm := make([]int, 0, len(n))
	for i, x := range n {
		if !x.Permanent {
			nonPerm = append(nonPerm, i)
		}
	}

	// Pick a node to delete
	idx := nonPerm[rand.Intn(len(nonPerm))]
	return removeNode(n, idx)
}

type linkInfo struct {
	owner int
	idx   int
}
