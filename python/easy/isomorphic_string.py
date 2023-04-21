"""
Given two strings s and t, determine if they are isomorphic.
Two strings s and t are isomorphic if the characters in s can be replaced to get t.
All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character, but a character may map to itself.
"""
class Solution:
    # not solved yet
    def isIsomorphic(self, s: str, t: str) -> bool:
        s_previous, t_previous = None, None

        for s_letter, t_letter in zip(s, t):
            if any([s_letter == s_previous and t_letter != t_previous,
                    s_letter != s_previous and t_letter == t_previous]):
                return False
            s_previous, t_previous = s_letter, t_letter

        return True



if __name__ == '__main__':
    solution = Solution()

    # assert solution.isIsomorphic(s='foo', t='bar') is False
    # assert solution.isIsomorphic(s='egg', t='add') is True
    assert solution.isIsomorphic(s='badc', t='baba') is False
