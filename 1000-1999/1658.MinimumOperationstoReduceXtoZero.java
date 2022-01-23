class Solution {
    public int minOperations(int[] nums, int x) {
        int n = nums.length, target = -x;
        for (int num : nums) {
            target += num;
        }
        if (target < 0) {
            return -1;
        }
        if (target == 0) {
            return n;
        }
        Map<Integer, Integer> map = new HashMap<>();
        map.put(0, -1);
        int sum = 0, len = Integer.MIN_VALUE;
        for (int i = 0; i < n; i++) {
            sum += nums[i];
            if (map.containsKey(sum-target)) {
                len = Math.max(len, i - map.get(sum-target));
            }
            map.put(sum, i);
        }
        return len==Integer.MIN_VALUE ? -1 : n-len;
    }
}