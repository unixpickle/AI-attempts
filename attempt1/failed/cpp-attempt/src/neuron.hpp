#ifndef __NNN1_NEURON_HPP__
#define __NNN1_NEURON_HPP__

#include "life-status.hpp"

namespace nnn1 {

class Link;
class Network;

class Neuron {
public:
  Neuron(const Neuron &) = delete;
  Neuron & operator=(const Neuron &) = delete;
  
  Neuron();
  
  /**
   * Removes all links in and out of this neuron.
   */
  virtual ~Neuron();
  
  /**
   * Returns the number of incoming messages during the current cycle.
   */
  unsigned int CountMessages();
  
  /**
   * Cause this neuron to fire during the current cycle.
   */
  inline void Fire() {
    isFiring = true;
  }
  
  /**
   * Prevent this neuron from firing during the current cycle.
   */
  inline void Inhibit() {
    isFiring = false;
  }
  
  /**
   * Returns `true` if and only if the neuron is firing during the current
   * cycle.
   */
  inline bool IsFiring() {
    return isFiring;
  }
  
  /**
   * Returns the number of inputs for this neuron.
   */
  inline unsigned int GetInputCount() {
    return inputCount;
  }
  
  /**
   * Returns the number of outputs for this neuron.
   */
  inline unsigned int GetOutputCount() {
    return outputCount;
  }
  
  /**
   * Returns the status of this neuron's life.
   */
  inline LifeStatus & GetLifeStatus() {
    return lifeStatus;
  }
  
  /**
   * Based on internal state and [CountMessages], determine if this neuron
   * should fire on the next cycle.
   */
  virtual bool Cycle() = 0;
  
  /**
   * Return the first output link on this neuron.
   */
  inline Link * GetFirstOutput() {
    return firstOutput;
  }
  
  /**
   * Return the first input link on this neuron.
   */
  inline Link * GetFirstInput() {
    return firstInput;
  }
  
  /**
   * Return the next neuron in the network.
   */
  inline Neuron * GetNextNeuron() {
    return nextNeuron;
  }

protected:
  friend class Link;
  friend class Network;
  
  Link * firstOutput = nullptr;
  Link * firstInput = nullptr;
  
  unsigned int inputCount = 0;
  unsigned int outputCount = 0;
  
  bool isFiring = false;
  bool willFire = false;
  
  LifeStatus lifeStatus;
  
  Neuron * nextNeuron = nullptr;
  Neuron * lastNeuron = nullptr;
};

}

#endif
