class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        heapq.heapify(nums)
        for i in range(k):
            heapq.heappush(nums, heapq.heappop(nums)+1)
        prod, MOD = 1, 10**9+7
        for num in nums:
            prod = prod * num % MOD
        return prod