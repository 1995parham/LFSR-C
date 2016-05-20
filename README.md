# LFSR-Go
## Introduction
Go Implementation of Linear Feedback Shift Register with some useful tools
based on [Michael Foukarakis LFSR implementation and tools](https://github.com/mfukar/lfsr).  
## LFSR ?
In computing, a linear-feedback shift register (LFSR) is a shift register whose input bit
is a linear function of its previous state.
Applications of LFSRs include generating pseudo-random numbers, pseudo-noise sequences,
fast digital counters, and whitening sequences. Both hardware and software implementations of LFSRs are common.
### Fibonacci LFSR
Fibonacci LFSR implementation for `x^16 + x^14 + x^13 + x^11 + 1` polynomial.
```c
# include <stdint.h>
int main(void)
{
    uint16_t start_state = 0xACE1u;  /* Any nonzero start state will work. */
    uint16_t lfsr = start_state;
    uint16_t bit;                    /* Must be 16bit to allow bit<<15 later in the code */
    unsigned period = 0;

    do
    {
        /* taps: 16 14 13 11; feedback polynomial: x^16 + x^14 + x^13 + x^11 + 1 */
        bit  = ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 5) ) & 1;
        lfsr =  (lfsr >> 1) | (bit << 15);
        ++period;
    } while (lfsr != start_state);

    return 0;
}
```
