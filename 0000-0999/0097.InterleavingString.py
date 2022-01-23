class Solution(object):
    def isInterleave(self, s1, s2, s3):
        """
        :type s1: str
        :type s2: str
        :type s3: str
        :rtype: bool
        """
        ''' # O(mn) time; O(mn) space
        m, n = len(s1), len(s2)
        if m + n != len(s3):
            return False
        dp = [[False]*(n+1) for i in xrange(m+1)]
        dp[0][0] = True
        for i in xrange(1, m+1):
            dp[i][0] = dp[i-1][0] and s1[i-1] == s3[i-1]
        for j in xrange(1, n+1):
            dp[0][j] = dp[0][j-1] and s2[j-1] == s3[j-1]
        for i in xrange(1, m+1):
            for j in xrange(1, n+1):
                dp[i][j] = (dp[i-1][j] and s1[i-1] == s3[i+j-1]) or (dp[i][j-1] and s2[j-1] == s3[i+j-1])
        return dp[-1][-1]
        '''
        # O(mn) time; O(n) space
        m, n = len(s1), len(s2)
        if m + n != len(s3):
            return False
        dp = [False] * (n+1)
        dp[0] = True
        for j in xrange(1, n+1):
            dp[j] = dp[j-1] and s2[j-1] == s3[j-1]
        for i in xrange(1, m+1):
            dp[0] = dp[0] and s1[i-1] == s3[i-1]
            for j in xrange(1, n+1):
                dp[j] = (dp[j] and s1[i-1] == s3[i+j-1]) or (dp[j-1] and s2[j-1] == s3[i+j-1])
        return dp[-1]

        ''' # TLE
        if not s1 and not s2 and not s3:
            return True
        if len(s1) + len(s2) != len(s3):
            return False
        if s1 and s2 and s3 and s1[0] == s2[0] == s3[0]:
            return self.isInterleave(s1[1:], s2, s3[1:]) or self.isInterleave(s1, s2[1:], s3[1:])
        elif s1 and s3 and s1[0] == s3[0]:
            return self.isInterleave(s1[1:], s2, s3[1:])
        elif s2 and s3 and s2[0] == s3[0]:
            return self.isInterleave(s1, s2[1:], s3[1:])
        return False
        '''