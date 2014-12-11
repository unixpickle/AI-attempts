# Attempt 3

This attempt will use the neural network used in [attempt2](../attempt2) along with a new evolutionary model. I hope that this will allow neural structures to build on top of eachother to solve larger problems piece by piece.

# Evolution

In this evolutionary model, each organism will have a net rate of pain. That is, each organism will have an associated "goodness" level. The less pain an organism receives per cycle--i.e. the more completely it solves the given problem--the higher it's "goodness level".

With this model, I hope that organisms will gradually build on top of one another and develop more and more complete solutions to given problems.

# Reproduction

This time, I think that reproduction will be completely random. Occasionally, links will be broken or neurons deleted. More frequently, neurons will be added.

This reproduction model means that the evolutionary process will not directly implement Hebbian learning. I believe that, in future models, the structure of the neural networks themselves might have to change so that Hebbian learning is still possible.
