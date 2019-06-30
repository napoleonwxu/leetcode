class Solution(object):
    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        ans = []
        while matrix:
            ans += matrix.pop(0)
            if matrix and matrix[0]:
                for row in matrix:
                    ans.append(row.pop())
            if matrix:
                ans += matrix.pop()[::-1]
            if matrix and matrix[0]:
                for row in matrix[::-1]:
                    ans.append(row.pop(0))
        return ans

        """
        if not matrix:
            return []
        ans = []
        t, b = 0, len(matrix)-1
        l, r = 0, len(matrix[0])-1
        while t < b and l < r:
            ans.extend([matrix[t][j] for j in xrange(l, r)])
            ans.extend([matrix[i][r] for i in xrange(t, b)])
            ans.extend([matrix[b][j] for j in xrange(r, l, -1)])
            ans.extend([matrix[i][l] for i in xrange(b, t, -1)])
            '''
            for j in xrange(l, r):
                ans.append(matrix[t][j])
            for i in xrange(t, b):
                ans.append(matrix[i][r])
            for j in xrange(r, l, -1):
                ans.append(matrix[b][j])
            for i in xrange(b, t, -1):
                ans.append(matrix[i][l])
            '''
            t += 1
            b -= 1
            l += 1
            r -= 1
        if t == b:
            ans.extend([matrix[t][j] for j in xrange(l, r+1)])
            '''
            for j in xrange(l, r+1):
                ans.append(matrix[t][j])
            '''
            #return ans
        elif l == r:
            ans.extend([matrix[i][l] for i in xrange(t, b+1)])
            '''
            for i in xrange(t, b+1):
                ans.append(matrix[i][l])
            '''
        return ans
        """