class Solution(object):
    def combinationSum3(self, k, n):
        """
        :type k: int
        :type n: int
        :rtype: List[List[int]]
        """
        ans = []
        self.dfs(k, n, 1, [], ans)
        return ans

    def dfs(self, k, n, cur, temp, ans):
        if k == 0:
            if n == 0:
                ans.append(temp)
            return
        for i in xrange(cur, 10):
            if n < i:
                break
            self.dfs(k-1, n-i, i+1, temp+[i], ans)
