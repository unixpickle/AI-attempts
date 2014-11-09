#ifndef __NNN1_AND_NEURON_HPP__
#define __NNN1_AND_NEURON_HPP__

#include "neuron.hpp"

namespace nnn1 {

class AndNeuron : public Neuron {
public:
  virtual bool Cycle();
};

}

#endif
