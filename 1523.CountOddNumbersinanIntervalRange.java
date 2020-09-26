class Solution {
    public int countOdds(int low, int high) {
        return (high+1)/2 - low/2;
        /*
        if (low%2 == 1) {
            low--;
        }
        if (high%2 == 1) {
            high++;
        }
        return (high-low) >> 1;
        */
    }
}