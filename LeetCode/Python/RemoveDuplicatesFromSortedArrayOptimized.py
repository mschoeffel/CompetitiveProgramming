from collections import Counter

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nums[:] = list(Counter(nums).keys())
        return len(nums)