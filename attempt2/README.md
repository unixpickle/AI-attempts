# Model

In this attempt, I will use the same three types of neurons that I used in [attempt1](https://github.com/unixpickle/AI-attempts/tree/master/attempt1). However, instead of evolving a single organism, I will introduce the process of reproduction. This way, it is easy to revert changes (i.e. if organism A gives birth to a worse organism B, organism A can survive while organism B dies).

# Reproduction

Organisms will reproduce at a rate proportional to the amount of pleasure they have experienced. Their children will be altered semi-randomly, just like in attempt1.

# Problems

This model has a few problems.

The first problem rests in the way organisms are judged against one another. In this model, I use the total amount of pain that an organism has experienced to determine how fit it is. Instead, I should be using the average rate of pain that it receives. This way, even though organisms may experience more and more pain, better organisms outlive worse ones.

The second problem occurs when organisms grow out of control. In general, I need to figure out a way to evolve organisms so that they shed useless neural pathways. Perhaps I will only allow organisms to grow to a certain size, and that size will gradually increase as the simulation runs for longer and longer.

A lot of problems stem from the fact that I cannot branch limitlessly. Biological organisms on earth can continue to reproduce rather freely without being bounded (except, of course, for the "overpopulation" problem some politicians like to point to). If I could launch unlimited goroutines, competition wouldn't even be necessary and eventually I would arrive at the perfect organism. This is not completely "brute force" because I would still bias reproduction towards those organisms which performed their tasks better.
