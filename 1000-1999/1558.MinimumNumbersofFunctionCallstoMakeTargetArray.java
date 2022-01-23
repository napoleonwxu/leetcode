class Solution {
    private int[] nums;
    private int step = 0;
    
    public int minOperations(int[] nums) {
        // O(Nlg(10^9))
        int ones = 0, maxLen = 1;
        for (int num : nums) {
            int len = 0;
            while (num > 0) {
                ones += num & 1;
                num >>= 1;
                len++;
            }
            maxLen = Math.max(maxLen, len);
        }
        return ones + maxLen - 1;
        /*
        this.nums = nums;
        while (!isAllZero()) {
            if (isAllEven()) {
                divideTwo();
            } else {
                toEven();
            }
        }
        return step;
        */
    }
    
    private boolean isAllZero() {
        for (int num : nums) {
            if (num != 0) {
                return false;
            }
        }
        return true;
    }
    
    private boolean isAllEven() {
        for (int num : nums) {
            if (num%2 == 1) {
                return false;
            }
        }
        return true;
    }
    
    private void divideTwo() {
        for (int i = 0; i < nums.length; i++) {
            nums[i] /= 2;
        }
        step++;
    }
    
    private void toEven() {
        for (int i = 0; i < nums.length; i++) {
            if (nums[i]%2 == 1) {
                nums[i]--;
                step++;
            }
        }
    }
}