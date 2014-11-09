#include "random.hpp"
#include <cstdlib>

namespace nnn1 {

bool RandomBoolWithProbability(double probability) {
  if (probability >= 1.0) {
    return true;
  }
  return RandomDouble() >= probability;
}

double RandomDouble() {
  return (double)rand() / (double)RAND_MAX;
}

}
