# Attempt 3

This attempt will use the neural network used in [attempt2](../attempt2) along with a new evolutionary model. I hope that this will allow neural structures to build on top of eachother to solve larger problems piece by piece.

# Evolution

In this evolutionary model, each organism will have a net rate of pain. That is, each organism will have an associated "goodness" level. The less pain an organism receives per cycle--i.e. the more completely it solves the given problem--the higher it's "goodness level".

With this model, I hope that organisms will gradually build on top of one another and develop more and more complete solutions to given problems.

# Reproduction

This time, I think that reproduction will be completely random. Occasionally, links will be broken or neurons deleted. More frequently, neurons will be added.

This reproduction model means that the evolutionary process will not directly implement Hebbian learning. I believe that, in future models, the structure of the neural networks themselves might have to change so that Hebbian learning is still possible.

# Results

Sometimes, this algorithm outputs a fully functional XOR circuit in an instant. Other times, however, it runs indefinitely without ever finding a solution. In the latter case, it seems to sometimes stumble upon a half-solution, but then the organism in question quickly dies.

I am now wondering if my computers' limitations are my downfall. Sometimes, building a better organism will simply take tons of evolution--and in these cases, millions of children will be necessary in order for one of them to succeed. In my current simulation, I only allow for a few hundred organisms to live at one time. This makes it difficult for long strings of genes to gradually build up and evolve.

Additionally, I am questioning my evolution model. Adding neurons is much more likely than removing them, and this leads to a huge amount of unprecedented growth. This means that a lot of "bad" structures build up quickly that are hard to remove. The result is increasingly complex organisms which perform no better than their parents. I would try to tackle this by introducing an implementation of Occam's Razor, but then it would not be possible for intermediate structures to develop in the interim.

## Ideas for improvement

In my next model, I will probably track the fitness of every organism simultaneously and exactly. This way, I will be sure that the fittest organism is NEVER killed off.

I may modify this model so that deletion mutations occur on a neuron rather than on a link. This may fix the issue of unwanted growth.

## Update

I have performed two changes.

The first change is in the nature of mutations. Now, multiple mutations can happen (although the odds of n permutations are 1/3^(n-1)). A mutation can now either be an addition or a deletion of a neuron--no more deleting links.

The second change is in the nature of fitness. Now, the fitness of every organism is tracked live and compared completely fairly. This means that the fittest organism will *never* die and has a 100% chance of reproducing. While this sometimes means that the goal organism won't come about as quickly, it also ensures that the current cutting edge neural structure will be protected.

Now, it would seem that a good XOR circuit eventually turns up, although it does sometimes take a lot of time.

## Update 2

I am working on evolving a 2-digit adder. Now, the issue is delay: at first it evolves a nice network to perform addition of *most* cases. However, in order to do all cases, it takes two cycles, not one, and it is difficult for the entire network to evolve to delay itself by one extra cycle.

**UPDATE**: I think to do this, I will make it possible for new neurons to be inserted in the middle of existing links.
