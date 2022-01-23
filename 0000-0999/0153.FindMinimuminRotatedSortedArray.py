class Solution(object):
    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        left, right = 0, len(nums)-1
        while left < right:
            '''# slow
            mid = (left+right) / 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
            '''
            # fast
            if nums[left] < nums[right]:
                break
            mid = (left+right) / 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        return nums[left]
