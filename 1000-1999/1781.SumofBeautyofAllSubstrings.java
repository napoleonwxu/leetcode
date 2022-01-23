class Solution {
    public int beautySum(String s) {
        int sum = 0;
        for (int i = 0; i < s.length(); i++) {
            int[] cnts = new int[26];
            for (int j = i; j < s.length(); j++) {
                cnts[s.charAt(j)-'a']++;
                sum += getBeauty(cnts);
            }
        }
        return sum;
    }
    
    private int getBeauty(int[] cnts) {
        int maxValue = 0, minValue = Integer.MAX_VALUE;
        for (int cnt : cnts) {
            maxValue = Math.max(maxValue, cnt);
            if (cnt > 0) {
                minValue = Math.min(minValue, cnt);
            }
        }
        return maxValue - minValue;
    }
}