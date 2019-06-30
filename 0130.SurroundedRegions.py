class Solution(object):
    def solve(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        if not board:
            return
        row = len(board)
        col = len(board[0])
        queue = [ij for i in xrange(row) for ij in [(i, 0), (i, col-1)]]
        queue += [ij for j in xrange(1, col-1) for ij in [(0, j), (row-1, j)]]
        while queue:
            i, j = queue.pop(0)
            if 0 <= i < row and 0 <= j < col and board[i][j] == 'O':
                board[i][j] = 'A'
                queue += [(i-1, j), (i+1, j), (i, j-1), (i, j+1)]
        '''
        for i in xrange(row):
            for j in xrange(col):
                board[i][j] = 'XO'[board[i][j] == 'A']
        '''
        for r in board:
            for j, ch in enumerate(r):
                r[j] = 'XO'[ch == 'A']
