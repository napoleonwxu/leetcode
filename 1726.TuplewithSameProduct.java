class Solution {
    public int tupleSameProduct(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            for (int j = i+1; j < nums.length; j++) {
                int prod = nums[i] * nums[j];
                int cnt = map.getOrDefault(prod, 0);
                ans += 8 * cnt;
                map.put(prod, cnt+1);
            }
        }
        return ans;
    }
}