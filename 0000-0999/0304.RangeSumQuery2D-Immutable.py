class NumMatrix(object):
    def __init__(self, matrix):
        """
        initialize your data structure here.
        :type matrix: List[List[int]]
        """
        if not matrix:
            return
        self.m = matrix
        for j in xrange(1, len(matrix[0])):
            self.m[0][j] += self.m[0][j-1]
        for i in xrange(1, len(matrix)):
            self.m[i][0] += self.m[i-1][0]
            for j in xrange(1, len(matrix[i])):
                self.m[i][j] += self.m[i][j-1] + self.m[i-1][j] - self.m[i-1][j-1]

    def sumRegion(self, row1, col1, row2, col2):
        """
        sum of elements matrix[(row1,col1)..(row2,col2)], inclusive.
        :type row1: int
        :type col1: int
        :type row2: int
        :type col2: int
        :rtype: int
        """
        if row1 == 0 and col1 == 0:
            return self.m[row2][col2]
        elif row1 == 0:
            return self.m[row2][col2] - self.m[row2][col1-1]
        elif col1 == 0:
            return self.m[row2][col2] - self.m[row1-1][col2]
        else:
            return self.m[row2][col2] - self.m[row1-1][col2] - self.m[row2][col1-1] + self.m[row1-1][col1-1]


# Your NumMatrix object will be instantiated and called as such:
# numMatrix = NumMatrix(matrix)
# numMatrix.sumRegion(0, 1, 2, 3)
# numMatrix.sumRegion(1, 2, 3, 4)