
import itertools
import operator

def main():
    i = 0
    while True:
        factors = get_all_factors(nth_triangle_number(i))
        print i, nth_triangle_number(i), len(factors)
        if len(factors) > 500:
            break
        i += 1

def nth_triangle_number(n):
    return sum(range(n+1))

def get_all_factors (num):
    prime_factors = get_prime_factors(num)
    factors = set(prime_factors)
    factors.update([1, num])

    for i in range (2, len(prime_factors)):
        combinations = itertools.combinations(prime_factors, i)
        ifactors = [reduce(operator.mul, fs) for fs in combinations]
        factors.update(ifactors)

    return factors

def get_prime_factors(num):
    prime_factors = []
    prime = 2
    _num = num

    while prime < _num:
        if _num % prime == 0:
            prime_factors.append(prime)
            _num /= prime
            prime = 2
        else:
            prime = next_prime(prime)
    if _num != num: prime_factors.append(prime)

    return prime_factors

def is_prime(a):
    if (a <= 1): return False
    if (a == 2): return True
    if (a % 2 == 0): return False

    for i in range(3, a, 2):
        if (a % i == 0): return False
        if (i*i > a): break

    return True

def next_prime(n):
    #return the next prime larger than n
    while True:
        n += 1
        if is_prime(n): return n

main()
