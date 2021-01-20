class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        if len(nums) <= 1:
            return False
        visited = []
        for num in nums:
            if num in visited:
                return True
            visited.append(num)
        return False
