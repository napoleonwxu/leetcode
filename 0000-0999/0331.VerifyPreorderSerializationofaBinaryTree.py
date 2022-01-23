class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        nodes = preorder.split(',')
        """
        stack = []
        for node in nodes:
            if node == '#':
                while len(stack) >= 2 and stack[-1] == '#' and stack[-2] != '#':
                    stack.pop()
                    stack.pop()
            stack.append(node)
        return stack == ['#']
        """
        diff = 1
        for node in nodes:
            '''
            diff -= 1
            if diff < 0:
                return False
            if node != '#':
                diff += 2
            '''
            if diff == 0:
                return False
            if node == '#':
                diff -= 1
            else:
                diff += 1
        return diff == 0

print Solution().isValidSerialization('9,3,4,#,#,1,#,#,2,#,6,#,#')