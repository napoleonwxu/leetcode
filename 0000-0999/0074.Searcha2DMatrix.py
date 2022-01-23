class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix:
            return False
        m, n = len(matrix), len(matrix[0])
        left, right = 0, m*n-1
        while left <= right:
            mid = left + (right-left)/2
            temp = matrix[mid/n][mid%n]
            if target == temp:
                return True
            elif target < temp:
                right = mid - 1
            else:
                left = mid + 1
        return False
        '''
        m = len(matrix)
        i, j = 0, len(matrix[0]) - 1
        while i < m and j >= 0:
            if target == matrix[i][j]:
                return True
            elif target > matrix[i][j]:
                i += 1
            else:
                j -= 1
        return False
        '''