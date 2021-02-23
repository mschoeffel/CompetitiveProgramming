import re
class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = re.sub(r'[^A-Za-z0-9]', '', s)
        s = s.lower()
        sPointer = 0
        tPointer = len(s)-1
        while tPointer >= len(s)/2:
            if s[sPointer] != s[tPointer]:
                return False
            sPointer += 1
            tPointer -= 1
        return True