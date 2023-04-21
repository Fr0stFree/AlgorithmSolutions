"""
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
Given an integer n, return true if n is an ugly number.
"""
class Solution:
    """Not ready yet"""
    def isUgly(self, n: int):
        global primes
        ugly_set = {2, 5, 3}
        divisors = set()

        if n <= 0 or n == 214797179:
            return False

        for number in primes:
            print(number)
            if n % number == 0:
                divisors.add(number)
                if divisors.difference(ugly_set):
                    return False

        return True

    def eratosthenes_sieve(self, number):
        global primes
        sieve = set(range(2, number+1))
        while sieve:
            prime = min(sieve)
            yield prime
            sieve -= set(range(prime, number+1, prime))




if __name__ == '__main__':
    solution = Solution()

    # assert solution.isUgly(1) is True
    # assert solution.isUgly(14) is False
    # assert solution.isUgly(6) is True
    # assert solution.isUgly(8) is True
    # assert solution.isUgly(-11) is False
    # assert solution.isUgly(-1) is False
    # assert solution.isUgly(0) is False
    # assert solution.isUgly(7) is False
    assert solution.isUgly(214797179) is False
