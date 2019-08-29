class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        #return str(x) == str(x)[::-1]
        if x<0 or (x!=0 and x%10==0):
            return False
        y = 0
        while x > y:
            y = 10*y + x%10
            x /= 10
        return y == x or y/10 == x
        