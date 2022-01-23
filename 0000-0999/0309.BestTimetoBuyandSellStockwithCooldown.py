class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if len(prices) < 2:
            return 0
        preBuy, preSell = 0, 0
        buy, sell = -prices[0], 0
        for price in prices:
            preBuy = buy
            buy = max(preSell-price, preBuy)
            preSell = sell
            sell = max(preBuy+price, preSell)
        return sell
        