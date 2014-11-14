package nnn

import "testing"

func TestAdditionCircuit(t *testing.T) {
	runAdderFunction(t, mapFromInput(0, 0, 0, 0, 0))
	runAdderFunction(t, mapFromInput(0, 1, 0, 1, 0))
	runAdderFunction(t, mapFromInput(0, 0, 1, 0, 1))
	runAdderFunction(t, mapFromInput(0, 1, 1, 1, 1))
	runAdderFunction(t, mapFromInput(1, 0, 0, 1, 0))
	runAdderFunction(t, mapFromInput(1, 1, 0, 0, 1))
	runAdderFunction(t, mapFromInput(1, 0, 1, 1, 1))
	runAdderFunction(t, mapFromInput(1, 1, 1, 0, 0))
}

func mapFromInput(a, b0, b1, c0, c1 int) map[string]bool {
	return map[string]bool{"a0": a != 0, "b0": b0 != 0, "b1": b1 != 0,
		"c0": c0 != 0, "c1": c1 != 0}
}

func createSimpleAdder() (*Network, map[string]*Neuron) {
	result := NewNetwork()

	// Create the inputs.
	a0 := NewOrNeuron()
	b0 := NewOrNeuron()
	b1 := NewOrNeuron()
	// Create the outputs.
	c0 := NewOrNeuron()
	c1 := NewOrNeuron()
	// Create circuit for first bit.
	xor0 := NewXorNeuron()
	delay0 := NewOrNeuron()
	// Create circuit for second bit.
	delay1 := NewOrNeuron()
	xor1 := NewXorNeuron()
	and0 := NewAndNeuron()

	// Connect neurons
	NewLink(a0, xor0)
	NewLink(b0, xor0)
	NewLink(a0, and0)
	NewLink(b0, and0)
	NewLink(b1, delay1)
	NewLink(xor0, delay0)
	NewLink(delay1, xor1)
	NewLink(and0, xor1)
	NewLink(delay0, c0)
	NewLink(xor1, c1)

	// Add neurons to network
	result.AddNeuron(a0)
	result.AddNeuron(b0)
	result.AddNeuron(b1)
	result.AddNeuron(c0)
	result.AddNeuron(c1)
	result.AddNeuron(xor0)
	result.AddNeuron(delay0)
	result.AddNeuron(delay1)
	result.AddNeuron(xor1)
	result.AddNeuron(and0)

	controlMap := map[string]*Neuron{"a0": a0, "b0": b0, "b1": b1, "c0": c0,
		"c1": c1}

	return result, controlMap
}

func runAdderFunction(t *testing.T, values map[string]bool) {
	network, controls := createSimpleAdder()
	controls["a0"].firing = values["a0"]
	controls["b0"].firing = values["b0"]
	controls["b1"].firing = values["b1"]
	for i := 0; i < 3; i++ {
		network.Cycle()
	}
	if controls["c0"].firing != values["c0"] {
		t.Error("Unexpected c0 value:", values)
	} else if controls["c1"].firing != values["c1"] {
		t.Error("Unexpected c1 value:", values)
	}
	network.Cycle()
	if network.CountFiring() > 0 {
		t.Error("Neurons still firing after addition:", values)
	}
}
