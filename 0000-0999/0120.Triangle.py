class Solution(object):
    def minimumTotal(self, triangle):
        """
        :type triangle: List[List[int]]
        :rtype: int
        """
        if not triangle:
            return 0
        '''
        for row in xrange(len(triangle)-2, -1, -1):
            for col in xrange(len(triangle[row])):
                triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
        return triangle[0][0]
        '''
        temp = [0] * (len(triangle)+1)
        for row in xrange(len(triangle)-1, -1, -1):
            for col in xrange(len(triangle[row])):
                temp[col] = triangle[row][col] + min(temp[col], temp[col+1])
        return temp[0]
        