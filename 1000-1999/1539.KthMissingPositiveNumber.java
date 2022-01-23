class Solution {
    public int findKthPositive(int[] arr, int k) {
        int num = 1, i = 0;
        while (k > 0) {
            if (i < arr.length && num == arr[i]) {
                i++;
            } else {
                k--;
            }
            num++;
        }
        return num-1;
    }
}