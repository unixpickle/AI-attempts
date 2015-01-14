package brain

import (
	"testing"
)

func TestXORCircuit(t *testing.T) {
	network := Network{
		Node{NodeTypeXOR, false, []int{}},
		Node{NodeTypeXOR, false, []int{}},
		Node{NodeTypeXOR, false, []int{0, 1}},
	}
	brain := NewBrain(network)
	
	// Case: 0^0
	brain.Activity()[0] = false
	brain.Activity()[1] = false
	brain.Activity()[2] = false
	brain.Cycle()
	for _, x := range brain.Activity() {
		if x {
			t.Error("Invalid output for 0^0.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 0^0.")
	}
	
	// Case: 1^1
	brain.Activity()[0] = true
	brain.Activity()[1] = true
	brain.Activity()[2] = false
	brain.Cycle()
	for _, x := range brain.Activity() {
		if x {
			t.Error("Invalid output for 1^1.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 1^1.")
	}
	
	// Case: 1^0
	brain.Activity()[0] = true
	brain.Activity()[1] = false
	brain.Activity()[2] = false
	brain.Cycle()
	for i, x := range brain.Activity() {
		if (i == 2 && !x) || (i != 2 && x) {
			t.Error("Invalid output for 1^0.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 1^0.")
	}
	
	// Case: 0^1
	brain.Activity()[1] = false
	brain.Activity()[0] = true
	brain.Activity()[2] = false
	brain.Cycle()
	for i, x := range brain.Activity() {
		if (i == 2 && !x) || (i != 2 && x) {
			t.Error("Invalid output for 0^1.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 0^1.")
	}
}

func TestANDCircuit(t *testing.T) {
	network := Network{
		Node{NodeTypeXOR, false, []int{}},
		Node{NodeTypeXOR, false, []int{}},
		Node{NodeTypeAND, false, []int{0, 1}},
	}
	brain := NewBrain(network)
	
	// Case: 0&0
	brain.Activity()[0] = false
	brain.Activity()[1] = false
	brain.Activity()[2] = false
	brain.Cycle()
	for _, x := range brain.Activity() {
		if x {
			t.Error("Invalid output for 0&0.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 0&0.")
	}
	
	// Case: 1&1
	brain.Activity()[0] = true
	brain.Activity()[1] = true
	brain.Activity()[2] = false
	brain.Cycle()
	for i, x := range brain.Activity() {
		if (i != 2 && x) || (i == 2 && !x) {
			t.Error("Invalid output for 1&1.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 1&1.")
	}
	
	// Case: 1&0
	brain.Activity()[0] = true
	brain.Activity()[1] = false
	brain.Activity()[2] = false
	brain.Cycle()
	for _, x := range brain.Activity() {
		if x {
			t.Error("Invalid output for 1&0.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 1&0.")
	}
	
	// Case: 0&1
	brain.Activity()[1] = false
	brain.Activity()[0] = true
	brain.Activity()[2] = false
	brain.Cycle()
	for _, x := range brain.Activity() {
		if x {
			t.Error("Invalid output for 0&1.")
		}
	}
	brain.Cycle()
	if brain.Activity()[0] || brain.Activity()[1] || brain.Activity()[2] {
		t.Error("Invalid lasting signals for 0&1.")
	}
}
