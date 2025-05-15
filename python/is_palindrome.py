"""
Given an array of integers nums and an integer target, return indices of the two numbers such that
they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
"""


class Solution:
    def isPalindrome(self, x: int) -> bool:
        for index, _ in enumerate((palindrome := str(x))):
            if palindrome[index] != palindrome[-index - 1]:
                return False
        return True


if __name__ == "__main__":
    solution = Solution()

    assert solution.isPalindrome(123) is False
    assert solution.isPalindrome(121) is True
