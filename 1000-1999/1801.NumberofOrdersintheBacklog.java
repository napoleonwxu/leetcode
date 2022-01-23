class Solution {
    private static final int MOD = (int)1e9 + 7;
    
    public int getNumberOfBacklogOrders(int[][] orders) {
        PriorityQueue<int[]> buy = new PriorityQueue<>((a, b) -> b[0] - a[0]);
        PriorityQueue<int[]> sell = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        for (int[] order : orders) {
            if (order[2] == 0) {
                buy.offer(order);
            } else {
                sell.offer(order);
            }
            while (!buy.isEmpty() && !sell.isEmpty() && buy.peek()[0] >= sell.peek()[0]) {
                int amount = Math.min(buy.peek()[1], sell.peek()[1]);
                buy.peek()[1] -= amount;
                if (buy.peek()[1] == 0) {
                    buy.poll();
                }
                sell.peek()[1] -= amount;
                if (sell.peek()[1] == 0) {
                    sell.poll();
                }
            }
        }
        int cnt = 0;
        for (int[] order : buy) {
            cnt = (cnt + order[1]) % MOD;
        }
        for (int[] order : sell) {
            cnt = (cnt + order[1]) % MOD;
        }
        return cnt;
    }
}