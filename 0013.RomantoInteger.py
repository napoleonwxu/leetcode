class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return
        d = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
        ans = 0
        pre = s[0]
        for i in xrange(1, len(s)):
            if d[s[i]] > d[pre]:
                ans -= d[pre]
            else:
                ans += d[pre]
            pre = s[i]
        '''
        for i in xrange(1, len(s)):
            if d[s[i]] > d[s[i-1]]:
                ans -= d[s[i-1]]
            else:
                ans += d[s[i-1]]
        '''
        ans += d[s[-1]]
        return ans
