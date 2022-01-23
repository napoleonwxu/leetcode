class Solution {
    public int[] bestCoordinate(int[][] towers, int radius) {
        int[] ans = new int[2];
        int qMax = 0;
        for (int[] a : towers) {
            int q = 0;
            for (int[] b : towers) {
                int dd = (a[0]-b[0]) * (a[0]-b[0]) + (a[1]-b[1]) * (a[1]-b[1]);
                if (dd <= radius*radius) {
                    q += (int)(b[2] / (1 + Math.sqrt(dd)));
                }
            }
            if (q > qMax || (q == qMax && (a[0] < ans[0] || (a[0] == ans[0] && a[1] < ans[1])))) {
                qMax = q;
                ans[0] = a[0];
                ans[1] = a[1];
            }
        }
        return ans;
    }
}