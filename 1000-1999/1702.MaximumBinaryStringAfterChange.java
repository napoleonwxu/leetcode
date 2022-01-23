class Solution {
    public String maximumBinaryString(String binary) {
        int leadingOnes = 0, zeros = 0, n = binary.length();
        for (int i = 0; i < n; i++) {
            if (binary.charAt(i) == '0') {
                zeros++;
            } else if (zeros == 0) {
                leadingOnes++;
            }
        }
        if (leadingOnes >= n) {
            return binary;
        }
        return "1".repeat(leadingOnes+zeros-1) + "0" + "1".repeat(n-leadingOnes-zeros);
    }
}