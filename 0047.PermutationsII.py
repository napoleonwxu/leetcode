class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        ans = []
        self.dfs(sorted(nums), [], ans)
        return ans

    def dfs(self, nums, temp, ans):
        if not nums:
            ans.append(temp)
            return
        for i in xrange(len(nums)):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            self.dfs(nums[:i]+nums[i+1:], temp+[nums[i]], ans)
