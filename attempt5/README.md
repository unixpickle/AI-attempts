# Self-modifying code

I am going to attempt to make a simple bytecode which is ideal for interconnected, parallelized execution while also giving code the ability to modify itself.

My plan is to create some simple artificial "brains" using my self-modifying bytecode. At first, I will create pretty simple self-modifying programs which modify themselves in predictable ways to learn simple tasks. Once I demonstrate that this works, I will try to create a more advanced initial program which will have human-like capabilities, like short-term memory, image processing, etc.

# Bytecode overview

While I have not decided/figured out any exact details, here are some aspects that I think my bytecode will reflect:

 * Separate "modules" or "subroutines"--some innate, some modifiable
 * Inputs and outputs look just like any other subroutine
 * Subroutines can call each other and run asynchronously (with some kind of limit)
 * A subroutine can modify itself or any other subroutine which it references
 * A subroutine is Turing complete
 * Subroutines which are not referenced are destroyed
 * Subroutines can create new subroutines (with some kind of limit)
