#include "and-neuron.hpp"

namespace nnn1 {

bool AndNeuron::Cycle() {
  return CountMessages() == GetInputCount();
}

}
