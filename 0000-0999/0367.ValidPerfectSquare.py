class Solution(object):
    def isPerfectSquare(self, num):
        """
        :type num: int
        :rtype: bool
        """
        #3: Integer Newton
        r = num
        while r*r > num:
            r = (r + num/r)/2
        return r*r == num
        '''#2: 1 + 3 + 5 + ... + (2n-1) = (2n-1+1)n/2 = n*n
        i = 1
        while num > 0:
            num -= i
            i += 2
        return num == 0
        '''
        '''#1: Binary Search
        l, r = 1, num
        while l <= r:
            m = l + (r-l)/2
            if m*m == num:
                return True
            elif m*m < num:
                l = m + 1
            else:
                r = m - 1
        return False
        '''