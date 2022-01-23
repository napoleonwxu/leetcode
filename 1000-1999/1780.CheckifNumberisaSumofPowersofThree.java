class Solution {
    public boolean checkPowersOfThree(int n) {
        while (n > 1 && n%3 != 2) {
            n = n%3 == 0 ? n/3 : n-1;
        }
        return n == 1;
        /*
        if (n == 1) {
            return true;
        }
        if (n%3 == 2) {
            return false;
        }
        return n%3 == 0 ? checkPowersOfThree(n/3) : checkPowersOfThree(n-1);
        */
    }
}