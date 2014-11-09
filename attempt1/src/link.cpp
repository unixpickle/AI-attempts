#include "link.hpp"

namespace nnn1 {

Link::~Link() {
}

Link * Link::Create(Neuron & sender, Neuron & receiver) {
  return new Link(sender, receiver);
}

void Link::Remove() {
  // Remove this link from the sender
  if (!senderLast) {
    sender.firstOutput = senderNext;
  } else {
    senderLast->senderNext = senderNext;
  }
  if (senderNext) {
    senderNext->senderLast = senderLast;
  }
  --sender.outputCount;
  
  // Remove this link from the receiver
  if (!receiverLast) {
    receiver.firstInput = receiverNext;
  } else {
    receiverLast->receiverNext = receiverNext;
  }
  if (receiverNext) {
    receiverNext->receiverLast = nullptr;
  }
  --receiver.inputCount;
  
  delete this;
}

Link::Link(Neuron & sender, Neuron & receiver)
    : sender(sender),
      receiver(receiver) {
  senderNext = sender.firstOutput;
  if (senderNext) {
    senderNext->senderLast = this;
  }
  sender.firstOutput = this;
  ++sender.outputCount;
  receiverNext = receiver.firstInput;
  if (receiverNext) {
    receiverNext->receiverLast = this;
  }
  receiver.firstInput = this;
  ++receiver.inputCount;
}

}
