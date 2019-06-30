class Solution(object):
    def wiggleSort(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        nums.sort()
        half = len(nums[::2])
        nums[::2], nums[1::2] = nums[:half][::-1], nums[half:][::-1]

nums = [1, 5, 1, 1, 6, 4]
#nums = [1, 3, 2, 2, 3, 1]
Solution().wiggleSort(nums)
print nums