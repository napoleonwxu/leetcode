class Solution(object):
    def diffWaysToCompute(self, input):
        """
        :type input: str
        :rtype: List[int]
        """
        if input.isdigit():
            return [int(input)]
        ans = []
        for i in xrange(len(input)):
            if input[i] in '+-*':
                left = self.diffWaysToCompute(input[:i])
                right = self.diffWaysToCompute(input[i+1:])
                for n in left:
                    for m in right:
                        if input[i] == '+':
                            ans.append(n+m)
                        elif input[i] == '-':
                            ans.append(n-m)
                        else:
                            ans.append(n*m)
        '''
        if not ans:
            ans.append(int(input))
        '''
        return ans
