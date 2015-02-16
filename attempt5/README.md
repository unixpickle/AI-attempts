# Self-modifying code

I am going to attempt to make a programming language which is ideal for self-modifying, interconnected, and parallelized code.

After I create a runtime, I will create a few artificial "brains". I will start by making simple programs which modify themselves in predictable ways to learn simple tasks. Once I demonstrate that this works, I will try to create a more advanced program which will have human-like capabilities such as short-term memory, image processing, etc.

# Runtime overview

The runtime/bytecode will have some features to make it ideal for AI. While I have not decided and/or figured out any exact details, here are some aspects that I think it will reflect:

 * Separate "modules" or "subroutines"&mdash;some static, some modifiable
 * Inputs and outputs look just like any other subroutine
 * Subroutines can call each other and run asynchronously (with some kind of limit)
 * A subroutine can modify itself or any other subroutine which it references
 * A subroutine is Turing complete
 * Subroutines which are not referenced are destroyed
 * Subroutines can generate new subroutines (with some kind of limit)

# Runtime features

Here is my attempt to outline the runtime as I will implement it.

NOTE: The runtime won't really be "bytecode". However, it will consists of a series of "instructions", so it's kind of like a bytecode.

## Subroutines and addresses.

**Subroutines** are logical units of "memory" which contain zero or more instructions. Each instruction has a **memory address**. Memory addresses coordinates&mdash;ordered pairs of the form (subroutine, index). A subroutine is not *numerically addressed*&mdash;that is, it does not have a numerical identifier. Rather, subroutines can only be referenced by copying existing references. However, instructions within a subroutine are indexed, and these indexes can be incremented and decremented as numbers.

## Instructions

Each instruction takes one argument. Every argument is a memory address. Note that a subroutine can only run once at a time, so any calls to a running subroutine will be ignored.

This section will be kind of like library documentation. I will denote each instruction as "name(argument)". I will try to give the arguments meaningful names.

 * nop(unused) - This does nothing. This can be used to store an address for use somewhere else.
 * jmp(address) - Move the execution to a different address.
 * call(address) - Concurrently run a different address.
 * callback(output) - Get the address of the current instruction, add 2 to its instruction index, and set the result as the argument of the instruction at "output".
 * inc(instruction) - This increments the instruction index of the argument of "instruction".
 * dinc(instruction) - This is equivalent to running inc(instruction) twice.
 * dec(instruction) - This decrements the instruction index of the argument of "instruction"
 * eq(compare) - This looks the argument of the next instruction. If that argument is equal to "compare", this changes the instruction after the next instruction to be a "jmp". Otherwise, this changes said instruction to be a "nop". This does not affect any arguments, only the instruction itself.
 * setinst(address) - This sets the instruction at "address" to the instruction after the next instruction. It does not change the argument at "address", only the instruction.
 * getinst(address) - This sets the instruction after the next one to the instruction at "address". It does not change the argument, only the instruction.
 * setarg(address) - This sets the argument at "address" to the argument of next instruction.
 * getarg(address) - This sets the argument of the next instruction to be the argument at "address"
 * beginning(address) - This sets the instruction index of the argument at "address" to 0.
 * create(address) - This creates a new subroutine. It sets the argument at "address" to equal the coordinate (s, 0), where s is the new subroutine.
 * halt(unused) - This stops executing the current subroutine. It is basically a "return" statement, but there is nowhere no return to.
 * copy(address) - This performs an indirection in the form of a "set" operation. It puts the argument at "address" into the argument of the address which is the argument of the next instruction.

**Update:** I have drafted a NAND gate in this bytecode. In addition, I have created an addition circuit. Thus, I think it's pretty safe to say that the language is Turing complete.