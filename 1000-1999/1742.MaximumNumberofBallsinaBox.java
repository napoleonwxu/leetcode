class Solution {
    public int countBalls(int lowLimit, int highLimit) {
        int[] cnt = new int[5*9+1];
        int ans = 0;
        for (int ball = lowLimit; ball <= highLimit; ball++) {
            int box = 0;
            for (int tmp = ball; tmp > 0; tmp /= 10) {
                box += tmp % 10;
            }
            ans = Math.max(ans, ++cnt[box]);
        }
        return ans;
    }
}