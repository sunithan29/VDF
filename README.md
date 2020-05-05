# Verifiable Delay Function (VDF) 

Verifiable delay function is slow to compute and easy to verify. This is computed in a prescribed amount of time, t and 
cannot be performed faster even if it is parallelizable.

This implementation is based on sloth which is best characterized as a time-asymmetric encoding, offering a trade-off
in practice between computation and inversion (verification), and thus can be viewed as a pseudo-VDF. This a weak
VDF because it requires the honest Eval to use greater than polylog(t) parallelism to run 
in parallel time t (the delay parameter).

## Sloth

Wesolowski proposed creating a more difficult puzzle for a small p by chaining a series of such puzzles together
(interleaved with a simple permutation) in a construction called Sloth. However, this does not meet asymptotic definition
of a VDF because it does not offer (asymptotically) efficient verification. This is only used as a building block 
to help construct a practical VDF. Sloth is comparable to a hash chain of length t with t checkpoints provided as a proof.

Implementation notes:

To build from the source code you need to have a working Go environment with version 1.14 or greater installed.

       $ cd vdf
       
       $ make build

To run: 

      $ ./vdf --x=7 --t=1000
## References
1. [Sloth(Quadratic Residue) Verifiable Delay Functions](https://eprint.iacr.org/2018/601.pdf). Boneh, 2018
2. [Efficient Verifiable Delay Functions](https://eprint.iacr.org/2018/623.pdf). Wesolowski, 2018

### (Currently working)
