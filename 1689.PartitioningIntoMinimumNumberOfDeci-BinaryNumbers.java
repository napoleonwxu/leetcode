class Solution {
    public int minPartitions(String n) {
        int cnt = 0;
        for (int i = 0; i < n.length(); i++) {
            cnt = Math.max(cnt, n.charAt(i)-'0');
        }
        /*
        char[] num = n.toCharArray();
        boolean notZero = true;
        while (notZero) {
            notZero = false;
            for (int i = 0; i < num.length; i++) {
                num[i] -= num[i] > '0' ? 1 : 0;
                if (num[i] > '0') {
                    notZero = true;
                }
            }
            cnt++;
        } */
        return cnt;
    }
}