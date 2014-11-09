#include "life-status.hpp"

namespace nnn1 {

double LifeStatus::GetKeepProbability() const {
  if (permanent) {
    return 1.0;
  }
  double value = exp(-painScale * netPain);
  if (value > 1) {
    return 1.0;
  } else {
    return value;
  }
}

void LifeStatus::UpdateNetPain(unsigned long long currentTime,
                               double intensity) {
  double painValue = intensity / (double)(currentTime - GetLastUsed() + 1);
  netPain += painValue;
}

}
