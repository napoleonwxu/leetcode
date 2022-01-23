class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        #return bin(int(a, 2) + int(b, 2))[2:]
        c = 0
        ans = ''
        i, j = len(a)-1, len(b)-1
        while i >= 0 or j >= 0 or c:
            if i >= 0:
                c += int(a[i])
                i -= 1
            if j >= 0:
                c += int(b[j])
                j -= 1
            ans = str(c%2) + ans
            c >>= 1
        return ans
        