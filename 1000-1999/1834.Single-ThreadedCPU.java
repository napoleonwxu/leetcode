class Solution {
    public int[] getOrder(int[][] tasks) {
        int n = tasks.length;
        int[][] extTasks = new int[n][3];
        for (int i = 0; i < n; i++) {
            extTasks[i][0] = tasks[i][0];
            extTasks[i][1] = tasks[i][1];
            extTasks[i][2] = i;
        }
        Arrays.sort(extTasks, (a, b) -> a[0]-b[0]);
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[1] == b[1] ? a[2]-b[2] : a[1]-b[1]);
        int time = 0, idx = 0, i = 0;
        int[] seq = new int[n];
        while (i < n) {
            while (idx < n && extTasks[idx][0] <= time) {
                pq.offer(extTasks[idx++]);
            }
            if (pq.isEmpty()) {
                time = extTasks[i][0];
                continue;
            }
            int[] task = pq.poll();
            seq[i++] = task[2];
            time += task[1];
        }
        return seq;
    }
}