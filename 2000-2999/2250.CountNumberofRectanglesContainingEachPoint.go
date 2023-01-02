func countRectangles(rectangles [][]int, points [][]int) []int {
    MaxHeight := 100
    m := make(map[int][]int)
    for _, p := range rectangles {
        m[p[1]] = append(m[p[1]], p[0])
    }
    for _, lengths := range m {
        sort.Ints(lengths)
    }
    
    cnts := make([]int, len(points))
    for i, p := range points {
        for y := p[1]; y <= MaxHeight; y++ {
            if len(m[y]) > 0 {
                cnts[i] += len(m[y]) - binSearch(m[y], p[0])
            }
        }
    }
    return cnts
}

func binSearch(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}
