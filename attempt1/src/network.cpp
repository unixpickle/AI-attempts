#include "network.hpp"

namespace nnn1 {
  
Network::~Network() {
  while (firstNeuron) {
    Neuron * next = firstNeuron->nextNeuron;
    delete firstNeuron;
    firstNeuron = next;
  }
}

void Network::AddNeuron(Neuron & neuron) {
  neuron.nextNeuron = firstNeuron;
  firstNeuron = &neuron;
}

void Network::RemoveNeuron(Neuron & neuron) {
  if (neuron.lastNeuron) {
    neuron.lastNeuron->nextNeuron = neuron.nextNeuron;
  } else {
    firstNeuron = neuron.nextNeuron;
  }
  if (neuron.nextNeuron) {
    neuron.nextNeuron->lastNeuron = neuron.lastNeuron;
  }
}

void Network::Cycle() {
  // Record which neurons will fire on the next cycle
  Neuron * neuron = firstNeuron;
  while (neuron) {
    neuron->willFire = neuron->Cycle();
    neuron = neuron->nextNeuron;
  }
  // Fire away!
  neuron = firstNeuron;
  while (neuron) {
    neuron->isFiring = neuron->willFire;
    neuron->willFire = false;
    neuron = neuron->nextNeuron;
  }
}

unsigned int Network::CountFiring() {
  int count = 0;
  Neuron * neuron = firstNeuron;
  while (neuron) {
    if (neuron->IsFiring()) {
      ++count;
    }
    neuron = neuron->nextNeuron;
  }
  return count;
}

}
