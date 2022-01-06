class Solution(object):
    def getMoneyAmount(self, n):
        """
        :type n: int
        :rtype: int
        """
        dp = [[0]*(n+1) for i in xrange(n+1)]
        for i in xrange(n, 0, -1):
            for j in xrange(i+1, n+1):
                dp[i][j] = min(k + max(dp[i][k-1], dp[k+1][j]) for k in xrange(i, j))
                '''
                dp[i][j] = float('inf')
                for k in xrange(i, j):
                    dp[i][j] = min(dp[i][j], k + max(dp[i][k-1], dp[k+1][j]))
                '''
        return dp[1][n]