class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        ans = [0] * (num+1)
        for i in xrange(1, num+1):
            ans[i] = ans[i>>1] + (i&1)
        return ans
        '''
        if num == 0:
            return [0]
        bits = int(math.floor(math.log(num)/math.log(2))) + 1
        ones = [0]
        for i in xrange(bits):
            ones += [c+1 for c in ones]
        return ones[:num+1]
        '''
