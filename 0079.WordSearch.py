class Solution(object):
    def exist(self, board, word):
        """
        :type board: List[List[str]]
        :type word: str
        :rtype: bool
        """
        if not word:
            return True
        if not board:
            return False
        for i in xrange(len(board)):
            for j in xrange(len(board[0])):
                if self.dfs(board, word, i, j):
                    return True
        return False

    def dfs(self, board, word, i, j):
        if not word:
            return True
        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] != word[0]:
            return False
        temp = board[i][j]
        board[i][j] = '#'
        ans = self.dfs(board, word[1:], i-1, j) or self.dfs(board, word[1:], i+1, j) or self.dfs(board, word[1:], i, j-1) or self.dfs(board, word[1:], i, j+1)
        board[i][j] = temp
        return ans
