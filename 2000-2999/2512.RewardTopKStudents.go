import (
    "sort"
    "strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
    pos, neg := make(map[string]bool), make(map[string]bool)
    for _, word := range positive_feedback {
        pos[word] = true
    }
    for _, word := range negative_feedback {
        neg[word] = true
    }
    points := make([][]int, len(report))
    for i, r := range report {
        points[i] = []int{student_id[i], 0}
        for _, word := range strings.Split(r, " ") {
            if pos[word] {
                points[i][1] += 3
            }
            if neg[word] {
                points[i][1] -= 1
            }
        }
    }
    sort.Slice(points, func(i, j int) bool {
        if points[i][1] == points[j][1] {
            return points[i][0] < points[j][0]
        }
        return points[i][1] > points[j][1]
    })
    ans := make([]int, k)
    for i := range ans {
        ans[i] = points[i][0]
    }
    return ans
}