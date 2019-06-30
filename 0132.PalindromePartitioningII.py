class Solution(object):
    def minCut(self, s):
        """
        :type s: str
        :rtype: int
        """
        l = len(s)
        #ans = [x for x in xrange(-1, l)]
        ans = range(-1, l)
        '''
        for i in xrange(l):
            for j in xrange(i, l):
                if s[i:j] == s[j:i:-1]:
                    ans[j+1] = min(ans[j+1], ans[i]+1)
        return ans[-1]
        '''
        for i in xrange(l):
            r1 = r2 = 0
            while i-r1 >= 0 and i+r1 < l and s[i-r1] == s[i+r1]:
                ans[i+r1+1] = min(ans[i+r1+1], ans[i-r1]+1)
                r1 += 1
            while i-r2 >= 0 and i+r2+1 < l and s[i-r2] == s[i+r2+1]:
                ans[i+r2+2] = min(ans[i+r2+2], ans[i-r2]+1)
                r2 += 1
        return ans[-1]
