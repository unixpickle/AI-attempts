package nnn;

import "testing"

func TestAdditionCircuit(t *testing.T) {
	adder, controls := createSimpleAdder()
	// TODO: test them here
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
