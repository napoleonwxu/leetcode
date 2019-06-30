class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        # O(m+n) time, fast: 58%
        # http://yucoding.blogspot.com/2013/02/leetcode-question-123-wildcard-matching.html
        i = j = 0
        match = 0
        star = -1
        while i < len(s):
            if j < len(p) and (s[i] == p[j] or p[j] == '?'):
                i += 1
                j += 1
            elif j < len(p) and p[j] == '*':
                match = i
                star = j
                j += 1
            elif star != -1:
                j = star + 1
                match += 1
                i = match
            else:
                return False
        while j < len(p) and p[j] == '*':
            j += 1
        return j == len(p)

        ''' # O(mn) time + O(n) space, 30%
        n = len(s)
        if len(p) - p.count('*') > n:
            return False
        dp = [True] + [False] * n
        for ch in p:
            temp = [dp[0] and ch == '*']
            if ch == '?':
                temp += dp[:n]
            elif ch == '*':
                for j in xrange(n):
                    temp.append(temp[-1] or dp[j+1])
            else:
                for j in xrange(n):
                    temp.append(dp[j] and ch == s[j])
            dp = temp
        return dp[-1]
        '''

        ''' # O(mn) time + O(mn) space, slow: 8%
        m, n = len(p), len(s)
        if m - p.count('*') > n:  #avoid TLE
            return False
        dp = [[False]*(n+1) for i in xrange(m+1)]
        dp[0][0] = True
        for i in xrange(1, m+1):
            dp[i][0] = dp[i-1][0] and p[i-1] == '*'
        for i in xrange(1, m+1):
            for j in xrange(1, n+1):
                if p[i-1] == '?':
                    dp[i][j] = dp[i-1][j-1]
                elif p[i-1] == '*':
                    dp[i][j] = dp[i][j-1] or dp[i-1][j]
                else:
                    dp[i][j] = dp[i-1][j-1] and p[i-1]==s[j-1]
        return dp[-1][-1]
        '''