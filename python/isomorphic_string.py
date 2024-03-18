"""
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character, but a character may map to itself.
"""
class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        s_map = {}
        t_map = {}

        for s_letter, t_letter in zip(s, t):
            if s_letter in s_map and s_map[s_letter] != t_letter:
                return False

            if t_letter in t_map and t_map[t_letter] != s_letter:
                return False

            s_map[s_letter] = t_letter
            t_map[t_letter] = s_letter

        return True



if __name__ == '__main__':
    solution = Solution()

    assert solution.isIsomorphic(s='foo', t='bar') is False
    assert solution.isIsomorphic(s='egg', t='add') is True
    assert solution.isIsomorphic(s='badc', t='baba') is False
