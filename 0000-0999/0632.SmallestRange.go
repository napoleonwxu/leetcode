func merge(num1 [][]int, num2 [][]int) [][]int {
    var nums [][]int
    n1, n2 := len(num1), len(num2)
    i, j := 0, 0
    for i<n1 && j<n2 {
        if num1[i][0] < num2[j][0] {
            nums = append(nums, num1[i])
            i++
        } else {
            nums = append(nums, num2[j])
            j++
        }
    }
    if i < n1 {
        nums = append(nums, num1[i:]...)
    } else {
        nums = append(nums, num2[j:]...)
    }
    return nums
}

func divideConquer(nums_idx [][][]int, left int, right int) [][]int {
    if left == right {
        return nums_idx[left]
    }
    mid := (left+right) >> 1
    num_idx1 := divideConquer(nums_idx, left, mid)
    num_idx2 := divideConquer(nums_idx, mid+1, right)
    return merge(num_idx1, num_idx2)
}

func smallestRange(nums [][]int) []int {
    k := len(nums)
    var ns [][]int
    
    for i, num := range nums {
        for _, n := range num {
            ns = append(ns, []int{n, i})
        }
    }
    // 56ms < 47.62%; 10.6MB < 5.55%
    sort.Slice(ns, func (i, j int) bool {
        return ns[i][0] < ns[j][0]
    })
    /*
    var nums_idx [][][]int
    for i, num := range nums {
        var num_idx [][]int
        for _, n := range num {
            num_idx = append(num_idx, []int{n, i})
        }
        nums_idx = append(nums_idx, num_idx)
    }
    // 144ms < 13.64%; 26.8MB < 5.55% ???
    ns = divideConquer(nums_idx, 0, k-1)
    */
    var ans []int
    Map := make(map[int]int)
    left := 0
    for right, n := range ns {
        Map[n[1]]++
        for len(Map) == k {
            if len(ans)==0 || (ans[1]-ans[0] > ns[right][0]-ns[left][0]) {
                ans = []int{ns[left][0], ns[right][0]}
            }
            Map[ns[left][1]]--
            if Map[ns[left][1]] == 0 {
                delete(Map, ns[left][1])
            }
            left++
        }
    }
    return ans
}