#include "scoped-pass.hpp"
#include <nnn1>
#include <cassert>

using namespace nnn1;

Network * CreateAdder(Neuron *& a0, Neuron *& b0, Neuron *& b1, Neuron *& c0,
                      Neuron *& c1);

void TestAdd(bool a, bool b0, bool b1, bool c0, bool c1);

int main() {
  TestAdd(false, false, false, false, false);
  TestAdd(false, true, false, true, false);
  TestAdd(false, false, true, false, true);
  TestAdd(false, true, true, true, true);
  TestAdd(true, false, false, true, false);
  TestAdd(true, true, false, false, true);
  TestAdd(true, false, true, true, true);
  TestAdd(true, true, true, false, false);
  return 0;
}

Network * CreateAdder(Neuron *& a0, Neuron *& b0, Neuron *& b1, Neuron *& c0,
                      Neuron *& c1) {
  Network * network = new Network();
  
  // Inputs
  a0 = new OrNeuron();
  b0 = new OrNeuron();
  b1 = new OrNeuron();
  // Outputs
  c0 = new OrNeuron();
  c1 = new OrNeuron();
  // Circuit for first bit
  auto xor0 = new XorNeuron();
  auto delay0 = new OrNeuron();
  // Circuit for second bit
  auto delay1 = new OrNeuron();
  auto xor1 = new XorNeuron();
  auto and0 = new AndNeuron();
  
  // Connect neurons
  Link::Create(*a0, *xor0);
  Link::Create(*b0, *xor0);
  Link::Create(*a0, *and0);
  Link::Create(*b0, *and0);
  Link::Create(*b1, *delay1);
  Link::Create(*xor0, *delay0);
  Link::Create(*delay1, *xor1);
  Link::Create(*and0, *xor1);
  Link::Create(*delay0, *c0);
  Link::Create(*xor1, *c1);
  
  // Add neurons to network
  network->AddNeuron(*a0);
  network->AddNeuron(*b0);
  network->AddNeuron(*b1);
  network->AddNeuron(*c0);
  network->AddNeuron(*c1);
  network->AddNeuron(*xor0);
  network->AddNeuron(*delay0);
  network->AddNeuron(*delay1);
  network->AddNeuron(*xor1);
  network->AddNeuron(*and0);
  
  return network;
}

void TestAdd(bool aFlag, bool b0Flag, bool b1Flag, bool c0Flag, bool c1Flag) {
  ScopedPass("Adder [(", aFlag, "), (", b0Flag, ", ", b1Flag, ")]");
  Neuron * a0, * b0, * b1, * c0, * c1;
  Network * network = CreateAdder(a0, b0, b1, c0, c1);
  
  if (aFlag) a0->Fire();
  if (b0Flag) b0->Fire();
  if (b1Flag) b1->Fire();
  
  for (int i = 0; i < 3; ++i) {
    network->Cycle();
  }
  
  assert(c0->IsFiring() == c0Flag);
  assert(c1->IsFiring() == c1Flag);
  
  network->Cycle();
  assert(network->CountFiring() == 0);
  
  delete network;
}
