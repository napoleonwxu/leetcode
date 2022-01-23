class Solution {
    public double trimMean(int[] arr) {
        Arrays.sort(arr);
        int n = arr.length, sum = 0;
        for (int i = n/20; i < n - n/20; i++) {
            sum += arr[i];
        }
        return sum / (n - n/10d);
    }
}