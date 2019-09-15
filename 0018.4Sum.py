class Solution(object):
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        if len(nums) < 4:
            return []
        ans = []
        nums.sort()
        self.nSum(nums, 4, target, [], ans)
        return ans
        '''
        for i in xrange(len(nums)-3):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            for j in xrange(i+1, len(nums)-2):
                if j > i+1 and nums[j] == nums[j-1]:
                    continue
                l, r = j+1, len(nums)-1
                while l < r:
                    t = nums[i] + nums[j] + nums[l] + nums[r]
                    if t < target:
                        l += 1
                    elif t > target:
                        r -= 1
                    else:
                        ans.append([nums[i], nums[j], nums[l], nums[r]])
                        while l < r and nums[l] == nums[l+1]:
                            l += 1
                        while l < r and nums[r] == nums[r-1]:
                            r -= 1
                        l += 1
                        r -= 1
        return ans
        '''
    def twoSum(self, nums, target, temp, ans):
        l, r = 0, len(nums)-1
        while l < r:
            t = nums[l] + nums[r]
            if t < target:
                l += 1
            elif t > target:
                r -= 1
            else:
                ans.append(temp + [nums[l], nums[r]])
                while l < r and nums[l] == nums[l+1]:
                    l += 1
                while l < r and nums[r] == nums[r-1]:
                    r -= 1
                l += 1
                r -= 1
        return
        
    def nSum(self, nums, N, target, temp, ans):
        if N < 2 or len(nums) < N:
            return
        if N == 2:
            self.twoSum(nums, target, temp, ans)
        else:
            if target > N*nums[-1]:
                return
            for i in xrange(len(nums)-N+1):
                if target < N*nums[i]:
                    break
                if i > 0 and nums[i] == nums[i-1]:
                    continue
                self.nSum(nums[i+1:], N-1, target-nums[i], temp+[nums[i]], ans)
        return
    