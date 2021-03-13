class Solution {
    public int[] decrypt(int[] code, int k) {
        int n = code.length;
        int[] ans = new int[n];
        if (k == 0) {
            return ans;
        }
        // O(n)
        int start = 1, end = k, sum = 0;
        if (k < 0) {
            start = n + k;
            end = n - 1;
        }
        for (int i = start; i <= end; i++) {
            sum += code[i];
        }
        for (int i = 0; i < n; i++) {
            ans[i] = sum;
            sum += code[++end%n];
            sum -= code[start++%n];
        }
        /* O(kn)
        for (int i = 0; i < n; i++) {
            if (k > 0) {
                for (int j = 1; j <= k; j++) {
                    ans[i] += code[(i+j)%n];
                }
            } else {
                for (int j = k; j <= -1; j++) {
                    ans[i] += code[(i+j+n)%n];
                }
            }
        }*/
        return ans;
    }
}