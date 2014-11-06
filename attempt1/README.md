# Model

### Neurons, links, and messages

In this model of neural networks, there are two fundamental structures: **neurons** and **links**. A neuron takes zero or more inputs and can provide zero or more outputs. Links go between the output of one neuron to the input of the next.

Data is transferred in a neural network using **messages**. A message itself carries no data: it is simply an impulse. Every **neural cycle**, every message is sent over exactly one link. Thus, for a message to traverse a path through 20 links, it would take 20 neural cycles.

### Types of neurons

There are six types of neurons:

 * **Input neurons** have zero inputs and any number of outputs. These neurons represent sensory data.
 * **Output neurons** have any number of inputs and zero outputs. These neurons represent motor control.
 * **Delay neurons** have one input and any number of outputs. Delay neurons simply forward messages from their input to their outputs. They may be used to "delay" a message by one neural cycle.
 * **And neurons** have two inputs and any number of outputs. If both inputs receive a message on the same neural cycle, the neuron sends a message to all its outputs during the next cycle.
 * **Or neurons** have two inputs and any number of outputs. If either input receives a message on the same neural cycle, the neuron sends a message to all its outputs during the next cycle.
 * **Xor neurons** have two inputs and any number of outputs. If exactly input receives a message on a neural cycle, the neuron sends a message to all its outputs during the next cycle.

### Evolution

I am still trying to figure out how these neural networks will improve themselves over time to learn from experience. We'll see what I come up with.
