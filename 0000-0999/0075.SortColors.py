class Solution(object):
    def sortColors(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        i = front = 0
        back = len(nums) - 1
        while i<=back:
            if nums[i] == 0:
                if i > front:
                    nums[front], nums[i] = nums[i], nums[front]
                i += 1
                front += 1
            elif nums[i] == 2:
                nums[back], nums[i] = nums[i], nums[back]
                back -= 1
            else:
                i += 1
