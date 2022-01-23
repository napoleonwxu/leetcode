class Solution {
    public int getWinner(int[] arr, int k) {
        int cnt = 0, ans = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > ans) {
                ans = arr[i];
                cnt = 0;
            }
            if (++cnt >= k) {
                return ans;
            }
        }
        return ans;
    }
}