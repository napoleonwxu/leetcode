class Solution(object):
    def integerBreak(self, n):
        """
        :type n: int
        :rtype: int
        """
        if 1 <= n <= 3:
            return n - 1
        '''# math
        if n%3 == 0:
            return 3**(n/3)
        if n%3 == 1:
            return 3**(n/3-1) * 4
        if n%3 == 2:
            return 3**(n/3) * 2
        '''
        # DP
        dp = [0] * (n+1)
        dp[2] = 2
        dp[3] = 3
        dp[4] = 4
        for i in xrange(5, n+1):
            dp[i] = dp[i-3] * 3
        return dp[n]

