class Solution(object):
    def countSmaller(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        ans = [0] * len(nums)
        self.mergeSort(list(enumerate(nums)), ans)
        return ans

    def mergeSort(self, enum, ans):
        if len(enum) <= 1:
            return enum
        half = len(enum) / 2
        left = self.mergeSort(enum[:half], ans)
        right = self.mergeSort(enum[half:], ans)
        #for i in range(len(enum))[::-1]:
        for i in xrange(len(enum)-1, -1, -1):
            if not right or left and left[-1][1] > right[-1][1]:
                ans[left[-1][0]] += len(right)
                enum[i] = left.pop()
            else:
                enum[i] = right.pop()
        return enum
