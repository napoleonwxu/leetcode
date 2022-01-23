func count(nums []int) []int {
    cnt := make([]int, 7)
    for _, n := range nums {
        cnt[n]++
    }
    return cnt
}

func getMinR(A []int, B []int, n int) int {
    cnt := 0
    for i := 0; i < len(A); i++ {
        if A[i] != n {
            if B[i] != n {
                return -1
            } else {
                cnt++
            }
        }
    }
    return cnt
}

func minDominoRotations(A []int, B []int) int {
	// Python solution with set is better
    cntA, cntB := count(A), count(B)
    num := 0
    for i := 1; i <= 6; i++ {
        if cntA[i] + cntB[i] >= len(A) {
            num = i
            break
        }
    }
    if num == 0 {
        return -1
    }
    minA, minB := getMinR(A, B, num), getMinR(B, A, num)
    if minA == -1 {
        return minB
    }
    if minB == -1 {
        return minA
    }
    if minA <= minB {
        return minA
    }
    return minB
}