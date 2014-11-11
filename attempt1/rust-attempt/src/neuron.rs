mod neuron {

  pub enum NeuronType {
    XorNeuron,
    AndNeuron,
    OrNeuron
  }

  pub struct Neuron {
    inputs: Vec<RefCell<Link>>,
    outputs: Vec<RefCell<Link>>,
    firing: bool,
    will_fire: bool,
    type: NeuronType
  }

  impl Neuron {

    pub fn new(t: NeuronType) -> Neuron {
      Neuron {inputs: vec![], outputs: vec![], firing: false,
              will_fire: false, type: t}
    }
  
    pub fn next_cycle(&self, in_count: uint) -> bool {
      match self.type {
        XorNeuron => in_count % 2 != 0,
        AndNeuron => in_count == self.inputs.len(),
        OrNeuron => in_count != 0
      }
    }

  }

}
