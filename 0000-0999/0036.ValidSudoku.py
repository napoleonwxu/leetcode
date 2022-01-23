class Solution(object):
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """
        '''
        for row in board:
            test = {}
            for n in row:
                if n == '.':
                    continue
                test[n] = test.get(n, 0) + 1
                if test[n] > 1:
                    return False
        for i in xrange(9):
            test = {}
            for j in xrange(9):
                n = board[j][i]
                if n == '.':
                    continue
                test[n] = test.get(n, 0) + 1
                if test[n] > 1:
                    return False
        for ii in xrange(0, 9, 3):
            for jj in xrange(0, 9, 3):
                test = {}
                for i in xrange(ii, ii+3):
                    for j in xrange(jj, jj+3):
                        n = board[i][j]
                        if n == '.':
                            continue
                        test[n] = test.get(n, 0) + 1
                        if test[n] > 1:
                            return False
        '''
        test = set()
        for i in xrange(9):
            for j in xrange(9):
                cur = board[i][j]
                if cur != '.':
                    if (i, cur) in test or (cur, j) in test or (i/3, j/3, cur) in test:
                        return False
                    test.add((i, cur))
                    test.add((cur, j))
                    test.add((i/3, j/3, cur))
        return True
