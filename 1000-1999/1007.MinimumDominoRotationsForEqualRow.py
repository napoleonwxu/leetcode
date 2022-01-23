class Solution(object):
    def minDominoRotations(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        Set = set(range(1, 7))
        for p in zip(A, B):
            Set &= set(p)
            if not Set:
                return -1
        n = Set.pop()
        return min(len(A)-A.count(n), len(B)-B.count(n))