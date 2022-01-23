class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        return self.binSearch(nums, target, 0, len(nums)-1)

    def binSearch(self, nums, target, left, right):
        if left > right:
            return left
        mid = (left+right)/2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
        return self.binSearch(nums, target, left, right)
