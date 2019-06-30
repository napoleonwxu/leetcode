class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        left = self.searchLeft(nums, target, 0, len(nums)-1)
        right = self.searchRight(nums, target, 0, len(nums)-1)
        return [left, right]

    def searchLeft(self, nums, target, left, right):
        if left > right:
            return -1
        mid = (left+right) / 2
        if nums[mid] == target:
            if mid == 0 or (mid > 0 and nums[mid-1] != target):
                return mid
            else:
                right = mid - 1
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
        return self.searchLeft(nums, target, left, right)

    def searchRight(self, nums, target, left, right):
        if left > right:
            return -1
        mid = (left+right) / 2
        if nums[mid] == target:
            if mid == len(nums)-1 or (mid < len(nums)-1 and nums[mid+1] != target):
                return mid
            else:
                left = mid + 1
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
        return self.searchRight(nums, target, left, right)
