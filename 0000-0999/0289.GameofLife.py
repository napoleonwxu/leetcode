class Solution(object):
    def gameOfLife(self, board):
        """
        :type board: List[List[int]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        // O(mn) + O(1), cool
        m = len(board)
        n = len(board[0])
        for i in xrange(m):
            for j in xrange(n):
                self.judge(board, m, n, i, j)
        for i in xrange(m):
            for j in xrange(n):
                board[i][j] >>= 1

    def judge(self, board, m, n, i, j):
        live = -board[i][j]
        for ii in xrange(max(i-1, 0), min(i+2, m)):
            for jj in xrange(max(j-1, 0), min(j+2, n)):
                live += board[ii][jj] & 1
        if live == 3 or (board[i][j] == 1 and live == 2):
            board[i][j] += 2

