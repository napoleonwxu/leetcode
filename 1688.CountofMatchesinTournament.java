class Solution {
    public int numberOfMatches(int n) {
        //return n - 1;
        int cnt = 0;
        while (n > 1) {
            cnt += n / 2;
            n = (n+1) / 2;
        }
        return cnt;
    }
}