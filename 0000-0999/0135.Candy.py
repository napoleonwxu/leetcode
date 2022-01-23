class Solution(object):
    def candy(self, ratings):
        """
        :type ratings: List[int]
        :rtype: int
        """
        length = len(ratings)
        candys = [1] * length
        for i in xrange(1, length):
            if ratings[i] > ratings[i-1]:
                candys[i] = candys[i-1] + 1
        for i in xrange(length-1, 0, -1):
            if ratings[i-1] > ratings[i]:
                candys[i-1] = max(candys[i-1], candys[i]+1)
        return sum(candys)