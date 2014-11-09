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
  
  /**
   * Get the number of neural cycles that this network has performed.
   */
  inline unsigned long long GetCycleCount() {
    return cycleCount;
  }
  
  /**
   * Get the first neuron in the network.
   */
  inline Neuron * GetFirstNeuron() {
    return firstNeuron;
  }
  
  /**
   * Modify all neurons and links to reflect pain of a given [intensity].
   *
   * A negative [intensity] is interpreted as pleasure and can be used for
   * positive reinforcement.
   */
  void UpdateLivesWithPain(double intensity);
  
private:
  Neuron * firstNeuron = nullptr;
  
  unsigned long long cycleCount = 0;
};

}

#endif
