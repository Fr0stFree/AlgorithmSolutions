"""
Given a valid (IPv4) IP address, return a defanged version of that IP address.
A defanged IP address replaces every period "." with "[.]".
"""
class Solution:
    def defangIPaddr(self, address: str) -> str:
        return '[.]'.join(address.split('.'))


if __name__ == '__main__':
    solution = Solution()
    
    assert solution.defangIPaddr("255.100.50.0") == "255[.]100[.]50[.]0"
