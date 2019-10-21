/*
Given three integer arrays arr1, arr2 and arr3 sorted in strictly increasing order, return a sorted array of only the integers that appeared in all three arrays.

Example 1:
Input: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
Output: [1,5]
Explanation: Only 1 and 5 appeared in the three arrays.
 
Constraints:
1 <= arr1.length, arr2.length, arr3.length <= 1000
1 <= arr1[i], arr2[i], arr3[i] <= 2000
*/

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    minMax := min(arr1[len(arr1)-1], arr2[len(arr2)-1])
    minMax = min(minMax, arr3[len(arr3)-1])
    Map1 := make(map[int]bool)
    for _, num := range arr1 {
        if num > minMax {
            break
        }
        Map1[num] = true
    }
    Map2 := make(map[int]bool)
    for _, num := range arr2 {
        if num > minMax {
            break
        }
        if Map1[num] {
            Map2[num] = true
        }
    }
    ans := []int{}
    for _, num := range arr3 {
        if num > minMax {
            break
        }
        if Map2[num] {
            ans = append(ans, num)
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}