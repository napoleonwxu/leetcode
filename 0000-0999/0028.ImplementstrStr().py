class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        '''
        if not needle:
            return 0
        '''
        k = len(needle)
        for i in xrange(len(haystack)-k+1):
            if haystack[i:i+k] == needle:
                return i
        return -1