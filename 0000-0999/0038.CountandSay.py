class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        if n < 1:
            return
        nStr = '1'
        for _ in xrange(n-1):
            '''
            ans = ''
            i = 0
            while i < len(nStr):
                t = 1
                while i < len(nStr)-1 and nStr[i] == nStr[i+1]:
                    t += 1
                    i += 1
                ans += str(t) + nStr[i]
                i += 1
            nStr = ans
            '''
            nStr = self.helper(nStr)
        return nStr

    def helper(self, nStr):
        ans = ''
        i = 0
        while i < len(nStr):
            t = 1
            while i < len(nStr)-1 and nStr[i] == nStr[i+1]:
                t += 1
                i += 1
            ans += str(t) + nStr[i]
            i += 1
        return ans
