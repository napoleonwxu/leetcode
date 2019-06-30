// 1. sort: O(NlgN)
// 2. sort with heap: O(NlgK), refer 692.Top K Frequent Words
// 3. quick select: average O(N) <= O(N) + O(N/2) + O(N/4) + ... => worst: O(N^2)

func kClosest(points [][]int, K int) [][]int {
    // quick select: average O(N), and result is not sorted
    left, right := 0, len(points)-1
    for left <= right {
        mid := partition(points, left, right)
        if mid == K {
            break
        } else if mid < K {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    /* sort: O(NlgN)
    sort.Slice(points, func(i, j int) bool {
        return dis(points[i]) < dis(points[j])
    })
    */
    return points[:K]
}

func partition(points [][]int, l, r int) int {
    pivot, pivot_dis := points[l], dis(points[l])
    for l < r {
        for ; l < r && dis(points[r]) >= pivot_dis; r-- {}
        points[l] = points[r]
        for ; l < r && dis(points[l]) <= pivot_dis; l++ {}
        points[r] = points[l]
    }
    points[l] = pivot
    return l
}

func dis(p []int) int {
    return p[0]*p[0] + p[1]*p[1]
}
