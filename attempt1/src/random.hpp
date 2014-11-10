#ifndef __NNN1_RANDOM_HPP__
#define __NNN1_RANDOM_HPP__

namespace nnn1 {

/**
 * Return a bool which has a certain [probability] of being `true`.
 */
bool RandomBool(double probability = 0.5);

/**
 * Return a uniformly random `double` between 0 and 1.
 */
double RandomDouble();

/**
 * Generate a random number between 0 (inclusive) and [max] (exclusive).
 */
unsigned int RandomNumber(unsigned int max);

}

#endif
