// 1. sort: O(NlgN)
// 2. sort with heap: O(NlgK), refer 692.Top K Frequent Words
// 3. quick select: average O(N) <== O(N) + O(N/2) + O(N/4) + ... = O(2N) ==> worst: O(N^2)

func kClosest(points [][]int, K int) [][]int {
    // quick select: average O(N), and result is not sorted
    left, right := 0, len(points)-1
    for left < right {
        mid := partition(points, left, right)
        if mid == K-1 {
            break
        } else if mid < K-1 {
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
    dis := points[r][0]*points[r][0] + points[r][1]*points[r][1]
    mid := l
    for i := l; i < r; i++ {
        if points[i][0]*points[i][0] + points[i][1]*points[i][1] < dis {
            points[i], points[mid] = points[mid], points[i]
            mid++
        }
    }
    points[mid], points[r] = points[r], points[mid]
    return mid
}
/*
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
*/