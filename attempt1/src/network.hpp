#ifndef __NNN1_NETWORK_HPP__
#define __NNN1_NETWORK_HPP__

#include "neuron.hpp"

namespace nnn1 {

/**
 * An array of neurons which are synchronized on a per-cycle basis.
 */
class Network {
public:
  /**
   * Deallocate a network and delete all its Neurons and their links.
   */
  ~Network();
  
  /**
   * Add a neuron.
   */
  void AddNeuron(Neuron &);
  
  /**
   * Remove a neuron.
   */
  void RemoveNeuron(Neuron &);
  
  /**
   * Perform a neural cycle.
   */
  void Cycle();
  
  /**
   * Returns the number of neurons that are firing.
   */
  unsigned int CountFiring();
  
private:
  Neuron * firstNeuron = nullptr;
};

}

#endif
