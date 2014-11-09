#include "xor-neuron.hpp"

namespace nnn1 {

bool XorNeuron::Cycle() {
  return (CountMessages() & 1) != 0;
}

}
