part of nnn;

class NeuralNetwork {
  List<Neuron> neurons = [];
  
  void cycle() {
    for (var neuron in neurons) {
      neuron._willFire = neuron.nextCycle();
    }
    for (var neuron in neurons) {
      neuron._firing = neuron._willFire;
    }
  }
  
  int firingCount() {
    int count = 0;
    for (var neuron in neurons) {
      if (neuron.firing) {
        ++count;
      }
    }
    return count;
  }
}
