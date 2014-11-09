#ifndef __NNN1_RANDOM_HPP__
#define __NNN1_RANDOM_HPP__

namespace nnn1 {

/**
 * Return a bool which has a certain [probability] of being `true`.
 */
bool RandomBool(double probability);

/**
 * Return a uniformly random `double` between 0 and 1.
 */
double RandomDouble();

}

#endif
