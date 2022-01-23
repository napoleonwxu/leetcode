class Solution {
    private static final int MOD = (int)1e9+7;
    
    public int rangeSum(int[] nums, int n, int left, int right) {
        // O(N^2 * lgN) + O(N^2)
        ArrayList<Integer> arr = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            int sum = 0;
            for (int j = i; j < n; j++) {
                sum += nums[j];
                arr.add(sum);
            }
        }
        Collections.sort(arr);
        int ans = 0;
        for (int i = left-1; i < right; i++) {
            ans = (ans + arr.get(i)) % MOD;
        }
        return ans;
    }
}