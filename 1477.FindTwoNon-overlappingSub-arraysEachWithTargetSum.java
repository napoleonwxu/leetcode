class Solution {
    public int minSumOfLengths(int[] arr, int target) {
        HashMap<Integer, Integer> prefix = new HashMap<>();
        prefix.put(0, -1);
        int sum = 0, minSize = Integer.MAX_VALUE, ans = Integer.MAX_VALUE;
        for (int i = 0; i < arr.length; i++) {
            sum += arr[i];
            prefix.put(sum, i);
        }
        sum = 0;
        for (int i = 0; i < arr.length; i++) {
            sum += arr[i];
            if (prefix.get(sum-target) != null) {
                minSize = Math.min(minSize, i-prefix.get(sum-target));
            }
            if (minSize < Integer.MAX_VALUE && prefix.get(sum+target) != null) {
                ans = Math.min(ans, prefix.get(sum+target)-i+minSize);
            }
        }
        return ans < Integer.MAX_VALUE ? ans : -1;
    }
}