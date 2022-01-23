class Solution(object):
    def countBattleships(self, board):
        """
        :type board: List[List[str]]
        :rtype: int
        """
        if not board:
            return 0
        ans = 0
        m, n = len(board), len(board[0])
        for i in xrange(m):
            for j in xrange(n):
                if board[i][j] == 'X' and (i==0 or board[i-1][j]!='X') and (j==0 or board[i][j-1]!='X'):
                    ans += 1
        return ans
