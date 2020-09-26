class Solution {
    public int kthFactor(int n, int k) {
        List<Integer> factors = new ArrayList<>();
        for (int i = 1; i*i <= n; i++) {
            if (n%i == 0) {
                if (--k == 0) {
                    return i;
                }
                if (i*i != n) {
                    factors.add(i);
                }
            }
        }
        if (factors.size() >= k) {
            return n / factors.get(factors.size()-k);
        }
        return -1;
    }
}