import re
class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = re.sub(r'[^A-Za-z0-9]', '', s)
        s = s.lower()
        lPointer = 0
        rPointer = len(s)-1
        while rPointer >= len(s)/2:
            if s[lPointer] != s[rPointer]:
                return False
            lPointer += 1
            rPointer -= 1
        return True