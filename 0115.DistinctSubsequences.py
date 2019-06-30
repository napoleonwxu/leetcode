class Solution(object):
    def numDistinct(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        ''' # O(mn) time, O(mn) space
        m, n = len(t)+1, len(s)+1
        dp = [[0]*n for i in xrange(m)]
        for j in xrange(n):
            dp[0][j] = 1
        for i in xrange(1, m):
            for j in xrange(1, n):
                if t[i-1] == s[j-1]:
                    dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
                else:
                    dp[i][j] = dp[i][j-1]
        return dp[-1][-1]
        '''
        # O(mn) time, O(n) space
        m, n = len(t)+1, len(s)+1
        dp = [1]*n
        for i in xrange(1, m):
            pre = dp[0]
            dp[0] = 0
            for j in xrange(1, n):
                temp = dp[j]
                if t[i-1] == s[j-1]:
                    dp[j] = pre + dp[j-1]
                else:
                    dp[j] = dp[j-1]
                pre = temp
        return dp[-1]
