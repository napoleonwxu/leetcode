class Solution {
    private static final int N = 26;
    
    public boolean closeStrings(String word1, String word2) {
        int n1 = word1.length(), n2 = word2.length();
        if (n1 != n2) {
            return false;
        }
        int[] chs1 = new int[N], chs2 = new int[N];
        for (int i = 0; i < n1; i++) {
            chs1[word1.charAt(i)-'a']++;
            chs2[word2.charAt(i)-'a']++;
        }
        for (int i = 0; i < N; i++) {
            if ((chs1[i] == 0 && chs2[i] != 0) || (chs1[i] != 0 && chs2[i] == 0)) {
                return false;
            }
        }
        Arrays.sort(chs1);
        Arrays.sort(chs2);
        for (int i = 0; i < N; i++) {
            if (chs1[i] != chs2[i]) {
                return false;
            }
        }
        return true;
    }
}