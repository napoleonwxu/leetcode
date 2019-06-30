class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        if len(nums) <= 1:
            return
        i = len(nums)-2
        while i >= 0 and nums[i] >= nums[i+1]:
            i -= 1
        if i >= 0:
            j = i+1
            while j < len(nums)-1 and nums[j+1] > nums[i]:
                j += 1
            nums[i], nums[j] = nums[j], nums[i]
            nums[i+1:] = reversed(nums[i+1:])
            #nums[i+1:] = sorted(nums[i+1:])
        else:
            nums.reverse()
