# Model

### Neurons, links, and messages

In this model of neural networks, there are two fundamental structures: **neurons** and **links**. A neuron takes zero or more inputs and can provide zero or more outputs. Links go from the output of one neuron to the input of the next.

Data is transferred in a neural network using **messages**. A message itself carries no data: it is simply an impulse. During a **neural cycle**, every active message is sent over exactly one link. Thus, for a message to traverse a path of 20 links, it would take 20 neural cycles.

### Types of neurons

There are three types of neurons:

 * **And neurons**: If all inputs receive a message on the same neural cycle, the neuron sends a message to all its outputs during the next cycle.
 * **Or neurons**: If any input receives a message on a neural cycle, the neuron sends a message to all its outputs during the next cycle.
 * **Xor neurons**: If an odd number of inputs receive a message on a neural cycle, the neuron sends a message to all its outputs during the next cycle.

### Evolution

I am still trying to figure out how these neural networks will improve themselves over time to learn from experience. We'll see what I come up with.
