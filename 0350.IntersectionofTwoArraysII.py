class Solution(object):
    def intersect(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        # O(m+n)
        dic = {}
        ans = []
        for num in nums1:
            dic[num] = dic.get(num, 0) + 1
        for num in nums2:
            if dic.get(num, 0) > 0:
                ans.append(num)
                dic[num] -= 1
        return ans
        ''' # O(max(m,n)logmax(m,n))
        nums1.sort()
        nums2.sort()
        ans = []
        i = j = 0
        while i < len(nums1) and j < len(nums2):
            if nums1[i] < nums2[j]:
                i += 1
            elif nums1[i] > nums2[j]:
                j += 1
            else:
                ans.append(nums1[i])
                i += 1
                j += 1
        return ans
        '''
        ''' Slow
        ans = []
        set1, set2 = set(nums1), set(nums2)
        for num in set1:
            if num in set2:
                ans += [num] * min(nums1.count(num), nums2.count(num))
        return ans
        '''
