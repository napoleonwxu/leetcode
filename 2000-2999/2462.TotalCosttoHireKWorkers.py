class Solution:
    def totalCost(self, costs: List[int], k: int, candidates: int) -> int:
        i, j = candidates, max(candidates, len(costs)-candidates)
        p, q = costs[:i], costs[j:]
        heapify(p)
        heapify(q)
        ans = 0
        for _ in range(k):
            if not q or (p and p[0] <= q[0]):
                ans += heappop(p)
                if i < j:
                    heappush(p, costs[i])
                    i += 1
            else:
                ans += heappop(q)
                if i < j:
                    j -= 1
                    heappush(q, costs[j])
        return ans
        