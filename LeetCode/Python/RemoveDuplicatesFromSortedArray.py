class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        size = len(nums)
        if size <= 1:
            return size
        i = 1
        lastNum = nums[0]
        while(i < size):
            e = nums[i]
            if e == lastNum:
                nums.pop(i)
                size -= 1
            else:
                lastNum = e
                i += 1
        return len(nums)