class Solution(object):
    def solveSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """
        self.board = board
        self.solve()

    def solve(self):
        row, col = self.findLocation()
        if row == -1 and col == -1:
            return True
        for n in xrange(1, 10):
            ch = str(n)
            if self.isVaild(row, col, ch):
                self.board[row][col] = ch
                if self.solve():
                    return True
                self.board[row][col] = '.'
        return False

    def findLocation(self):
        for row in xrange(9):
            for col in xrange(9):
                if self.board[row][col] == '.':
                    return row, col
        return -1, -1

    def isVaild(self, row, col, ch):
        if self.isVaildRow(row, ch) and self.isVaildCol(col, ch) and self.isVaildSqure(row - row%3, col - col%3, ch):
            return True
        return False

    def isVaildRow(self, row, ch):
        for s in self.board[row]:
            if s == ch:
                return False
        return True

    def isVaildCol(self, col, ch):
        for row in xrange(9):
            if self.board[row][col] == ch:
                return False
        return True

    def isVaildSqure(self, rowSq, colSq, ch):
        for row in xrange(rowSq, rowSq+3):
            for col in xrange(colSq, colSq+3):
                if self.board[row][col] == ch:
                    return False
        return True

a = [['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.'],
     ['.', '.', '.', '.', '.', '.', '.', '.', '.']]

b = [['5', '3', '.', '.', '7', '.', '.', '.', '.'],
     ['6', '.', '.', '1', '9', '5', '.', '.', '.'],
     ['.', '9', '8', '.', '.', '.', '.', '6', '.'],
     ['8', '.', '.', '.', '6', '.', '.', '.', '3'],
     ['4', '.', '.', '8', '.', '3', '.', '.', '1'],
     ['7', '.', '.', '.', '2', '.', '.', '.', '6'],
     ['.', '6', '.', '.', '.', '.', '2', '8', '.'],
     ['.', '.', '.', '4', '1', '9', '.', '.', '5'],
     ['.', '.', '.', '.', '8', '.', '.', '7', '9']]
test = Solution()
test.solveSudoku(a)
for i in xrange(9):
    print test.board[i]