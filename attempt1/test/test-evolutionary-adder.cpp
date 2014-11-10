#include "scoped-pass.hpp"
#include <nnn1>

using namespace nnn1;

int main() {
  ScopedPass pass("Evolver [single-bit adder]");
  
  Network network;
  Evolver evolver(network);
  Neuron * input0 = new OrNeuron();
  Neuron * input1 = new OrNeuron();
  Neuron * output = new OrNeuron();
  network.AddNeuron(*input0);
  network.AddNeuron(*input1);
  network.AddNeuron(*output);
  
  while (true) {
    bool flag0 = RandomBool();
    bool flag1 = RandomBool();
    if (flag0) input0->Fire();
    else input0->Inhibit();
    if (flag1) input1->Fire();
    else input1->Inhibit();
    bool gotResponse = false;
    bool expectingResponse = (flag0 && !flag1) || (flag1 && !flag0);
    // Give it ten cycles to respond with the answer
    for (int i = 0; i < 10; ++i) {
      network.Cycle();
      if (output->IsFiring()) {
        gotResponse = true;
      }
      evolver.Prune();
      evolver.Grow();
    }
    if (gotResponse != expectingResponse) {
      std::cout << "wrong!" << std::endl;
      network.UpdateLivesWithPain(0.5);
    } else {
      std::cout << "right!" << std::endl;
    }
    // Give the network some time to flush out its activity
    for (int i = 0; i < 5; ++i) {
      network.Cycle();
      evolver.Prune();
      evolver.Grow();
    }
  }
  
  return 0;
}
