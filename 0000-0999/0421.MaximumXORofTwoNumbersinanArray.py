class Solution(object):
    def findMaximumXOR(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        ans = 0
        mask = 0
        for i in xrange(31, -1, -1):
            mask |= (1 << i)
            Set = {num & mask for num in nums}
            '''
            Set = set()
            for num in nums:
                Set.add(num & mask)
            '''
            tmp = ans | (1 << i)
            for prefix in Set:
                if tmp ^ prefix in Set:
                    ans = tmp
                    break
        return ans
        