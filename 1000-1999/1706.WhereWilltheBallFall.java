class Solution {
    public int[] findBall(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[] ans = new int[n];
        for (int c = 0; c < n; c++) {
            int i = 0, j = c;
            for (; i < m; i++) {
                if (grid[i][j] == 1) {
                    if (j < n-1 && grid[i][j+1] == 1) {
                        j++;
                    } else {
                        break;
                    }
                } else {
                    if (j > 0 && grid[i][j-1] == -1) {
                        j--;
                    } else {
                        break;
                    }
                }
            }
            ans[c] = i >= m ? j : -1;
        }
        return ans;
    }
}