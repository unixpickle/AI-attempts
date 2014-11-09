# Model

### Neurons, links, and messages

In this model of neural networks, there are two fundamental structures: **neurons** and **links**. A neuron takes zero or more inputs and can provide zero or more outputs. Links go from the output of one neuron to the input of the next.

Data is transferred in a neural network using **messages**. A message itself carries no data: it is simply an impulse. During a **neural cycle**, every active message is sent over exactly one link. Thus, for a message to traverse a path of 20 links, it would take 20 neural cycles.

### Types of neurons

There are three types of neurons:

 * **And neurons**: If all inputs receive a message on the same neural cycle, the neuron sends a message to all its outputs during the next cycle.
 * **Or neurons**: If any input receives a message on a neural cycle, the neuron sends a message to all its outputs during the next cycle.
 * **Xor neurons**: If an odd number of inputs receive a message on a neural cycle, the neuron sends a message to all its outputs during the next cycle.

## Evolution

While I am not sure what model will prove to be effective, I have an idea in mind.

### Destruction of neurons and links

When the neural network experiences pain, neurons and links which fired more recently will be more affected by the pain.  When a neuron or link has experienced enough net pain, it's deletion will become probable.

### Creation and linking of neurons

Periodically, new random neurons will be generated and added to the network. These neurons will have several semi-random links.  The more recently a neuron has fired, the more likely it is to be one of these links' inputs or outputs.

### Hebbian learning

According to [Hebbian theory](http://en.wikipedia.org/wiki/Hebbian_theory), neurons which fire at nearly the same time tend to become associated. This is important to the functionality of associative memory.

The process of creating and linking neurons will favor links with neurons which recently fired. This will allow thoughts which occur in parallel to be linked together.
