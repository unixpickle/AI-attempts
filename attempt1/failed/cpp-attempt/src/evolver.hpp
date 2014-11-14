#ifndef __NNN1_EVOLVER_HPP__
#define __NNN1_EVOLVER_HPP__

#include "network.hpp"

namespace nnn1 {

class Evolver {
public:
  Evolver(Network &);
  
  /**
   * Randomly delete neurons and links based on their pain factor.
   */
  void Prune();
  
  /**
   * Insert a neuron of a random type with semi-random connections.
   */
  void Grow();

protected:
  /**
   * Picks random neurons based on their weights and returns them in an 
   * [output] list.
   */
  virtual void RandomNeurons(Neuron ** output, unsigned int count);
  
  /**
   * Generate a random neuron and return it.
   */
  virtual Neuron * GenerateNeuron();
  
  /**
   * Returns a numerical representation of the "bias" that this neuron should
   * receive when selecting neurons to connect.
   
   * The probability of [neuron] being chosen for a connection becomes `R/N`,
   * where R is the weight of [neuron] and N is the sum total of all the
   * neurons' weight.
   *
   * By default, this simply returns the reciprocal of the age of the neuron,
   * measured in cycles.
   */
  virtual double Weight(Neuron * neuron);

private:
  Network & network;
};

}

#endif
