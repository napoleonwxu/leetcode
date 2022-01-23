class Solution {
    public int maxAbsoluteSum(int[] nums) {
        int maxSum = 0, minSum = 0;
        for (int num : nums) {
            maxSum = Math.max(0, maxSum+num);
            minSum = Math.min(0, minSum+num);
        }
        return maxSum - minSum;
    }
}