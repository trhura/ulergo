#! /usr/bin/env python3

from prob12 import get_all_factors
from functools import lru_cache

def main():
    amiable_numbers = set([])
    for i in range (2, 10000):
        d = spd(i)
        if spd(d) == i and d != i:
            amiable_numbers.update([i, d])
    print(amiable_numbers)
    print(sum(amiable_numbers))

@lru_cache(maxsize=10000)
def spd(num):
    return sum(get_proper_divisors(num))

def get_proper_divisors (num):
    factors = get_all_factors(num)
    factors.remove(num)
    return factors

if __name__ == "__main__":
    main()
