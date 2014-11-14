#include "scoped-pass.hpp"
#include <nnn1>
#include <cassert>

using namespace nnn1;

void TestCreate();
void TestRemove();

int main() {
  TestCreate();
  TestRemove();
  return 0;
}

void TestCreate() {
  ScopedPass pass("Link::Create()");
  OrNeuron neuron1;
  OrNeuron neuron2;
  Link * link = Link::Create(neuron1, neuron2);
  assert(neuron1.GetOutputCount() == 1);
  assert(neuron1.GetInputCount() == 0);
  assert(neuron2.GetOutputCount() == 0);
  assert(neuron2.GetInputCount() == 1);
  assert(&link->GetSender() == &neuron1);
  assert(&link->GetReceiver() == &neuron2);
  assert(link->GetSenderNext() == nullptr);
  assert(link->GetReceiverNext() == nullptr);
  
  Link * link1 = Link::Create(neuron1, neuron2);
  assert(neuron1.GetOutputCount() == 2);
  assert(neuron1.GetInputCount() == 0);
  assert(neuron2.GetOutputCount() == 0);
  assert(neuron2.GetInputCount() == 2);
  assert(&link1->GetSender() == &neuron1);
  assert(&link1->GetReceiver() == &neuron2);
  assert(link1->GetSenderNext() == link);
  assert(link1->GetReceiverNext() == link);
  assert(link->GetSenderNext() == nullptr);
  assert(link->GetReceiverNext() == nullptr);
  
  link->Remove();
  link1->Remove();
}

void TestRemove() {
  ScopedPass pass("Link::Remove()");
  OrNeuron neuron1;
  OrNeuron neuron2;
  
  Link * link = Link::Create(neuron1, neuron2);
  Link * link1 = Link::Create(neuron1, neuron2);
  
  assert(link1->GetSenderNext() == link);
  assert(link1->GetReceiverNext() == link);
  link->Remove();
  assert(neuron1.GetOutputCount() == 1);
  assert(neuron1.GetInputCount() == 0);
  assert(neuron2.GetOutputCount() == 0);
  assert(neuron2.GetInputCount() == 1);
  assert(!link1->GetSenderNext());
  assert(!link1->GetReceiverNext());
  link1->Remove();
  assert(neuron1.GetOutputCount() == 0);
  assert(neuron1.GetInputCount() == 0);
  assert(neuron2.GetOutputCount() == 0);
  assert(neuron2.GetInputCount() == 0);
  
  link = Link::Create(neuron1, neuron2);
  link1 = Link::Create(neuron2, neuron1);
  assert(link->GetReceiverNext() == nullptr);
  assert(link->GetSenderNext() == nullptr);
  assert(link1->GetReceiverNext() == nullptr);
  assert(link1->GetSenderNext() == nullptr);
  assert(neuron1.GetOutputCount() == 1);
  assert(neuron1.GetInputCount() == 1);
  assert(neuron2.GetOutputCount() == 1);
  assert(neuron2.GetInputCount() == 1);
  link->Remove();
  link1->Remove();
  assert(neuron1.GetOutputCount() == 0);
  assert(neuron1.GetInputCount() == 0);
  assert(neuron2.GetOutputCount() == 0);
  assert(neuron2.GetInputCount() == 0);
}
