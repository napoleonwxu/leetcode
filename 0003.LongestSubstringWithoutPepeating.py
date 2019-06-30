class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if len(s) <= 1:
            return len(s)
        ans = 0
        start = 0
        d = {}
        for i, ch in enumerate(s):
            if ch in d:
                ans = max(ans, i-start)
                start = max(start, d[ch]+1)
            d[ch] = i
        return max(ans, len(s)-start)
        '''
        ans = 1
        l, r = 0, 1
        while r < len(s):
            while l < r and s[r] in s[l:r]:
                l += 1
            if l < r:
                ans = max(ans, r-l+1)
            r += 1
        return ans
        '''

