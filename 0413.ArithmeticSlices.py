class Solution(object):
    def numberOfArithmeticSlices(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        cur = ans = 0
        for i in xrange(2, len(A)):
            if A[i]-A[i-1] == A[i-1]-A[i-2]:
                cur += 1
                ans += cur
            else:
                cur = 0
        return ans
        '''
        n = len(A)
        ans = 0
        for l in xrange(n-2):
            d = A[l+1] - A[l]
            for r in xrange(l+2, n):
                if A[r]-A[r-1] == d:
                    ans += 1
                else:
                    break
        return ans
        '''