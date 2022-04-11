class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        size = len(nums)
        i = 0
        while i < size:
            if nums[i] == 0:
                nums.append(nums.pop(i))
                size -= 1
            else:
                i += 1