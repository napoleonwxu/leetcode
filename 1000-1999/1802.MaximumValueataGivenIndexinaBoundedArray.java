class Solution {
    public int maxValue(int n, int index, int maxSum) {
        maxSum -= n;
        int left = 0, right = maxSum;
        while (left < right) {
            int mid = (left + right + 1) / 2;
            if (check(n, index, mid) <= maxSum) {
                left = mid;
            } else {
                right = mid - 1;
            }
        }
        return left + 1;
    }
    
    private long check(int n, int index, int max) {
        int min = Math.max(max-index, 0);
        long sum = (long)(min+max) * (max-min+1) / 2;
        min = Math.max(max-(n-1-index), 0);
        sum += (long)(min+max) * (max-min+1) / 2;
        return sum - max;
    }
}