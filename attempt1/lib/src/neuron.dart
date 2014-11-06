part of nnn;

abstract class Neuron {
  List<Neuron> destinations = [];
  List<Neuron> inputs = [];
  bool _willFire = false;
  bool _firing = false;
  
  bool get firing => _firing;
  
  int get messageCount {
    int count = 0;
    for (var input in inputs) {
      if (input._firing) {
        ++count;
      }
    }
    return count;
  }
  
  bool nextCycle();
}

class InputNeuron extends Neuron {
  bool _fireNext = false;
  
  void trigger() {
    _fireNext = true;
  }
  
  bool nextCycle() {
    bool result = _fireNext;
    _fireNext = false;
    return result;
  }
}

class OutputNeuron extends Neuron {
  bool _triggered = false;
  
  bool get triggered => _triggered;
  
  bool nextCycle() {
    _triggered = (messageCount != 0);
    return false;
  }
}

class DelayNeuron extends Neuron {
  bool nextCycle() {
    return messageCount != 0;
  }
}

class AndNeuron extends Neuron {
  bool nextCycle() {
    return messageCount == 2;
  }
}

class OrNeuron extends Neuron {
  bool nextCycle() {
    return messageCount > 0;
  }
}

class XorNeuron extends Neuron {
  bool nextCycle() {
    return messageCount == 1;
  }
}
