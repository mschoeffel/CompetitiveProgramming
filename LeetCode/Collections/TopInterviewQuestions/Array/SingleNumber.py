class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        while len(nums) > 1:
            num = nums.pop()
            if num not in nums:
                return num
            else:
                nums.pop(nums.index(num))
        return nums[0]