class Solution:
    def isPalindrome(self, s: str) -> bool:
        lPointer, rPointer = 0, len(s) - 1
        while lPointer < rPointer:
            if not s[lPointer].isalnum():
                lPointer += 1
            elif not s[rPointer].isalnum():
                rPointer -= 1
            else:
                if s[lPointer].lower() != s[rPointer].lower():
                    return False
                else:
                    lPointer += 1
                    rPointer -= 1
        return True