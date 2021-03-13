class Solution {
    public int countGoodRectangles(int[][] rectangles) {
        int cnt = 0, maxLen = 0;
        for (int[] rectangle : rectangles) {
            int len = Math.min(rectangle[0], rectangle[1]);
            if (len > maxLen) {
                maxLen = len;
                cnt = 1;
            } else if (len == maxLen) {
                cnt++;
            }
        }
        return cnt;
    }
}