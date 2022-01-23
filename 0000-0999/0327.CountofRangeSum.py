class Solution(object):
    def countRangeSum(self, nums, lower, upper):
        """
        :type nums: List[int]
        :type lower: int
        :type upper: int
        :rtype: int
        """
        sums = [0] * (len(nums)+1)
        for i in xrange(len(nums)):
            sums[i+1] = sums[i] + nums[i]
        return self.merge(sums, 0, len(sums)-1, lower, upper)

    def merge(self, sums, start, end, lower, upper):
        if start >= end:
            return 0
        mid = (start+end)/2
        left = self.merge(sums, start, mid, lower, upper)
        right = self.merge(sums, mid+1, end, lower, upper)
        count = left + right
        i = j = mid+1
        for index in xrange(start, mid+1):
            while i <= end and sums[i]-sums[index] < lower:
                i += 1
            while j <= end and sums[j]-sums[index] <= upper:
                j += 1
            count += j - i
        sums[start:end+1] = sorted(sums[start:end+1])
        return count

test = Solution()
print test.countRangeSum([-2, 5, -1], -2, 2)

