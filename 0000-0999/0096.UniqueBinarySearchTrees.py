class Solution(object):
    def numTrees(self, n):
        """
        :type n: int
        :rtype: int
        """
        ans = [0] * (n+1)
        ans[0] = ans[1] = 1
        for i in xrange(2, n+1):
            for j in xrange(i):
                ans[i] += ans[j] * ans[i-j-1]
        return ans[n]
