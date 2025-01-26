class Solution(object):
    def maxEnvelopes(self, envelopes):
        """
        :type envelopes: List[List[int]]
        :rtype: int
        """
        if not envelopes:
            return 0
        envs = sorted(envelopes, key = lambda x: (x[0], -x[1]))
        end = [0] * len(envs)
        high = 0
        # O(nlogn)
        for w, h in envs:   # O(n)
            #i = bisect.bisect_left(end, h, 0, high) # O(logn)
            i = self.my_bisect_left(end, h, 0, high)    # O(logn)
            end[i] = h
            if i == high:
                high += 1
        return high
        
        ''' # O(n^2), TLE
        env = sorted(envelopes)
        dp = [1] * len(env)
        for i in xrange(1, len(env)):
            for j in xrange(i):
                if env[i][0] > env[j][0] and env[i][1] > env[j][1]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return max(dp)
        '''
        
    def my_bisect_left(self, nums, target, l, r):
        while l < r:
            m = l + (r-l)/2
            if target <= nums[m]:
                r = m
            else:
                l = m + 1
        return l
                