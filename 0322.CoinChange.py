class Solution(object):
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        # DP
        dp = [0] + [float('inf')] * amount
        for i in xrange(1, amount+1):
            for c in coins:
                if c <= i:
                    dp[i] = min(dp[i], dp[i-c] + 1)
        if dp[amount] > amount:
            return -1
        else:
            return dp[amount]
        '''
        # BFS
        if amount == 0:
            return 0
        visited = [True] + [False] * amount
        queue1 = [0]
        queue2 = []
        ans = 0
        while queue1:
            ans += 1
            for v in queue1:
                for c in coins:
                    new = v + c
                    if new == amount:
                        return ans
                    elif new > amount:
                        continue
                    elif not visited[new]:
                        queue2.append(new)
                        visited[new] = True
            queue1, queue2 = queue2, []
        return -1
        '''
