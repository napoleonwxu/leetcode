class Solution {
    public int[] getSumAbsoluteDifferences(int[] nums) {
        int n = nums.length;
        int[] ans = new int[n];
        for (int i = 1; i < n; i++) {
            ans[0] += nums[i] - nums[0];
        }
        for (int i = 1; i < n; i++) {
            int d = nums[i] - nums[i-1];
            ans[i] = ans[i-1] + i*d - (n-i)*d;
        }
        return ans;
    }
}