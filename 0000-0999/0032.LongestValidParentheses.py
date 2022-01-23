class Solution(object):
    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return 0
        #dp = [0] * (len(s)+1)
        dp = [0] * len(s)
        stack = []
        for i in xrange(len(s)):
            if s[i] == '(':
                stack.append(i)
            elif stack:
                l = stack.pop()
                #dp[i+1] = dp[l] + i - l + 1
                dp[i] = dp[l-1] + i - l + 1
        return max(dp)
