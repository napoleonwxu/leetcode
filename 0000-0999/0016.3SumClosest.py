class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        if len(nums) < 3:
            return
        ans = sum(nums[:3])
        nums.sort()
        for i in xrange(len(nums)-2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            l, r = i+1, len(nums)-1
            while l < r:
                temp = nums[i] + nums[l] + nums[r]
                if temp == target:
                    return target
                if abs(temp-target) < abs(ans-target):
                    ans = temp
                if temp < target:
                    l += 1
                    '''
                    while l < r and nums[l] == nums[l-1]:
                        l += 1
                    '''
                else:
                    r -= 1
                    '''
                    while l < r and nums[r] == nums[r+1]:
                        r -= 1
                    '''
        return ans
