class Solution(object):
    def integerReplacement(self, n):
        """
        :type n: int
        :rtype: int
        """
        #3: math
        count = 0
        while n > 1:
            if n%2 == 0:
                n /= 2
            elif (n+1)%4 == 0 and n-1 != 2:
                n += 1
            else:
                n -= 1
            count += 1
        return count
        
        ''' #2: recursive, faster 
        self.dic = {1: 0}
        self.rec(n)
        return self.dic[n]
        '''
        ''' #1: recursive, slow
        if n == 1:
            return 0
        if n&1 == 0:
            return self.integerReplacement(n/2) + 1
        else:
            return min(self.integerReplacement(n-1), self.integerReplacement(n+1)) + 1
        '''

    def rec(self, n):
        if n in self.dic:
            return
        if n&1 == 0:
            self.rec(n/2)
            self.dic[n] = self.dic[n/2] + 1
        else:
            self.rec(n-1)
            self.rec(n+1)
            self.dic[n] = min(self.dic[n-1], self.dic[n+1]) + 1
        
        