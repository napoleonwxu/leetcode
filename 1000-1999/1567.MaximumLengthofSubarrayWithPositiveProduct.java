class Solution {
    public int getMaxLen(int[] nums) {
        int ans = 0, cnt_neg = 0, first_neg = -1, idx_zero = -1;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] < 0) {
                cnt_neg++;
                if (first_neg == -1) {
                    first_neg = i;
                }
            }
            if (nums[i] != 0) {
                if (cnt_neg%2 == 0) {
                    ans = Math.max(ans, i-idx_zero);
                } else {
                    ans = Math.max(ans, i-first_neg);
                }
            } else {
                cnt_neg = 0;
                idx_zero = i;
                first_neg = -1;
            }
        }
        return ans;
    }
}