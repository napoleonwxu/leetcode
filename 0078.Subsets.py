class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        '''
        ans = [[]]
        for num in sorted(nums):
            ans += [t+[num] for t in ans]
        return ans
        '''
        '''
        ans = []
        self.dfs(sorted(nums), 0, [], ans)
        return ans
        '''
        '''
        ans = []
        nums.sort()
        for i in xrange(1<<len(nums)):
            temp = []
            for j in xrange(len(nums)):
                if i>>j & 1: #if i & 1<<j:
                    temp.append(nums[j])
            ans.append(temp)
        return ans
        '''
        ans = [[] for i in xrange(2**len(nums))]
        nums.sort()
        for i in xrange(len(nums)):
            for j in xrange(len(ans)):
                if j>>i & 1:
                    ans[j].append(nums[i])
        return ans

    def dfs(self, nums, index, temp, ans):
        ans.append(temp)
        for i in xrange(index, len(nums)):
            self.dfs(nums, i+1, temp+[nums[i]], ans)