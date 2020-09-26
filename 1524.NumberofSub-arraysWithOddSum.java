class Solution {
    private static final int MOD = (int)1e9+7;
    
    public int numOfSubarrays(int[] arr) {
        // O(n) + O(1)
        int cnt = 0, odd = 0, even = 0;
        for (int a : arr) {
            if (a%2 == 1) {
                int tmp = odd;
                odd = even + 1;
                even = tmp;
            } else {
                even++;
            }
            cnt = (cnt+odd) % MOD;
        }
        return cnt;
    }
}