class Solution {
    private static final int MOD = (int)1e9 + 7;
    
    public int maxSumRangeQuery(int[] nums, int[][] requests) {
        int n = nums.length;
        int[] cnts = new int[n];
        for (int[] req : requests) {
            cnts[req[0]]++;
            if (req[1]+1 < n) {
                cnts[req[1]+1]--;   
            }
        }
        for (int i = 1; i < n; i++) {
            cnts[i] += cnts[i-1];
        }
        Arrays.sort(cnts);
        Arrays.sort(nums);
        long sum = 0;
        for (int i = 0; i < n; i++) {
            sum = (sum + (1L*cnts[i]*nums[i]) % MOD) % MOD;
        }
        return (int)sum;
    }
}