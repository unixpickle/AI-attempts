#ifndef __NNN1_OR_NEURON_HPP__
#define __NNN1_OR_NEURON_HPP__

#include "neuron.hpp"

namespace nnn1 {

class OrNeuron : public Neuron {
public:
  virtual bool Cycle();
};

}

#endif
