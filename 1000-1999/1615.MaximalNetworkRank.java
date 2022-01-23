class Solution {
    public int maximalNetworkRank(int n, int[][] roads) {
        // O(n^2)
        boolean[][] connected = new boolean[n][n];
        int[] cnts = new int[n];
        for (int[] road :roads) {
            cnts[road[0]]++;
            cnts[road[1]]++;
            connected[road[0]][road[1]] = true;
            connected[road[1]][road[0]] = true;
        }
        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i+1; j < n; j++) {
                ans = Math.max(ans, cnts[i] + cnts[j] - (connected[i][j] ? 1:0));
            }
        }
        return ans;
    }
}