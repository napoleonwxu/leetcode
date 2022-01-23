class Solution {
    public int[] decode(int[] encoded) {
        int n = encoded.length + 1, x = 0;
        for (int i = 1; i <= n; i++) {
            x ^= i;
        }
        int[] perm = new int[n];
        perm[0] = x;
        for (int i = 1; i < n-1; i += 2) {
            perm[0] ^= encoded[i];
        }
        for (int i = 1; i < n; i++) {
            perm[i] = perm[i-1] ^ encoded[i-1];
        }
        return perm;
    }
}