class Solution(object):
    def evalRPN(self, tokens):
        """
        :type tokens: List[str]
        :rtype: int
        """
        stack = []
        for n in tokens:
            if n not in '+-*/':
                stack.append(int(n))
            elif n in '+-*/' and len(stack) >= 2:
                b = stack.pop()
                a = stack.pop()
                if n == '+':
                    stack.append(a + b)
                elif n == '-':
                    stack.append(a - b)
                elif n == '*':
                    stack.append(a * b)
                elif b != 0:
                    stack.append(int(float(a) / b))
                    '''
                    if (a >= 0 and b > 0) or (a <= 0 and b < 0):
                        stack.append(a / b)
                    else:
                        stack.append(-(abs(a) / abs(b)))
                    '''
        return stack[0]

test = Solution()
print test.evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"])
