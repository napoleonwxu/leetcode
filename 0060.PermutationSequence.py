class Solution(object):
    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        m = n
        n_ = 1
        while m > 1:
            n_ *= m
            m -= 1
        if k > n_:
            return
        return self.permute(n, n_, k-1, [i for i in xrange(1, n+1)])

    def permute(self, n, n_, k_, nums):
        #if len(nums) == 1:
        if n == 1:
            return str(nums[0])
        n_ /= n
        i = k_/n_
        n -= 1
        k_ %= n_
        '''
        if i == len(nums)-1:
            return str(nums[i]) + self.permute(n, n_, k_, nums[:i])
        '''
        return str(nums[i]) + self.permute(n, n_, k_, nums[:i]+nums[i+1:])
