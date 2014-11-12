use std::cell::RefCell;
use neuron::Neuron;

pub struct Link<'a> {
  pub input: &'a RefCell<Neuron<'a>>,
  pub output: &'a RefCell<Neuron<'a>>
}

impl<'a> Link<'a> {
  pub fn new(output: &'a RefCell<Neuron<'a>>, input: &'a RefCell<Neuron<'a>>)
      -> Link<'a> {
    Link {input: input, output: output}
  }
}
