class Solution:
    def twoSum(self, nums, target):
        store = {}
        for i, number in enumerate(nums):
            search = target - number
            if search in store:
                return [store[search], i]
            else:
                store[number] = i