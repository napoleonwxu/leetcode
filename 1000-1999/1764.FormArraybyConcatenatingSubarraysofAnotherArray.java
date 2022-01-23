class Solution {
    public boolean canChoose(int[][] groups, int[] nums) {
        int n = groups.length, i = 0, j = 0;
        while (i < n && j < nums.length) {
            if (isSame(groups[i], nums, j)) {
                j += groups[i].length;
                i++;
            } else {
                j++;
            }
        }
        return i >= n ? true : false;
    }
    
    private boolean isSame(int[] group, int[] nums, int j) {
        for (int i = 0; i < group.length; i++) {
            if (j+i >= nums.length || group[i] != nums[j+i]) {
                return false;
            }
        }
        return true;
    }
}