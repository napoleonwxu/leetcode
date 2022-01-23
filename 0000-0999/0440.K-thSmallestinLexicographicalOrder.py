class Solution(object):
    def findKthNumber(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: int
        """
        ans = 1
        k -= 1
        while k > 0:
            count = 0
            interval = [ans, ans+1]
            while interval[0] <= n:
                count += min(n+1, interval[1]) - interval[0]
                interval = [interval[0]*10, interval[1]*10]
            if k >= count:
                ans += 1
                k -= count
            else:
                ans *= 10
                k -= 1
        return ans
