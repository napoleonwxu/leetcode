from heapq import heappush, heappop
class MedianFinder:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.min_heap = []
        self.max_heap = []

    def addNum(self, num):
        """
        Adds a num into the data structure.
        :type num: int
        :rtype: void
        """
        if not self.min_heap:
            heappush(self.min_heap, num)
        elif num >= self.min_heap[0]:
            heappush(self.min_heap, num)
            if len(self.min_heap) - len(self.max_heap) > 1:
                heappush(self.max_heap, -heappop(self.min_heap))
        else:
            heappush(self.max_heap, -num)
            if len(self.max_heap) > len(self.min_heap):
                heappush(self.min_heap, -heappop(self.max_heap))

    def findMedian(self):
        """
        Returns the median of current data stream
        :rtype: float
        """
        if (len(self.min_heap)+len(self.max_heap))%2:
            return self.min_heap[0]
        else:
            return (self.min_heap[0]-self.max_heap[0])/2.0

# Your MedianFinder object will be instantiated and called as such:
# mf = MedianFinder()
# mf.addNum(1)
# mf.findMedian()