class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        t = ab = reduce(operator.xor, nums)
        t &= -t
        a = reduce(operator.xor, filter(lambda x: x&t, nums))
        return [a, a^ab]
        '''
        ab = reduce(operator.xor, nums)
        tmp = 1
        while tmp&ab == 0:
            tmp <<= 1
        nums1 = []
        for num in nums:
            if num & tmp:
                nums1.append(num)
        a = reduce(operator.xor, nums1)
        return [a, a^ab]
        '''