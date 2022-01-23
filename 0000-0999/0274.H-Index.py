class Solution(object):
    def hIndex(self, citations):
        """
        :type citations: List[int]
        :rtype: int
        """
        #return sum(i < c for i, c in enumerate(sorted(citations, reverse = True)))
        ''' # O(logN)
        citations.sort()
        l = len(citations)
        for i in xrange(l):
            if citations[i] >= l-i:
                return l-i
        return 0
        '''
        # O(N)
        l = len(citations)
        array = [0] * (l+1)
        for c in citations:
            array[min(l, c)] += 1
            '''
            if c > l:
                array[l] += 1
            else:
                array[c] += 1
            '''
        t = 0
        for i in xrange(l, -1, -1):
            t += array[i]
            if t >= i:
                return i
        return 0
