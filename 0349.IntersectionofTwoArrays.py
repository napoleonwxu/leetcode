class Solution(object):
    def intersection(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        #return list(set(nums1) & set(nums2))
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
                if not ans or nums1[i] != ans[-1]:
                    ans.append(nums1[i])
                i += 1
                j += 1
        return ans