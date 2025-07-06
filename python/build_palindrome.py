"""
Task
Given a string str, find the shortest possible string which can be achieved by adding characters to the end of initial
string to make it a palindrome.

Example
For str = "abcdc", the output should be "abcdcba"
"""


def build_palindrome(st):
    step = 0
    new_word = st
    while not new_word == new_word[::-1]:
        step += 1
        new_word = st + st[step - 1 :: -1]

    return new_word


if __name__ == "__main__":
    assert build_palindrome("abcdc") == "abcdcba"
    assert build_palindrome("ababab") == "abababa"
