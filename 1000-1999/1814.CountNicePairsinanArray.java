class Solution {
    private static final int MOD = (int)1e9 + 7;
    
    public int countNicePairs(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        int cnt = 0;
        for (int num : nums) {
            int sub = num - rev(num);
            int pre = map.getOrDefault(sub, 0);
            cnt = (cnt + pre) % MOD;
            map.put(sub, pre+1);
        }
        return cnt;
    }
    
    private int rev(int num) {
        int rev = 0;
        while (num > 0) {
            rev = 10*rev + num%10;
            num /= 10;
        }
        return rev;
    }
}