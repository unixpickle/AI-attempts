#ifndef __NNN1_NEURON_HPP__
#define __NNN1_NEURON_HPP__

namespace nnn1 {

class Link;

class Neuron {
public:
  Neuron(const Neuron &) = delete;
  Neuron & operator=(const Neuron &) = delete;
  
  Neuron();
  virtual ~Neuron();
  unsigned int CountMessages();
  
  /**
   * Cause this neuron to fire during the current cycle.
   */
  inline void Fire() {
    isFiring = true;
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
   * Based on internal state and [CountMessages], determine if this neuron
   * should fire on the next cycle.
   */
  virtual bool Cycle() = 0;

protected:
  friend class Link;
  
  Link * firstOutput = nullptr;
  Link * firstInput = nullptr;
  
  unsigned int inputCount = 0;
  unsigned int outputCount = 0;
  
  bool isFiring = false;
  bool willFire = false;
};

}

#endif
