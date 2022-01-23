class Solution(object):
    def rotate(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: void Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        for i in xrange(n/2):
            p = n-1-i
            for j in xrange(i, p):
                q = n-1-j
                temp = matrix[i][j]
                matrix[i][j] = matrix[q][i]
                matrix[q][i] = matrix[p][q]
                matrix[p][q] = matrix[j][p]
                matrix[j][p] = temp
        """
        for i in xrange(n):
            for j in xrange(i):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
                '''
                temp = matrix[i][j]
                matrix[i][j] = matrix[j][i]
                matrix[j][i] = temp
                '''
        for j in xrange(n/2):
            for i in xrange(n):
                matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
                '''
                temp = matrix[i][j]
                matrix[i][j] = matrix[i][n-1-j]
                matrix[i][n-1-j] = temp
                '''
        """