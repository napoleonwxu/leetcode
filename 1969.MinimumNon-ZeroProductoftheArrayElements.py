class Solution:
    def minNonZeroProduct(self, p: int) -> int:
        MOD = 10**9 + 7
        return pow(2**p-2, 2**(p-1)-1, MOD) * (2**p-1) % MOD
        #return pow((1<<p)-2, (1<<(p-1))-1, MOD) * ((1<<p)-1) % MOD
    