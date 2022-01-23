class Solution {
    public boolean checkOnesSegment(String s) {
        // s[0] must be '1'
        return !s.contains("01");
        /* s[0] can be '0' or '1'
        boolean start = false, end = false;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '1') {
                if (!start) {
                    start = true;
                } else if (end) {
                    return false;
                }
            } else {
                if (start) {
                    end = true;
                }
            }
        }
        return true;
        */
    }
}