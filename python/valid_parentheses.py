"""
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
"""


class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for element in s:
            if element in ("(", "{", "["):
                stack.append(element)
                continue

            try:
                previous = stack.pop()
            except IndexError:
                return False

            if not any([
                previous == "{" and element == "}",
                previous == "[" and element == "]",
                previous == "(" and element == ")",
            ]):
                return False

        return not stack


if __name__ == '__main__':
    s = Solution()

    assert s.isValid("()") is True
    assert s.isValid(")(") is False
