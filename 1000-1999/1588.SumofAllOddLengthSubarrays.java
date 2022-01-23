class Solution {
    public int sumOddLengthSubarrays(int[] arr) {
        int n = arr.length;
        int ans = 0;
        //O(n)
        for (int i = 0; i < n; i++) {
            ans += ((i+1) * (n-i) + 1)/2 * arr[i];
        }
        /* O(n^2)
        int[] sum = new int[n+1];
        for (int i = 0; i < n; i++) {
            sum[i+1] = sum[i] + arr[i];
        }
        for (int i = 0; i < n; i++) {
            for (int j = i+1; j <= n; j += 2) {
                ans += sum[j] - sum[i];
            }
        }
        */
        return ans;
    }
}