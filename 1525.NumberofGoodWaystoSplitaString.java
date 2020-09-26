class Solution {
    public int numSplits(String s) {
        int n = s.length(), ans = 0;
        int[] prefix = new int[n];
        HashSet set = new HashSet<>();
        for (int i = 0; i < n; i++) {
            set.add(s.charAt(i));
            prefix[i] = set.size();
        }
        set.clear();
        for (int i = n-1; i > 0; i--) {
            set.add(s.charAt(i));
            if (set.size() == prefix[i-1]) {
                ans++;
            }
        }
        return ans;
    }
}