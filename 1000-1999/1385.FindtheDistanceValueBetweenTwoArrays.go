func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// O(mn)
    cnt := 0
    for _, num1 := range arr1 {
        flag := 1
        for _, num2 := range arr2 {
            if abs(num1, num2) <= d {
                flag = 0
                break
            }
        }
        cnt += flag
    }
    return cnt
}

func abs(num1, num2 int) int {
    if num1 >= num2 {
        return num1-num2
    }
    return num2-num1
}