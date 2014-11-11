use std::cell::RefCell;

pub struct Neuron<'a>;

pub struct Link<'a> {
  pub input: &'a RefCell<Neuron<'a>>,
  pub output: &'a RefCell<Neuron<'a>>
}
  
impl<'a> Link<'a> {
  pub fn new(input: &'a RefCell<Neuron<'a>>, output: &'a RefCell<Neuron<'a>>) 
      -> Link<'a> {
    Link {input: input, output: output}
  }
}
