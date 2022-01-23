class Solution {
    public double averageWaitingTime(int[][] customers) {
        long time = 0, cost = 0;
        for (int[] customer : customers) {
            time = Math.max(time, customer[0]) + customer[1];
            cost += time - customer[0];
        }
        return (double)cost / customers.length;
    }
}