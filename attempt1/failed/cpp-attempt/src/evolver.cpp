#include "evolver.hpp"
#include "link.hpp"
#include "random.hpp"
#include "xor-neuron.hpp"
#include "and-neuron.hpp"
#include "or-neuron.hpp"
#include <cassert>

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
  // This code is pretty lame.
  
  if (network.GetNeuronCount() < 2) {
    return;
  }
  
  // So yeah, idk why I use numbers for this; I just wanna see if this works.
  unsigned int inputCount = 1;
  unsigned int outputCount;
  if (network.GetNeuronCount() == 2) {
    outputCount = 1;
  } else {
    outputCount = RandomBool() ? 1 : 2;
  }
  
  // Generate a random neuron and hook it up to random things.
  
  // TODO: here, customize how easily links vs. neurons get deleted.
  Neuron * neuron = GenerateNeuron();
  network.AddNeuron(*neuron);
  Neuron * neurons[3];
  RandomNeurons(neurons, inputCount + outputCount);
  Link::Create(*neurons[0], *neuron)->GetLifeStatus().SetPermanent(false);
  Link::Create(*neuron, *neurons[1])->GetLifeStatus().SetPermanent(false);
  if (outputCount == 2) {
    Link::Create(*neuron, *neurons[2])->GetLifeStatus().SetPermanent(false);
  }
}

Neuron * Evolver::GenerateNeuron() {
  // Basically, this is just an unweighted random neuron factory.
  unsigned int num = RandomNumber(3);
  Neuron * result = nullptr;
  if (num == 0) {
    result = new XorNeuron();
  } else if (num == 1) {
    result = new AndNeuron();
  } else {
    result = new OrNeuron();
  }
  result->GetLifeStatus().SetPermanent(false);
  return result;
}

void Evolver::RandomNeurons(Neuron ** output, unsigned int count) {
  // This could easily be the worst code I've ever written; I doubt it works.
  assert(count <= network.GetNeuronCount());
  
  // This structure will seem important once you realize the silly way I'm
  // trying to do this...
  struct ListPlace {
    Neuron * neuron;
    double weight;
    double priorWeight;
  };
  
  // Create a list of every neuron, it's weight, and the amount of weight
  // before it.
  ListPlace * places = new ListPlace[network.GetNeuronCount()];
  double totalWeight = 0.0;
  unsigned int i = 0;
  Neuron * neuron = network.GetFirstNeuron();
  while (neuron) {
    assert(i < network.GetNeuronCount());
    places[i].neuron = neuron;
    places[i].weight = Weight(neuron);
    places[i].priorWeight = totalWeight;
    totalWeight += places[i].weight;
    ++i;
    neuron = neuron->GetNextNeuron();
  }
  
  // "Randomly" choose [count] neurons from the list.
  unsigned int listSize = i;
  for (i = 0; i < count; ++i) {
    assert(listSize > 0);
    double thresh = RandomDouble() * totalWeight;
    unsigned int pickIndex;
    for (pickIndex = 0; pickIndex < listSize; ++pickIndex) {
      if (places[pickIndex].weight + places[pickIndex].priorWeight >=
          thresh) {
        break;
      }
    }
    // Seriously, this shouldn't happen, but who knows with floating points...
    if (pickIndex == listSize) {
      pickIndex = listSize - 1;
    }
    output[i] = places[pickIndex].neuron;
    
    // This is my convoluted way of removing an item from the list
    double subWeight = places[pickIndex].weight;
    totalWeight -= subWeight;
    --listSize;
    for (unsigned int j = pickIndex; j < listSize; ++j) {
      places[j] = places[j + 1];
      places[j].priorWeight -= subWeight;
    }
  }
  delete places;
}

double Evolver::Weight(Neuron * neuron) {
  unsigned long long age = network.GetCycleCount() -
      neuron->GetLifeStatus().GetLastUsed();
  return 1.0 / (double)(age + 1);
}

}
