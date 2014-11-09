#include "evolver.hpp"
#include "link.hpp"
#include "random.hpp"

namespace nnn1 {

Evolver::Evolver(Network & network) : network(network) {}

void Evolver::Prune() {
  Neuron * neuron = network.GetFirstNeuron();
  while (neuron) {
    if (!RandomBool(neuron->GetLifeStatus().GetKeepProbability())) {
      // Throw out the entire neuron
      Neuron * current = neuron;
      neuron = neuron->GetNextNeuron();
      network.RemoveNeuron(*current);
      delete current;
      continue;
    }
    // Remove links randomly
    Link * link = neuron->GetFirstOutput();
    while (link) {
      if (!RandomBool(link->GetLifeStatus().GetKeepProbability())) {
        Link * current = link;
        link = link->GetSenderNext();
        current->Remove();
      } else {
        link = link->GetSenderNext();
      }
    }
    neuron = neuron->GetNextNeuron();
  }
}

void Evolver::Grow() {
  // TODO: create a list of every neuron based on it's recentness; then, pick
  // one using a scaled random number.
}

double Evolver::Recentness(Neuron * neuron) {
  unsigned long long age = network.GetCycleCount() -
      neuron->GetLifeStatus().GetLastUsed();
  return 1.0 / (double)(age + 1);
}

}
