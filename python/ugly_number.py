"""
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
Given an integer n, return true if n is an ugly number.
"""
import time


class Solution:
    def isUgly(self, n: int):
        if -1 <= n <= 0:
            return False

        for divisor in (2, 3, 5):
            while n % divisor == 0:
                n //= divisor

        return n == 1

    def prime_generator(self):
        num = 1
        while True:
            num += 1
            if self.is_prime(num):
                yield num

    @staticmethod
    def is_prime(n):
        if n <= 1:
            return False

        for i in range(2, int(n ** 0.5) + 1):
            if n % i == 0:
                return False

        return True




if __name__ == '__main__':
    solution = Solution()

    assert solution.isUgly(14) is False
    assert solution.isUgly(6) is True
    assert solution.isUgly(8) is True
    assert solution.isUgly(7) is False
    assert solution.isUgly(0) is False
    assert solution.isUgly(-1) is False
    assert solution.isUgly(1) is True

