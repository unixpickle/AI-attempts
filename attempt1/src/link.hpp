#ifndef __NNN1_LINK_HPP__
#define __NNN1_LINK_HPP__

#include "neuron.hpp"

namespace nnn1 {

/**
 * A unidirectional link between two neurons.
 */
class Link {
public:
  Link(const Link &) = delete;
  Link & operator=(const Link &) = delete;
  
  virtual ~Link();
  
  /**
   * Allocate a new link between a [sender] and a [receiver], add it to both,
   * and return it.
   */
  static Link * Create(Neuron & sender, Neuron & receiver);
  
  /**
   * Remove this link from the sender and receiver and deallocate it.
   */
  void Remove();
  
  /**
   * Get the sender.
   */
  inline Neuron & GetSender() {
    return sender;
  }
  
  /**
   * Get the receiver.
   */
  inline Neuron & GetReceiver() {
    return receiver;
  }
  
  /**
   * Get the next [Link] in the sender's output linked-list.
   */
  inline Link * GetSenderNext() {
    return senderNext;
  }
  
  /**
   * Get the next [Link] in the receiver's input linked-list.
   */
  inline Link * GetReceiverNext() {
    return receiverNext;
  }
  
private:
  Link(Neuron & sender, Neuron & receiver);
  
  Neuron & sender;
  Neuron & receiver;
  Link * senderNext, * senderLast = nullptr;
  Link * receiverNext, * receiverLast = nullptr;
};

}

#endif
