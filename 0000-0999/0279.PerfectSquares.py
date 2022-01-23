import math
class Solution(object):
    def numSquares(self, n):
        """
        :type n: int
        :rtype: int
        """
        # BFS: 1772ms, beats 5.59%
        '''
        sqs = []
        i = 1
        sq = i * i
        while sq <= n:
            sqs.append(sq)
            i += 1
            sq = i * i

        count = 0
        toCheck = set([n])
        while 0 not in toCheck:
            temp = set()
            for check in toCheck:
                for sq in sqs:
                    if check < sq:
                        break
                    temp.add(check - sq)
            toCheck = temp
            count += 1

        return count
        '''
        # DFS: Time Limit Exceeded
        """
        dp = [0] * (n+1)
        for i in xrange(1, n+1):
            dp[i] = min(dp[i-j*j]+1 for j in xrange(1, int(math.sqrt(i))+1))
        return dp[n]
        """
        # DFS:
        dp = [0] * (n+1)
        for i in xrange(1, n+1):
            m = i
            j = 1
            while i - j*j >= 0:
                m = min(dp[i - j*j]+1, m)
                j += 1
            dp[i] = m
        return dp[n]

test = Solution()
print test.numSquares(5778)
'''
for i in xrange(15):
    print i, test.numSquares(i)
'''
