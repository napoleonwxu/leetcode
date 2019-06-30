class Solution(object):
    def hIndex(self, citations):
        """
        :type citations: List[int]
        :rtype: int
        """
        '''
        if not citations or citations[-1] == 0:
            return 0
        cLen = len(citations)
        l, r = 0, cLen-1
        while l < r:
            m = (l+r) / 2
            if citations[m] < cLen-m:
                l = m + 1
            else:
                r = m
        return cLen-l
        '''
        n = len(citations)
        '''
        l, r = 0, n
        while l < r:
            m = (l+r)/2
            if citations[m] == n-m:
                return n-m
            elif citations[m] < n-m:
                l = m + 1
            else:
                r = m
        return n-l
        '''
        l, r = 0, n-1
        while l <= r:
            m = (l+r) / 2
            if citations[m] < n-m:
                l = m + 1
            else:
                r = m - 1
        return n-l