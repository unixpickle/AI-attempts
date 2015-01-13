# Evolution with implicit inhibition

This model will be similar to my [last attempt](attempt3), but it will lack OR neurons. I call this "implicit inhibition" because AND and XOR neurons are inherently inhibitable: if an AND neuron has an unused input, it is inhibited; if an XOR neuron gets an extra signal, it's result is flipped. This is in sharp contrast to OR neurons which fire easily and can only be inhibited through evolutionary deletion.

I will use selection to breed organisms which do increasingly well at various tasks. At this time, I have no special plans for improving my evolution technique.

## Motivation

A problem with my last two attempts was that circuits did not evolve to be more and more complex. Rather, a circuit would evolve to solve a subset of the problem too fast. These circuits were unlikely to produce children which solved the *entire* problem; such an occurance would usually involve a slower intermediate step wherein the organism did not improve but rather slowed down.

One approach to fixing the problem described above may be sexual reproduction. This is a path I am considering, but I think I'll have to learn some graph theory before I can do it.

In the meantime, I think removing OR neurons will simplify networks and make inhibition easier. Inhibition will allow a partial solution to inhibit a different partial solution depending on the input.