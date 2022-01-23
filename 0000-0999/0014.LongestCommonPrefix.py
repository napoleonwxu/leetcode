#import os
class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        #return os.path.commonprefix(strs)
        if not strs:
            return ''
        minLen = len(strs[0])
        for s in strs[1:]:
            minLen = min(len(s), minLen)
        '''
        lenStrs = [len(s) for s in strs]
        minLen = min(lenStrs)
        '''
        ans = strs[0][:minLen]
        for s in strs[1:]:
            for i in xrange(minLen):
                if ans[i] != s[i]:
                    minLen = i
                    break
            if minLen == 0:
                return ''
        return ans[:minLen]