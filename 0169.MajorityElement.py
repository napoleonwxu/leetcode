class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        """
        majLen = len(nums)/2
        d = {}
        for n in nums:
            '''
            if not d.get(n):
                d[n] = 1
            else:
                d[n] += 1
            '''
            d[n] = d.get(n, 0) + 1
            if d[n] > majLen:
                return n
        """ 
        '''
        nums.sort()
        return nums[len(nums)/2]
        '''
        cad, count = 0, 0
        for n in nums:
            if n == cad:
                count += 1
            elif count == 0:
                cad, count = n, 1
            else:
                count -= 1
        return cad
                
        