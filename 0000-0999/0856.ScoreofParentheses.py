class Solution(object):
    def scoreOfParentheses(self, S):
        """
        :type S: str
        :rtype: int
        """
        return eval(S.replace(')(', ')+(').replace('()', '1').replace(')', ')*2'))
        