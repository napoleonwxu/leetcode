from heapq import heappush, heappop

class Solution(object):
    def findKthLargest(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        #return sorted(nums)[::-1][k-1]     # Quick Sort: O(NlgN)
        #return self.kthLargest(nums, k, 0, len(nums)-1)    # Divide and Conquer: O(NlgK)
        # Min Heap: O(NlgK)
        heap = []
        for num in nums:
            heappush(heap, num)
            while len(heap) > k:
                heappop(heap)
        return heap[0]

    def kthLargest(self, nums, k, left, right):
        mid = self.partition(nums, left, right)
        if mid == k - 1:
            return nums[mid]
        elif mid > k - 1:
            return self.kthLargest(nums, k, left, mid - 1)
        else:
            return self.kthLargest(nums, k, mid + 1, right)

    def partition(self, nums, left, right):
        pivot = nums[right]
        mid = left - 1
        for i in xrange(left, right):
            if nums[i] > pivot: # Kth Smallest: nums[i] < pivot
                mid += 1
                if i != mid:
                    nums[i], nums[mid] = nums[mid], nums[i]
        mid += 1
        nums[mid], nums[right] = nums[right], nums[mid]
        return mid
