class Solution {
    public double average(int[] salary) {
        double sum = 0;
        int max = -1, min = (int)1e6 + 1;
        for (int s : salary) {
            min = Math.min(min, s);
            max = Math.max(max, s);
            sum += s;
        }
        return (sum - min - max) / (salary.length - 2);
    }
}