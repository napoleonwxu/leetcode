class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1 or len(s) <= numRows:
            return s
        rows = [''] * numRows
        period = 2*(numRows-1)
        for i, ch in enumerate(s):
            t = i%period
            if t < numRows:
                rows[t] += ch
            else:
                rows[period-t] += ch
        return ''.join(rows)
