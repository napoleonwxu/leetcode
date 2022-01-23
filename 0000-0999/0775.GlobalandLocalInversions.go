// global inversions == local inversions: only swap the adjacent value 
func isIdealPermutation(A []int) bool {
    n := len(A)
    for i := 0; i < n; i++ {
        if abs(A[i]-i) > 1 {
            return false
        }
    }
    /*
    for i := 0; i < n-1; {
        if A[i] == i {
            i++
        } else if A[i] == i+1 && A[i+1] == i {
            i += 2
        } else {
            return false
        }
    }*/
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}