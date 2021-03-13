class Solution {
    public int countStudents(int[] students, int[] sandwiches) {
        int[] cnt = {0, 0};
        for (int s : students) {
            cnt[s]++;
        }
        for (int i = 0; i < sandwiches.length; i++) {
            if (cnt[sandwiches[i]] <= 0) {
                return sandwiches.length - i;
            }
            cnt[sandwiches[i]]--;
        }
        return 0;
    }
}