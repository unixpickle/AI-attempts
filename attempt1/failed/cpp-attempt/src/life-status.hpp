#ifndef __NNN1_LIFE_STATUS_HPP__
#define __NNN1_LIFE_STATUS_HPP__

#include <cmath>

namespace nnn1 {

class LifeStatus {
public:
  inline LifeStatus() {}
  
  inline bool GetPermanent() const {
    return permanent;
  }
  
  inline void SetPermanent(bool flag) {
    permanent = flag;
  }
  
  inline unsigned long long GetLastUsed() const {
    return lastUsed;
  }
  
  inline void SetLastUsed(unsigned long long value) {
    lastUsed = value;
  }
  
  inline double GetNetPain() const {
    return netPain;
  }
  
  inline void SetNetPain(double value) {
    netPain = value;
  }
  
  inline double GetPainScale() const {
    return painScale;
  }
  
  inline void SetPainScale(double value) {
    painScale = value;
  }
  
  /**
   * Returns the probability that this neuron will survive the next painful
   * experience.
   */
  double GetKeepProbability() const;
  
  /**
   * Apply pain of a given [intensity] to this life status.
   * 
   * The [currentTime] is used to compute the difference between [GetLastUsed]
   * and the present time.
   */
  void UpdateNetPain(unsigned long long currentTime, double intensity);
  
protected:
  bool permanent = true;
  unsigned long long lastUsed = 0;
  double netPain = 0.0;
  double painScale = 1.0;
};

}

#endif
