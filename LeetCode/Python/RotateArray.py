class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        last = len(nums)-1
        k = k % len(nums)
        while k > 0:
            nums.insert(0, nums.pop(last)) 
            k-=1