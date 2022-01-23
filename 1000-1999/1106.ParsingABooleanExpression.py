class Solution(object):
    def parseBoolExpr(self, expression, t=True, f=False):
        """
        :type expression: str
        :rtype: bool
        """
        # "|(&(t,f,t),!(t))"
        return eval(expression.replace('!', 'not &').replace('&(', 'all([').replace('|(', 'any([').replace(')', '])'))
        # 'any([all([t,f,t]),not all([t])])'
        #return eval(expression.replace('!', 'not |').replace('&(', 'all([').replace('|(', 'any([').replace(')', '])'))
        # 'any([all([t,f,t]),not any([t])])'
        