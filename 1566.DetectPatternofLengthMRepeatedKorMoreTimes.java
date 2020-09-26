class Solution {
    public boolean containsPattern(int[] arr, int m, int k) {
        int cnt = 0;
        for (int i = 0; i < arr.length-m; i++) {
            if (arr[i] != arr[i+m]) {
                cnt = 0;
            } else if (++cnt >= (k-1)*m) {
                return true;
            }
        }
        return false;
    }
}