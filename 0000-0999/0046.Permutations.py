class Solution(object):
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        # return list(itertools.permutations(nums))
        ans = []
        self.dfs(nums, [], ans)
        return ans

    def dfs(self, nums, temp, ans):
        if not nums:
            ans.append(temp)
            return
        for i in xrange(len(nums)):
            self.dfs(nums[:i]+nums[i+1:], temp+[nums[i]], ans)
            '''
            nums[0], nums[i] = nums[i], nums[0]
            self.dfs(nums[1:], temp+[nums[0]], ans)
            '''
