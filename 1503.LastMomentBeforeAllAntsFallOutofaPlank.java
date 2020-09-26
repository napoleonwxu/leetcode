class Solution {
    public int getLastMoment(int n, int[] left, int[] right) {
        int ans = 0;
        for (int start : left) {
            ans = Math.max(ans, start);
        }
        for (int start : right) {
            ans = Math.max(ans, n-start);
        }
        return ans;
    }
}