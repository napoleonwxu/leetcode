class Solution(object):
    def minDistance(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: int
        """
        '''# O(mn) Space
        l1 = len(word1) + 1
        l2 = len(word2) + 1
        dp = [[0]*l2 for i in xrange(l1)]
        for i in xrange(1, l1):
            dp[i][0] = i
        for j in xrange(1, l2):
            dp[0][j] = j
        for i in xrange(1, l1):
            for j in xrange(1, l2):
                if word1[i-1] == word2[j-1]:
                    dp[i][j] = dp[i-1][j-1]
                else:
                    dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
        return dp[-1][-1]
        '''
        # O(n) Space
        m, n = len(word1)+1, len(word2)+1
        dp = range(n)
        for i in xrange(1, m):
            pre = dp[0]
            dp[0] = i
            for j in xrange(1, n):
                temp = dp[j]
                if word1[i-1] == word2[j-1]:
                    dp[j] = pre
                else:
                    dp[j] = min(pre, dp[j-1], dp[j]) + 1
                pre = temp
        return dp[-1]