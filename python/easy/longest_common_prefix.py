"""
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string
"""
from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        strs.sort(key=lambda word: len(word))
        common_word = strs[0]

        for word in strs:
            if word == common_word: continue

            if not word:
                return ''

            for index, (common_letter, letter) in enumerate(zip(common_word, word)):
                if common_letter != letter:
                    common_word = word[:index]
                    break

        return common_word




if __name__ == '__main__':
    solution = Solution()

    assert solution.longestCommonPrefix(["flower", "flow", "flight"]) == "fl"
    assert solution.longestCommonPrefix(["dog", "racecar", "car"]) == ""
    assert solution.longestCommonPrefix(["ab", "a"]) == "a"
