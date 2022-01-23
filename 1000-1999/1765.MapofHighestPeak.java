class Solution {
    public int[][] highestPeak(int[][] isWater) {
        int m = isWater.length, n = isWater[0].length;
        int[][] ans = new int[m][n];
        List<int[]> queue = new ArrayList<>();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (isWater[i][j] == 1) {
                    queue.add(new int[]{i, j});
                }
                ans[i][j] = -1;
            }
        }
        int depth = 0;
        while (queue.size() > 0) {
            List<int[]> next = new ArrayList<>();
            for (int[] loc : queue) {
                int i = loc[0], j = loc[1];
                if (i < 0 || i >= m || j < 0 || j >= n || ans[i][j] != -1) {
                    continue;
                }
                ans[i][j] = depth;
                next.add(new int[]{i-1, j});
                next.add(new int[]{i+1, j});
                next.add(new int[]{i, j-1});
                next.add(new int[]{i, j+1});
            }
            depth++;
            queue = next;
        }
        return ans;
    }
}