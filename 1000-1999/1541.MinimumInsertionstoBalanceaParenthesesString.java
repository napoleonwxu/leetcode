class Solution {
    public int minInsertions(String s) {
        int ans = 0, right = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                if (right%2 == 1) {
                    right--;
                    ans++;
                }
                right += 2;
            } else {
                right--;
                if (right < 0) {
                    right += 2;
                    ans++;
                }
            }
        }
        return ans + right;
    }
}