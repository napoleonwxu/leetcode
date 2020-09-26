class Solution {
    public int minOperations(int n) {
        // O(1)
        return n*n/4;
        /* O(n)
        int ans = 0;
        for (int i = 0; i < n/2; i++) {
            ans += n - (2*i+1);
        }
        return ans;
        */
    }
}