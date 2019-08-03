class Solution(object):
    def maxProfit(self, k, prices):
        """
        :type k: int
        :type prices: List[int]
        :rtype: int
        """
        if k == 0 or len(prices) <= 1:
            return 0
        if k >= len(prices)/2:
            profit = 0
            for i in xrange(1, len(prices)):
                profit += max(prices[i]-prices[i-1], 0)
            return profit
        '''
        buy = [-65536]*k
        sel = [0]*k
        for i in xrange(len(prices)):
            buy[0] = max(buy[0], -prices[i])
            sel[0] = max(sel[0], buy[0]+prices[i])
            for j in xrange(1, k):
                buy[j] = max(buy[j], sel[j-1]-prices[i])
                sel[j] = max(sel[j], buy[j]+prices[i])
        '''
        buy = [-float('inf')]*(k+1)
        sel = [0]*(k+1)
        for i in xrange(len(prices)):
            for j in xrange(1,k+1):
                buy[j] = max(buy[j], sel[j-1]-prices[i])
                sel[j] = max(sel[j], buy[j]+prices[i])
                
        return max(sel)