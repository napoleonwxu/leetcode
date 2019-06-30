class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: bool
        """
        left, right = 0, len(nums)-1
        while left <= right:
            mid = (left+right)/2
            if target == nums[mid]:
                return True
            '''
            if nums[left] == nums[mid] == nums[right]:
                return self.searchInOrder(nums, target, left, right)
            '''
            while left < mid and nums[left] == nums[mid]:
                left += 1
            if nums[left] <= nums[mid]:
                if nums[left] <= target < nums[mid]:
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                if nums[mid] < target <= nums[right]:
                    left = mid + 1
                else:
                    right = mid - 1
        return False

    def searchInOrder(self, nums, target, left, right):
        for i in xrange(left, right+1):
            if target == nums[i]:
                return True
        return False