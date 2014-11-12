use std::cell::RefCell;
use link::Link;

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

#[test]
fn xor_neuron_test() {
  let neuron = Neuron::new(XorNeuron);
  assert!(neuron.next_cycle(0) == false);
  assert!(neuron.next_cycle(1) == true);
  assert!(neuron.next_cycle(2) == false);
  assert!(neuron.next_cycle(3) == true);
  assert!(neuron.next_cycle(4) == false);
}

#[test]
fn and_neuron_test() {
  let ref1 = RefCell::new(Neuron::new(AndNeuron));
  let ref2 = RefCell::new(Neuron::new(AndNeuron));
  let ref3 = RefCell::new(Link::new(&ref1, &ref2));
  ref1.borrow_mut().outputs.push(&ref3);
  ref2.borrow_mut().inputs.push(&ref3);
  assert!(ref1.borrow().next_cycle(0) == true);
  assert!(ref2.borrow().next_cycle(0) == false);
  assert!(ref2.borrow().next_cycle(1) == true);
}
