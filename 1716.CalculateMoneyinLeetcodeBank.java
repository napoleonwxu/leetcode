class Solution {
    public int totalMoney(int n) {
        int weeks = n/7, extra = n%7;
        return (1+7)*7/2*weeks + 7*weeks*(weeks-1)/2 + (1+extra)*extra/2 + weeks*extra;
    }
}