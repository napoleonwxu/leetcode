#coding: utf-8
'''
https://leetcode.com/discuss/99229/leetbook-2种思路中文详（附解释，图示，代码）
'''

class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        dp = [[False]*(len(s)+1) for _ in xrange(len(p)+1)]
        dp[0][0] = True
        for i in xrange(2, len(p)+1):
            dp[i][0] = dp[i-2][0] and p[i-1] == '*'
        for i in xrange(1, len(p)+1):
            for j in xrange(1, len(s)+1):
                if p[i-1] != '*':
                    dp[i][j] = dp[i-1][j-1] and (p[i-1] == s[j-1] or p[i-1] == '.')
                else:
                    dp[i][j] = dp[i-2][j] or dp[i-1][j]
                    if p[i-2] == s[j-1] or p[i-2] == '.':
                        dp[i][j] |= dp[i][j-1]
        return dp[-1][-1]

test = Solution()
print test.isMatch('a', 'a*')