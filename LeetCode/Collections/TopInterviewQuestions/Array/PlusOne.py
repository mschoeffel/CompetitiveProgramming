class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        result = []
        overflow = 1
        i = len(digits)-1
        while i >= 0:
            num = digits[i] + overflow
            if num >= 10:
                overflow = 1
                num -= 10
            else:
                overflow = 0
            result.insert(0, num)
            i -= 1
        if overflow > 0:
            result.insert(0, overflow)
        return result