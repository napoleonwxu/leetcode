class Solution(object):
    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        ans = []
        self.dfs(sorted(nums), [], ans)
        return ans

    def dfs(self, nums, path, ans):
        ans.append(path)
        for i in xrange(len(nums)):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            self.dfs(nums[i+1:], path+[nums[i]], ans)
