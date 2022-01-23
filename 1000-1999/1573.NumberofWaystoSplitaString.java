class Solution {
    private static final int MOD = (int)1e9+7;
    
    public int numWays(String s) {
        int n = s.length(), ones = 0;
        for (int i = 0; i < n; i++) {
            ones += s.charAt(i) - '0';
        }
        if (ones == 0) {
            return (int)((n-1L) * (n-2L) / 2 % MOD);
        }
        if (ones%3 != 0) {
            return 0;
        }
        ones /= 3;
        long zeros1 = 0, zeros2 = 0;
        for (int i = 0, cnt = 0; i < n; i++) {
            cnt += s.charAt(i) - '0';
            if (cnt == ones) {
                zeros1++;
            } else if (cnt == ones*2) {
                zeros2++;
            }
        }
        return (int)(zeros1 * zeros2 % MOD);
    }
}