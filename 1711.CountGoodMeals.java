class Solution {
    private static final int MOD = (int)1e9 + 7;
    
    public int countPairs(int[] deliciousness) {
        // O(32n)
        Map<Integer, Integer> map = new HashMap<>();
        long cnt = 0;
        for (int d : deliciousness) {
            int power = 1;
            for (int i = 0; i < 32; i++) {
                if (map.containsKey(power - d)) {
                    cnt = (cnt + map.get(power - d)) % MOD;
                }
                power <<= 1;
            }
            map.put(d, map.getOrDefault(d, 0) + 1);
        }
        return (int)cnt;
    }
}