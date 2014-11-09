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
   * Returns a numerical representation of how recently a neuron was fired. The
   * probability of [neuron] being chosen for a connection becomes `R/N`, where
   * R is the recentness of [neuron] and N is the sum total of all the neurons'
   * recentness values.
   *
   * By default, this simply returns the reciprocal of the age of the neuron,
   * measured in cycles.
   */
  virtual double Recentness(Neuron * neuron);

private:
  Network & network;
};

}

#endif
