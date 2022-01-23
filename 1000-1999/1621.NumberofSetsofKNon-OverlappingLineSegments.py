class Solution:
    def numberOfSets(self, n: int, k: int) -> int:
        MOD = 10**9 + 7
        cnt = 1
        for i in range(1, 2*k+1):
            cnt = cnt * (n+k-i) // i
        return cnt % MOD
        