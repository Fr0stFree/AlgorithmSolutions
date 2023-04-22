"""
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
Given an integer n, return true if n is an ugly number.
"""
import time


class Solution:
    # Fuck this task I am out
    def isUgly(self, n: int):
        ugly_set = {2, 5, 3}
        divisors = set()

        for divisor in self.prime_generator():
            if divisor > n:
                return True

            if n % divisor == 0:
                divisors.add(divisor)
                n //= divisor
                if divisors.difference(ugly_set):
                    return False

            if divisor > 5 and not divisors:
                return False

    def is_prime(self, n):
        if n <= 1:
            return False

        for i in range(2, int(n ** 0.5) + 1):
            if n % i == 0:
                return False

        return True

    def prime_generator(self):
        num = 1
        while True:
            num += 1
            if self.is_prime(num):
                yield num


if __name__ == '__main__':
    solution = Solution()

    assert solution.isUgly(14) is False
    assert solution.isUgly(6) is True
    assert solution.isUgly(8) is True
    assert solution.isUgly(7) is False
