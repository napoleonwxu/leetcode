class Solution {
    public int getXORSum(int[] arr1, int[] arr2) {
        return getXOR(arr1) & getXOR(arr2);
    }
    
    private int getXOR(int[] arr) {
        int xor = 0;
        for (int a : arr) {
            xor ^= a;
        }
        return xor;
    }
}