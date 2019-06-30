class Solution(object):
    def maxEqualRowsAfterFlips(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: int
        """
        return max(collections.Counter(tuple(c^r[0] for c in r) for r in matrix).values())