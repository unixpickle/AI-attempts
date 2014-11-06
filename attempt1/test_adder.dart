import 'lib/nnn.dart';

void main() {
  // Test all single-bit plus double-bit number addition cases.
  testAddition([false, false, false], [false, false]);
  testAddition([false, true, false], [true, false]);
  testAddition([false, false, true], [false, true]);
  testAddition([false, true, true], [true, true]);
  testAddition([true, false, false], [true, false]);
  testAddition([true, true, false], [false, true]);
  testAddition([true, false, true], [true, true]);
  testAddition([true, true, true], [false, false]);
  print('passed!');
}

void testAddition(List<bool> inputs, List<bool> outputs) {
  // Inputs are [a0, b0, b1]; outputs are [c0, c1].
  // Finds a + b and truncates the result into c.
  
  // Create the entire network
  NeuralNetwork network = new NeuralNetwork();
  // Inputs
  var a0 = new InputNeuron();
  var b0 = new InputNeuron();
  var b1 = new InputNeuron();
  // Outputs
  var c0 = new OutputNeuron();
  var c1 = new OutputNeuron();
  // Circuit for first bit
  var xor0 = new XorNeuron();
  var delay0 = new DelayNeuron();
  // Circuit for second bit
  var delay1 = new DelayNeuron();
  var xor1 = new XorNeuron();
  var and0 = new AndNeuron();
  
  // Connect neurons
  a0.connect(xor0);
  b0.connect(xor0);
  a0.connect(and0);
  b0.connect(and0);
  b1.connect(delay1);
  xor0.connect(delay0);
  delay1.connect(xor1);
  and0.connect(xor1);
  delay0.connect(c0);
  xor1.connect(c1);
  
  // Add neurons to network
  network.neurons.addAll([a0, b0, b1, c0, c1, xor0, delay0, delay1, xor1, 
      and0]);
  
  // Set initial states
  if (inputs[0]) a0.trigger();
  if (inputs[1]) b0.trigger();
  if (inputs[2]) b1.trigger();
  
  // Perform 4 neural cycles
  for (var i = 0; i < 4; ++i) {
    network.cycle();
  }
  if (network.firingCount() != 0) {
    throw 'unexpected firing count: ${network.firingCount()}';
  } else if (c0.triggered != outputs[0]) {
    throw 'unexpected first output: ${c0.triggered}';
  } else if (c1.triggered != outputs[1]) {
    throw 'unexpected second output: ${c1.triggered}';
  }
}
