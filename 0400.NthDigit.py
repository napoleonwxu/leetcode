class Solution(object):
    def findNthDigit(self, n):
        """
        :type n: int
        :rtype: int
        """
        start = 1
        length = 1
        count = 9
        while n > length * count:
            n -= length * count
            start *= 10
            length += 1
            count *= 10
        num = start + (n-1)/length
        return int(str(num)[(n-1)%length])
        