class Solution(object):
    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        left, right = 0, len(nums)-1
        while left < right:
            mid = (left+right) / 2
            if nums[mid] > nums[right]:
                left = mid + 1
            elif nums[mid] < nums[right]:
                right = mid
            else:
                right -= 1
        '''
        while left < right:
            while left < right and nums[left] == nums[left+1]:
                left += 1
            while right > left and nums[right] == nums[right-1]:
                right -= 1
            if left == right:
                break
            mid = (left+right) / 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        '''
        return nums[left]