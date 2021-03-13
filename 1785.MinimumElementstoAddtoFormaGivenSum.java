class Solution {
    public int minElements(int[] nums, int limit, int goal) {
        long sum = (long)goal;
        for (int num : nums) {
            sum -= num;
        }
        sum = Math.abs(sum);
        return (int)((sum + limit - 1)/limit);
        //return (int)(sum/limit) + (sum%limit > 0 ? 1 : 0);
    }
}