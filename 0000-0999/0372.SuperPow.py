class Solution(object):
    def superPow(self, a, b):
        """
        :type a: int
        :type b: List[int]
        :rtype: int
        """
        #1: naive
        #return pow(a, int(''.join(map(str, b))), 1337)
        
        #2: a^1234%k = (a^1230%k)*(a^4%k)%k = [(a^123%k)^10%k]*(a^4%k)%k
        if not b:
            return 1
        return pow(a, b.pop(), 1337) * self.superPow(pow(a, 10, 1337), b) % 1337
        
        #3: Euler's theorem, phi(1337) = phi(7) * phi(191) = 6 * 190 = 1140
        '''
        if a % 1337 == 0:
            return 0
        return pow(a, reduce(lambda x, y: (x*10 + y) % 1140, b) + 1140, 1337)
        '''