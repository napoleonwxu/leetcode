class Solution(object):
    def fizzBuzz(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        return [str(i) if i%3 != 0 and i%5 != 0 else 'Fizz'*(i%3==0) + 'Buzz'*(i%5==0) for i in xrange(1, n+1)]
        ans = []
        for i in xrange(1, n+1):
            tmp = ''
            if i%3 == 0:
                tmp = 'Fizz'
            if i%5 == 0:
                tmp += 'Buzz'
            if not tmp:
                tmp = str(i)
            ans.append(tmp)
        return ans
        