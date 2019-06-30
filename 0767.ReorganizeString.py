class Solution(object):
    def reorganizeString(self, S):
        """
        :type S: str
        :rtype: str
        """
        heap = [(-S.count(ch), ch) for ch in set(S)]
        if any(-cnt > (len(S)+1)>>1 for cnt, ch in heap):
            return ''
        heapq.heapify(heap)
        ans = ''
        while len(heap) >= 2:
            cnt1, ch1 = heapq.heappop(heap)
            cnt2, ch2 = heapq.heappop(heap)
            ans = ''.join([ans, ch1, ch2])
            if cnt1+1:
                heapq.heappush(heap, (cnt1+1, ch1))
            if cnt2+1:
                heapq.heappush(heap, (cnt2+1, ch2))
        if heap:
            ans = ''.join([ans, heap[0][1]])
        return ans
        