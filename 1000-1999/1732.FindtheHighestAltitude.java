class Solution {
    public int largestAltitude(int[] gain) {
        int highest = 0, altitude = 0;
        for (int g : gain) {
            altitude += g;
            highest = Math.max(highest, altitude);
        }
        return highest;
    }
}