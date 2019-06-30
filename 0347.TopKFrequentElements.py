class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        #return [num for num, freq in collections.Counter(nums).most_common(k)]
        #return zip(*collections.Counter(nums).most_common(k))[0]
        # bucket sort
        dic = {}
        for num in nums:
            dic[num] = dic.get(num, 0) + 1
        bucket = [[] for i in xrange(len(nums)+1)]
        for num, freq in dic.items():
            bucket[freq].append(num)
        ans = []
        i = len(nums)
        while k > 0:
            ans += bucket[i]
            k -= len(bucket[i])
            i -= 1
        return ans
