class Solution(object):
    def convertToTitle(self, n):
        """
        :type n: int
        :rtype: str
        """
        '''
        s = ''
        while n > 0:
            s = chr(ord('A')+(n-1)%26) + s
            n = (n-1)/26
        return s
        '''
        if n == 0:
            return ''
        else:
            return self.convertToTitle((n-1)/26) + chr(ord('A')+(n-1)%26)