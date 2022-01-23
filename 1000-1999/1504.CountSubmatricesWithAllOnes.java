class Solution {
    public int numSubmat(int[][] mat) {
        // O(mmn)
        int m = mat.length, n = mat[0].length;
        int ans = 0;
        for (int i = 0; i < m; i++) {
            int[] arr = new int[n];
            Arrays.fill(arr, 1);
            for (int j = i; j < m; j++) {
                for (int k = 0; k < n; k++) {
                    arr[k] &= mat[j][k];
                }
                ans += numSubArr(arr);
            }
        }
        return ans;
    }
    
    private int numSubArr(int[] arr) {
        int ans = 0, cnt = 0;
        for (int a : arr) {
            cnt = a==0 ? 0 : cnt+1;
            ans += cnt;
        }
        return ans;
    }
}