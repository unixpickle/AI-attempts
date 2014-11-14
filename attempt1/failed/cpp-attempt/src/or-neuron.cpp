#include "or-neuron.hpp"

namespace nnn1 {

bool OrNeuron::Cycle() {
  return CountMessages() != 0;
}

}
