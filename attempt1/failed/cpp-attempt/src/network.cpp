#include "network.hpp"
#include "link.hpp"

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
  ++neuronCount;
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
  --neuronCount;
}

void Network::Cycle() {
  ++cycleCount;
  
  // Record which neurons will fire on the next cycle
  Neuron * neuron = firstNeuron;
  while (neuron) {
    neuron->willFire = neuron->Cycle();
    neuron = neuron->nextNeuron;
  }
  // Fire away!
  neuron = firstNeuron;
  while (neuron) {
    if (neuron->willFire) {
      // The neuron is being fired during this cycle.
      neuron->GetLifeStatus().SetLastUsed(cycleCount);
    }
    if (neuron->isFiring) {
      // The neuron was fired during the last cycle, so all its outputs should
      // be updated accordingly.
      Link * link = neuron->firstOutput;
      while (link) {
        link->GetLifeStatus().SetLastUsed(cycleCount - 1);
        link = link->GetSenderNext();
      }
    }
    neuron->isFiring = neuron->willFire;
    neuron->willFire = false;
    neuron = neuron->nextNeuron;
  }
}

unsigned int Network::CountFiring() {
  unsigned int count = 0;
  Neuron * neuron = firstNeuron;
  while (neuron) {
    if (neuron->IsFiring()) {
      ++count;
    }
    neuron = neuron->nextNeuron;
  }
  return count;
}

void Network::UpdateLivesWithPain(double intensity) {
  Neuron * neuron = firstNeuron;
  while (neuron) {
    neuron->GetLifeStatus().UpdateNetPain(cycleCount, intensity);
    Link * link = neuron->firstOutput;
    while (link) {
      link->GetLifeStatus().UpdateNetPain(cycleCount, intensity);
      link = link->GetSenderNext();
    }
    neuron = neuron->nextNeuron;
  }
}

}
