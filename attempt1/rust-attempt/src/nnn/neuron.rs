use std::cell::RefCell;

pub struct Link<'a>;

pub enum NeuronKind {
  XorNeuron,
  AndNeuron,
  OrNeuron
}

pub struct Neuron<'a> {
  pub inputs: Vec<&'a RefCell<Link<'a>>>,
  pub outputs: Vec<&'a RefCell<Link<'a>>>,
  pub firing: bool,
  pub will_fire: bool,
  pub kind: NeuronKind
}

impl<'a> Neuron<'a> {

 pub fn new(t: NeuronKind) -> Neuron<'a> {
   Neuron {inputs: vec![], outputs: vec![], firing: false,
           will_fire: false, kind: t}
 }

 pub fn next_cycle(&self, in_count: uint) -> bool {
   match self.kind {
     XorNeuron => in_count % 2 != 0,
     AndNeuron => in_count == self.inputs.len(),
     OrNeuron => in_count != 0
   }
 }

}
