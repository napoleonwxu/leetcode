class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        '''
        for i in xrange(len(matrix)):
            for j in xrange(len(matrix[0])):
                if matrix[i][j] == target:
                    return True
                if matrix[i][j] > target:
                    break
        return False
        '''
        m = len(matrix)
        i, j = 0, len(matrix[0])-1
        while i < m and j >= 0:
            if target == matrix[i][j]:
                return True
            if target < matrix[i][j]:
                j -= 1
            else:
                i += 1
        return False
        