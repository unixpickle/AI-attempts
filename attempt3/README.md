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
