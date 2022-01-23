class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        length = len(s)
        if length <= 1:
            return s
        base, maxLen = 0, 1
        i = 0
        while i < length:
            if length - base <= maxLen/2:
                break
            l = r = i
            while r < length-1 and s[r+1] == s[r]:
                r += 1
            i = r + 1
            while l > 0 and r < length-1 and s[l-1] == s[r+1]:
                l -= 1
                r += 1
            if r-l+1 > maxLen:
                base, maxLen = l, r-l+1
        return s[base:base+maxLen]

print Solution().longestPalindrome('dfgggfk')