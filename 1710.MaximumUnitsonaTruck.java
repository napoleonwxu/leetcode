class Solution {
    public int maximumUnits(int[][] boxTypes, int truckSize) {
        Arrays.sort(boxTypes, (a, b) -> b[1] - a[1]);
        int ans = 0;
        for (int i = 0; i < boxTypes.length && truckSize > 0; i++) {
            int cnt = Math.min(truckSize, boxTypes[i][0]);
            ans += cnt * boxTypes[i][1];
            truckSize -= cnt;
        }
        return ans;
    }
}