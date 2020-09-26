class Solution {
    public int longestSubarray(int[] nums) {
        int n = nums.length;
        int[] left = new int[n], right = new int[n];
        int l = 0, r = 0;
        for (int i = 0; i < n; i++) {
            left[i] = l;
            right[n-1-i] = r;
            if (nums[i] == 1) {
                l++;
            } else {
                l = 0;
            }
            if (nums[n-1-i] == 1) {
                r++;
            } else {
                r = 0;
            }
        }
        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, left[i]+right[i]);
        }
        return ans;
    }
}