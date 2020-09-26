class Solution {
    private char[][] grid;
    private boolean[][] visited;
    private int m, n;
    private int[] dir = {0, -1, 0, 1, 0};
    
    public boolean containsCycle(char[][] grid) {
        this.grid = grid;
        m = grid.length;
        n = grid[0].length;
        visited = new boolean[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (!visited[i][j] && dfs(i, j, -1, -1, grid[i][j])) {
                    return true;
                }
            }
        }
        return false;
    }
    
    private boolean dfs(int i, int j, int pre_i, int pre_j, char ch) {
        if (visited[i][j]) {
            return true;
        }
        visited[i][j] = true;
        boolean flag = false;
        for (int k = 0; k < 4; k++) {
            int nxt_i = i + dir[k];
            int nxt_j = j + dir[k+1];
            if (nxt_i >= 0 && nxt_i < m && nxt_j >= 0 && nxt_j < n && (nxt_i != pre_i || nxt_j != pre_j) && grid[nxt_i][nxt_j] == ch) {
                flag |= dfs(nxt_i, nxt_j, i, j, ch);
            }
        }
        return flag;
    }
}