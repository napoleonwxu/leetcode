class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        if n < 1:
            return
        ans = []
        #self.dfs(2*n, n, n, '', ans)
        self.dfs2(n, n, '', ans)
        return ans

    def dfs(self, n, left, right, path, ans):
        if left > right:
            return
        if n == 0:
            ans.append(path)
            return
        if left > 0:
            self.dfs(n-1, left-1, right, path+'(', ans)
        if right > 0:
            self.dfs(n-1, left, right-1, path+')', ans)

    def dfs2(self, left, right, path, ans):
        if left > right:
            return
        if left == 0 and right == 0:
            ans.append(path)
            return
        if left > 0:
            self.dfs2(left-1, right, path+'(', ans)
        if right > 0:
            self.dfs2(left, right-1, path+')', ans)
