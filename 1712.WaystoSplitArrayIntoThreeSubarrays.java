class Solution {
    private static final int MOD = (int)1e9 + 7;
    
    public int waysToSplit(int[] nums) {
        int n = nums.length;
        int[] sums = new int[n];
        for (int i = 0; i < n; i++) {
            sums[i] = (i == 0 ? 0 : sums[i-1]) + nums[i];
        }
        int cnt = 0;
        for (int i = 1; i < n-1; i++) {
            int left = binarySearch(sums, sums[i-1], i, true);
            int right = binarySearch(sums, sums[i-1], i, false);
            if (left != -1 && right != -1) {
                cnt = (cnt + right - left + 1) % MOD;
            }
        }
        return cnt;
    }
    
    private int binarySearch(int[] sums, int leftSum, int idx, boolean searchLeft) {
        int n = sums.length, left = idx, right = n-2, ans = -1;
        while (left <= right) {
            int mid = (left + right) / 2;
            int midSum = sums[mid] - sums[idx-1], rightSum = sums[n-1] - sums[mid];
            if (leftSum <= midSum && midSum <= rightSum) {
                ans = mid;
                if (searchLeft) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else if (leftSum > midSum) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return ans;
    }
}