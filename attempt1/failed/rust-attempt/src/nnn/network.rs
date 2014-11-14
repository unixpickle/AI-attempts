use libc::{size_t, c_void, malloc, free};
use std::mem;
use std::ptr;

pub enum NeuronKind {
  XorNeuron,
  AndNeuron,
  OrNeuron,
}

pub struct Neuron<T> {
  inputs: Vec<*mut Link<T>>,
  outputs: Vec<*mut Link<T>>,
  pub firing: bool,
  pub will_fire: bool,
  pub kind: NeuronKind,
  pub data: T,
}

struct Link<T> {
  input: *mut Neuron<T>,
  output: *mut Neuron<T>,
}

pub struct Network<T> {
  neurons: Vec<*mut Neuron<T>>
}

fn alloc_link<T>(input: *mut Neuron<T>, output: *mut Neuron<T>)
    -> *mut Link<T> {
  unsafe {
    let res = malloc(mem::size_of::<Link<T>>() as size_t) as *mut Link<T>;
    assert!(!res.is_null());
    ptr::write(&mut *res, Link {input: input, output: output});
    res
  }
}

fn alloc_neuron<T>(kind: NeuronKind, data: T) -> *mut Neuron<T> {
  unsafe {
    let res = malloc(mem::size_of::<Neuron<T>>() as size_t) as *mut Neuron<T>;
    assert!(!res.is_null());
    let neuron = Neuron {inputs: vec![], outputs: vec![],
                         firing: false, will_fire: false, kind: kind,
                         data: data};
    ptr::write(&mut *res, neuron);
    res
  }
}

fn free_link<T>(link: *mut Link<T>) {
  unsafe {
    free(link as *mut c_void);
  }
}

fn free_neuron<T>(neuron: *mut Neuron<T>) {
  unsafe {
    // Trigger drop() on the neuron.
    ptr::read(neuron);
  
    // Deallocate the neuron.
    free(neuron as *mut c_void);
  }
}

#[unsafe_destructor]
impl<T> Drop for Neuron<T> {
  #[unsafe_destructor]
  pub fn drop(&mut self) {
    // Free inputs
    loop {
      let link = match self.inputs.pop() {
        None => break,
        Some(x) => x
      };
      unsafe {
        let otherNeuron = (*link).input;
        let outputs = &mut otherNeuron.outputs;
        (*outputs).retain(|x| x != link);
      }
      free_link(link);
    }
  
    // Free outputs
    loop {
      let link = match self.outputs.pop() {
        None => break,
        Some(x) => x
      };
      unsafe {
        let otherNeuron = (*link).output;
        let inputs = &mut otherNeuron.inputs;
        (*inputs).retain(|x| x != link);
      }
      free_link(link);
    }
  }
}

impl<T> Network<T> {
  pub fn new() -> Network<T> {
    return Network {neurons: vec![]}
  }
  
  pub fn input_count(&self) -> uint {
    let mut count: uint = 0;
    for link in self.inputs {
      unsafe {
        let source = &mut (*link).output;
        if source.firing {
          count += 1;
        }
      }
    }
    count
  }
  
  pub fn next_cycle(&self) -> bool {
    match self.kind {
      XorNeuron => self.input_count() % 2 != 0,
      AndNeuron => self.input_count() == self.inputs.len(),
      OrNeuron => self.input_count() != 0
    }
  }
}

#[unsafe_destructor]
impl<T> Drop for Network<T> {
  #[unsafe_destructor]
  pub fn drop(&mut self) {
    loop {
      let neuron = match self.neurons.pop() {
        None => break,
        Some(x) => x
      };
      free_neuron(neuron);
    }
  }
}
