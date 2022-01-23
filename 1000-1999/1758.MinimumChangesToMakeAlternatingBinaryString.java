class Solution {
    public int minOperations(String s) {
        int n = s.length(), cnt = 0;
        for (int i = 0; i < n; i++) {
            if (s.charAt(i)-'0' != i%2) {
                cnt++;
            }
        }
        return Math.min(cnt, n-cnt);
    }
}